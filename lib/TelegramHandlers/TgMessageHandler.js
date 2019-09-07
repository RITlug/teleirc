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



const TgHelpers = require('./TgHelpers.js');

/**
 * Handles the event when a user sends a message to the Telegram Group.
 */
class TgMessageHandler {

    // ---------------- Constructor ----------------

    /**
     * 
     * @param {*} ircConfig - IRC Configuration that specifies what the
     *                        username's prefix and suffix shall be, and the
     *                        maximum message length that can be sent through IRC.
     * @param {boolean} enabled - Is this handler enabled?
     * @param {*} action - The action to take when this handler is fired.
     *                     Only parameter is a string, which is the message
     *                     to send out.
     */
    constructor(ircConfig, enabled, action) {
        this._ircConfig = ircConfig;
        this.Enabled = enabled;
        this._action = action;
    }

    // ---------------- Functions ----------------

    /**
     * 
     * @param {*} from - Object that contains the information about the user name.
     * @param {string} userMessage - The message the user sent that we want to relay.
     */
    RelayMessage(from, userMessage) {
        const self = this;
        if (!self.Enabled) {
            return;
        }

        let username = '\x02' + TgHelpers.ResolveUserName(from) + '\x02';
        const messagePrefix = self._ircConfig.prefix + username + self._ircConfig.suffix + " ";
        const messageSplitRegex = new RegExp(`.\{1,${self._ircConfig.maxMessageLength - messagePrefix.length}}`, 'g');

        // split messages based on max message length, and prepend username on each new line
        userMessage.toString().split(/\r?\n/).filter(function(line) {
          return line.length > 0;
        }).forEach(function(line) {
          var linesToSend = line.match(messageSplitRegex);
          linesToSend.forEach(function(toSend) {
            self._action(messagePrefix + toSend);
          });
        });
    }
}

module.exports = TgMessageHandler;
