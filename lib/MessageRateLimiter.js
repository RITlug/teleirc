/*
The MIT License (MIT)

Copyright (c) 2016 RIT Linux Users Group

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/



'use strict';

class MessageBundle {

  constructor() {
    this.queue = [];
  }

  addMessage(message) {
    this.queue.push(message);
  }

  implodeAndClear() {
    let m = this.queue.join("\n");
    this.queue = [];
    return m;
  }
}

class MessageRateLimiter {

  /**
   * rate: how many messages we can send...
   * per: ...per this many seconds
   * sendAction: function to send a message
   *
   * Providing a rate of "0" disables rate limiting.
   */
  constructor(rate, per, sendAction) {
    this.rate = rate;
    this.per = per;
    this.allowance = rate;
    this.last_check = Date.now() / 1000;
    this.bundle = new MessageBundle();
    this.sendAction = sendAction;

    // We need to run periodically to make sure messages don't get stuck
    // in the queue.
    if (this.rate > 0) {
      setInterval(this.run.bind(this), 2000);
    }
  }

  queueMessage(message) {
    this.bundle.addMessage(message);
    // We call run here just in case we can immediately send
    // the message, instead of waiting for the setInterval to call
    // run for us.
    this.run();
  }

  run() {
    this.bumpAllowance();

    if (this.rate > 0 && this.allowance < 1) {
      console.log("A message has been received and rate limited");
      // Currently rate-limiting, so don't do anything.
    } else {
      if (this.bundle.queue.length > 0) {
        this.sendAction(this.bundle.implodeAndClear());
        this.allowance--;
      }
    }
  }

  bumpAllowance() {
    let current = Date.now() / 1000;
    let timePassed = current - this.last_check;
    this.last_check = current;
    this.allowance = this.allowance + (timePassed * this.rate / this.per);

    // Make sure we don't get to an allowance that's higher than the
    // rate we're actually allowed to send.
    if (this.allowance > this.rate) {
      this.allowance = this.rate;
    }
  }
}

module.exports = MessageRateLimiter;
