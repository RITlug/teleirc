const TgHelpers = require('./TgHelpers.js');

/**
 * Handles the event when a user leaves the telegram channel.
 * Named "PartHandler" to be consistent with the equivalent IRC class.
 */
class TgPartHandler {

    // ---------------- Constructor ----------------

    /**
     * 
     * @param {boolean} enabled - Is this handler enabled?
     * @param {*} action - The action to take when this handler is fired.
     *                     Only parameter is a string, which is the message
     *                     to send out.
     */
    constructor(enabled, action) {
        this.Enabled = enabled;
        this._action = action;
    }

    // ---------------- Functions ----------------

    /**
     * Constructs the message and fires this class's callback.
     * @param {*} partedUser - Object that contains the information about the user name.
     */
    RelayPartMessage(partedUser) {
        if (!this.Enabled) {
            return;
        }

        let username = TgHelpers.GetFullUserName(partedUser);
        let message = username + " has left the Telegram Group.";

        this._action(message);
    }
}

module.exports = TgPartHandler;
