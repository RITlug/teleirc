TeleIRC Code Conventions
========================

This page explains TeleIRC coding best practices.
This was originally written by Tim Zabel ([@Tjzabel](https://github.com/Tjzabel)).


## Naming Conventions

TeleIRC constitutes two halves: IRC and Telegram.
Function names should be agnostic to each platform. 
Lastly, function names should be consistent across IRC and Telegram where possible.

```go
func (*tg Client) SendMessage(msg string) {
    ...
}
```


## Handlers

Handlers are blocks of code responsible for "handling" specific message types.
Such handlers should be named appropriately in camelCase.

```go
func joinHandler(...)
```
