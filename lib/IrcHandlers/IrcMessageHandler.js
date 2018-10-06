/**
 * Handles the event when a user sends a message to the IRC channel.
 */
class IrcMessageHandler {

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
     * @param {string} username - The user who sent the message.
     * @param {string} channel - The channel the user has sent the message in.
     * @param {string} userMessage - The message the user sent.
     */
    ReportMessage(username, channel, userMessage) {
        if (this.Enabled == false ) {
            return;
        }

        // Search the black-listed names.
        var sendMessage;

        // If there is no black-list, always send the message.
        if ((this._ircBlackList === undefined) || (this._ircBlackList === null))
        {
            sendMessage = true;
        }
        else
        {
            let matchedName = this._ircBlackList.filter(
                (name) => { return username.toLowerCase() === name.toLowerCase(); }
            );
            sendMessage = (matchedName.length <= 0);
        }

        if (sendMessage) {
            let message = "<" + username + "> " + userMessage;
            this._action(message);
        }
    }
}

module.exports = IrcMessageHandler;
