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
