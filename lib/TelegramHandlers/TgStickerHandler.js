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
