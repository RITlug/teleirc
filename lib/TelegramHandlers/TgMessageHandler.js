const TgHelpers = require('./TgHelpers.js');

/**
 * Handles the event when a user sends a message to the Telegram Group.
 */
class TgMessageHandler {

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
     * @param {string} userMessage - The message the user sent that we want to relay.
     */
    RelayMessage(from, userMessage) {
        if (this.Enabled === false) {
            return;
        }

        let username = TgHelpers.CleanUpUserName(from);
        let message = this._prefixSuffixConfig.prefix + username + this._prefixSuffixConfig.suffix + " " + userMessage;

        this._action(message);
    }
}

module.exports = TgMessageHandler;