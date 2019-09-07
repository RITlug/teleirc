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



const Helpers = require("../Helpers.js");
const escape = require('escape-html');

/**
 * Handles the event when a user sends a message to the IRC channel.
 */
class IrcMessageHandler {

    // ---------------- Constructor ----------------

    /**
     * 
     * @param {Array} ircBlackList - List of names to ignore. null/undefined to not ignore anyone.
     * @param {*} ircConfig - The bot's IRC Configuration.
     * @param {boolean} enabled - If this handler enabled?
     * @param {Function} action - The action to take when this handler is fired.
     *                            Only parameter is a string, which is the message
     *                            to send out.
     */
    constructor(ircBlackList, ircConfig, enabled, action) {
        this._ircBlackList = ircBlackList;
        this._ircConfig = ircConfig;
        this._action = action;
        this.Enabled = enabled;
    }

    // ---------------- Functions ----------------

    /**
     * Constructs the message and fires this class's callback.
     * @param {string} username - The user who sent the message.
     * @param {string} channel - The channel the user has sent the message in.
     * @param {string} userMessage - The message the user sent.
     */
    RelayMessage(username, channel, userMessage) {
        if (this.Enabled === false) {
            return;
        }

        // we want to ignore the messages from not configured channels or direct messages

        if (channel) {
            if (channel.toLowerCase() !== this._ircConfig.channel.toLowerCase()) {
                console.log(`Ignoring: <${username}@${channel}> ${userMessage}`);
                return;
            }
        }
        else {
            // If our channel is null/undefined, ignore.
            // Otherwise, we don't know where the message came from.
            // It could be a private message, which we don't care about.
            return;
        }
        var blackListed = Helpers.StringExistsIgnoreCase(
            this._ircBlackList,
            username
        );

        if (blackListed === false) {
            let message = "&lt;<b>" + username + "</b>&gt; " + escape(userMessage);
            this._action(message);
        }
    }
}

module.exports = IrcMessageHandler;
