/**
 * Handles the event when a user joins the IRC channel.
 */
class IrcJoinHandler {

    // ---------------- Constructor ----------------

    /**
     * 
     * @param {boolean} enabled - If this handler enabled?
     * @param {Function} action - The action to take when this handler is fired.
     *                            Only parameter is a string, which is the join message
     *                            to send out.
     */
    constructor(enabled, action) {
        this._action = action;
        this.Enabled = enabled;
    }

    // ---------------- Functions ----------------

    /**
     * Constructs the message and fires this class's callback.
     * @param {string} channel - The channel the user has joined.
     * @param {string} username - The user that joined the channel.
     */
    ReportJoin(channel, username) {
        if (this.Enabled === false ) {
            return;
        }

        let message = username + " has joined " + channel + ".";
        this._action(message);
    }
}

module.exports = IrcJoinHandler;
