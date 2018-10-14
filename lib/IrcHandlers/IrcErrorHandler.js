/**
 * Handles an error event.
 */
class IrcErrorHandler {

    // ---------------- Constructor ----------------

    /**
     * 
     * @param {boolean} enabled - If this handler enabled?
     * @param {Function} action - The action to take when this handler is fired.
     *                            Only parameter is a string, which is the message
     *                            to send out.
     */
    constructor(enabled, action) {
        this._action = action;
        this.Enabled = enabled;
    }

    // ---------------- Functions ----------------

    /**
     * Constructs the message and fires this class's callback.
     * @param {*} message - message object to report.
     */
    ReportError(message) {
        if (this.Enabled === false ) {
            return;
        }

        let msg = "[IRC Debug] " + JSON.stringify(message);
        this._action(msg);
    }
}

module.exports = IrcErrorHandler;
