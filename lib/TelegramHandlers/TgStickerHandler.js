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
const Helpers = require('../Helpers.js');

/**
 * Handles the event when a user sends a sticker to the Telegram Group.
 */
class TgStickerHandler {

    // ---------------- Constructor ----------------

    /**
     * 
     * @param {*} prefixSuffixConfig - Configuration that specifies what the
     *                                 username's prefix and suffix
     *                                 shall be.
     * @param {boolean} enabled - Is this handler enabled?
     * @param {*} action - The action to take when this handler is fired.
     *                     Only parameter is a string, which is the message
     *                     to send out.
     */
    constructor(prefixSuffixConfig, enabled, action) {
        this._prefixSuffixConfig = prefixSuffixConfig;
        this.Enabled = enabled;
        this._action = action;
    }

    // ---------------- Functions ----------------

    /**
     * 
     * @param {*} from - Object that contains the information about the user name.
     * @param {*} sticker - Information about the sticker that was sent.
     */
    RelayStickerMessage(from, sticker) {
        if (!this.Enabled) {
            return;
        }
        else if (Helpers.IsNullOrUndefined(sticker.emoji)) {
            // Do nothing if there is no emoji to send.
            return;
        }

        let username = '\x02' + TgHelpers.ResolveUserName(from) + '\x02';
        let message = this._prefixSuffixConfig.prefix + username + this._prefixSuffixConfig.suffix + " " + sticker.emoji;

        this._action(message);
    }
}

module.exports = TgStickerHandler;
