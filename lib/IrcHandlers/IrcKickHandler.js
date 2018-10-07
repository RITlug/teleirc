/**
 * Handles the event when a user is kicked from the IRC channel.
 */
class IrcKickHandler {

    // ---------------- Constructor ----------------

    /**
     * 
     * @param {boolean} enabled - If this handler enabled?
     * @param {Function} action - The action to take when this handler is fired.
     *                            Only parameter is a string, which is the kick message
     *                            to send out.
     */
    constructor(enabled, action) {
        this._action = action;
        this.Enabled = enabled;
    }

    // ---------------- Functions ----------------

    /**
     * Constructs the message and fires this class's callback.
     * @param {string} channel - The channel the user was kicked from.
     * @param {string} username - The user that was kicked from the channel.
     * @param {string} by - The user who performed the kick.
     * @param {string} reason - The reason why the user was kicked.  Optional.
     */
    ReportKick(channel, username, by, reason) {
        if (this.Enabled === false ) {
            return;
        }

        if (typeof reason !== "string") {
            reason = "Kicked";
        }

        let message = username + " was kicked by " + by + " from " + channel + ": " +
        reason + ".";
        this._action(message);
    }
}

module.exports = IrcKickHandler;
