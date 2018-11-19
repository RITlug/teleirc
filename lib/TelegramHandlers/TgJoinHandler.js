const TgHelpers = require('./TgHelpers.js');

/**
 * Handles the event when a user joins the telegram channel.
 */
class TgJoinHandler {

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
     * @param {*} newUser - Object that contains the information about the user name.
     */
    RelayJoinMessage(newUser) {
        if (!this.Enabled) {
            return;
        }

        let username = TgHelpers.GetFullUserName(newUser);
        let message = username + " has joined the Telegram Group!";

        this._action(message);
    }
}

module.exports = TgJoinHandler;
