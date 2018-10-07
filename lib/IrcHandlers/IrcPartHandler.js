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
