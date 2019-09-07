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
 * Handles the event when a user parts (leaves) the IRC channel.
 */
class IrcPartHandler {

    // ---------------- Constructor ----------------

    /**
     * 
     * @param {boolean} enabled - If this handler enabled?
     * @param {Function} action - The action to take when this handler is fired.
     *                            Only parameter is a string, which is the part message
     *                            to send out.
     */
    constructor(enabled, action) {
        this._action = action;
        this.Enabled = enabled;
    }

    // ---------------- Functions ----------------

    /**
     * 
     * @param {string} channel - The channel the user has left.
     * @param {string} username - The user that parted (left) from the channel.
     * @param {string} reason - The reason why the user left.  Optional.
     */
    ReportPart(channel, username, reason) {
        if (this.Enabled === false ) {
            return;
        }

        if (typeof reason !== "string") {
            reason = "Parting..";
        }

        let message = username + " has left " + channel + ": " + reason + ".";
        this._action(message);
    }
}

module.exports = IrcPartHandler;
