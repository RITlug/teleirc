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



/**
 * Class that handles errors that can occur when using Telegram.
 */
class TgErrorHandler {

    // ----------------- Constructor -----------------

    constructor(enabled, action) {
        this.Enabled = enabled;
        this._action = action;
    }

    // ----------------- Functions -----------------

    /**
     * Use this handler when we get a chat id we were not expecting.
     * @param {*} chat - The chat object whose ID is invalid.
     * @param {*} msg - The message object that we received.
     */
    HandleBadChatId(chat, msg) {
        if (!this.Enabled) {
            return;
        }

        // Messages that are sent to the bot outside of the group chat should just be dumped
        // to the console for potential testing and debugging to do things like check chat IDs
        // and verify the JSON formats of various messages
        let message = "[TG Debug] - Unexpected chat ID '" + chat.id + "' from message: " + JSON.stringify(msg);
        this._action(message);
    }

    /**
     * Use this handler when we get a message whose type (e.g. sticker, photo, etc)
     * is unknown.
     * @param {*} msg - The message object that we received.
     */
    HandleUnknownMessageType(msg) {
        if (!this.Enabled) {
            return;
        }

        let message = "[TG Debug] Ignoring non-text message: " + JSON.stringify(msg);
        this._action(message);
    }
}

module.exports = TgErrorHandler;
