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
