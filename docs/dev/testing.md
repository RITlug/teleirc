Go testing: Best practices
==========================

This page explains our best practices for writing unit tests in Go.
This was originally written by contributor Nicholas Jones ([@Zedjones](https://github.com/Zedjones)).


## What is unit testing?

Unit testing is a type of testing where, as the name implies, we want to test an individual _unit_ of code.
In any given programming language, a _unit_ is the smallest piece of grouped code, usually a function.

In a unit test, we only want to test that this piece of code does its job.
We don't care if any other piece of code does its job, and optimally a unit test should pass even if every other unit of code that it calls fails.

To give an example from our current code base:

```go
func (c Client) StartBot(errChan chan<- error, sendMessage func(string)) {
	fmt.Println("Starting up IRC bot...")
	c.sendToTg = sendMessage
	c.addHandlers()
	if err := c.Connect(); err != nil {
		errChan <- err
	} else {
		errChan <- nil
	}
}
```

In this example, we care that our function does the following things:

* Passes "Starting up IRC bot..." to `fmt.Println`.
* Sets `c.sendToTg` to the `sendMessage` passed in.
* Calls `c.addHandlers`.
* Calls `c.Connect`
  * If an error is returned from this function, passed the error over the `errChan` channel passed in.
  * Otherwise, passes `nil` over the `errChan` channel.

Here are some things that we don't care about:

* What `fmt.Println` does with our string, that's its responsibility.
* What `c.addHandlers` does
* What `c.Connect` does, we only want its return values

In fact, we have an issue: if `c.Connect` _actually_ gets called, we can't guarantee that it will ever return what we want it to so we can actually test all of our cases.
Not to mention, we want our unit tests to pass even if there is no network connection available.

The solution to this: **mocking**.
Due to static typing and the lack of inheritance in Go, the only real way to do proper mocking is to use interfaces.


## Interfaces

An interface in Go is similar to most other languages.
It defines a contract that a concrete implementation must follow.

For example:

```go
type Shape interface {
    Area() float64
}
```

This interface defines what a shape is: anything that has an area.
A concrete implementation would be:

```go
type Rectangle struct {
    Width  float64
    Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
```

Another implementation might be:

```go
type Circle struct {
    Radius float64
}

func (c Circle) Area() float64 {
    return math.Pi * math.Pow(c.Radius, 2)
}
```

Now, we can create a function that accepts our interface:

```go
func PrintArea (s Shape) {
    fmt.Println(s.Area())
}

func main() {
	PrintArea(Rectangle{Width: 10, Height: 15}) // prints 150
	PrintArea(Circle{Radius: 15}) // prints 706.8583
}
```

### Why do we _need_ interfaces for testing?

Because we cannot modify the implementation of functions on structs directly, we must change our functions to accept _interfaces_ instead.

So, looking at one of our handlers:

```go
func connectHandler(c Client) func(*girc.Client, girc.Event) {
	return func(gc *girc.Client, e girc.Event) {
		c.Cmd.Join(c.Settings.Channel)
	}
}
```

In this case, there is a small problem: we called `c.Cmd.Join`.
So how does that work? How can we mock an inner struct inside of a struct?
Well, we define an interface as such:

```go
type IRCClient interface {
    Join(string)
}
```

Then, we modify our Client struct and add the following method:

```go
func (c Client) Join(channel string) {
	c.Cmd.Join(channel)
}
```

Now, our `Client` struct implements our interface and we can change the original function to accept our interface instead and change from `c.Cmd.Join` to `c.Join`:

```go
func connectHandler(c IRCClient) func(*girc.Client, girc.Event) {
	return func(gc *girc.Client, e girc.Event) {
		c.Join(c.Settings.Channel)
	}
}
```

While this does add complexity to our code, it also decouples the process of joining a channel from our exact library implementation.
Now, if we want to change our IRC library down the line, we just need to write a `Join` method that fulfills this contract.

As an added bonus, we can now write our own implementation of the `IRCClient` to use as a mock.
Something that fulfills the contract of the interface, but where the `Join` method doesn't actually do any work, just verifies that the string passed in is correct.
**Or**, even better, let's use a library to generate this for us.


## Generating mocks from interfaces with libraries

The most popular mock generation library is [`mockgen`](https://github.com/golang/mock).

So, to generate a mock from the previous interface, we would run the following commands:

```bash
$ GO111MODULE=on go get github.com/golang/mock/mockgen@latest
$ mockgen -source=./internal/handlers/irc.go
# OR
$ mockgen github.com/ritlug/teleirc/internal/handlers/irc IRCClient
```

Now, we could write the following code in a unit test:

```go
func TestConnectHandler(t *testing.T) {
  ctrl := gomock.NewController(t)

  // Assert that Join() is invoked.
  defer ctrl.Finish()

  m := NewMockIRCClient(ctrl)
  m.Settings = IRCSettings {Channel: "some channel"}

  // Asserts that the first and only call to Join() is passed "some channel".
  // Anything else will fail.
  m.
    EXPECT().
    Join(gomock.Eq("some channel"))

  connectHandler(m)(nil, nil) // Disclaimer: I didn't actually test this
}
```
