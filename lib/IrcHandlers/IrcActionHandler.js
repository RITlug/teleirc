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

/**
 * Handles the event when a user performs an action in the IRC channel.
 * An action, in this context, is a user sending the '/me' command.
 */
class IrcActionHandler {

    // ---------------- Constructor ----------------

    /**
     * 
     * @param {Array} ircBlackList - List of names to ignore. null/undefined to not ignore anyone.
     * @param {boolean} enabled - If this handler enabled?
     * @param {Function} action - The action to take when this handler is fired.
     *                            Only parameter is a string, which is the message
     *                            to send out.
     */
    constructor(ircBlackList, enabled, action) {
        this._ircBlackList = ircBlackList;
        this._action = action;
        this.Enabled = enabled;
    }

    // ---------------- Functions ----------------

    /**
     * Constructs the message and fires this class's callback.
     * @param {string} username - The user who performed the action.
     * @param {string} channel - The channel the user has performed the action in.
     * @param {string} userAction - The action the user performed.
     */
    ReportAction(username, channel, userAction) {
        if (this.Enabled === false ) {
            return;
        }

        var blackListed = Helpers.StringExistsIgnoreCase(
            this._ircBlackList,
            username
        );

        // Do not send the message if the user is black-listed.
        if (blackListed === false) {
            let message = username + " " + userAction;
            this._action(message);
        }
    }
}

module.exports = IrcActionHandler;
