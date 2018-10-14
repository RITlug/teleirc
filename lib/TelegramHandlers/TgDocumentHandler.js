const TgHelpers = require('./TgHelpers.js');

/**
 * Handles the event when a user sends a document to the Telegram Group.
 */
class TgDocumentHandler {

    // ---------------- Constructor ----------------

    /**
     * 
     * @param {boolean} enabled - Is this handler enabled?
     * @param {*} tgbot - Reference to the Telegram Bot API.
     * @param {*} action - The action to take when this handler is fired.
     *                     Only parameter is a string, which is the message
     *                     to send out.
     */
    constructor(enabled, tgbot, action) {
        this.Enabled = enabled;
        this._tgbot = tgbot;
        this._action = action;
    }

    // ---------------- Functions ----------------

    /**
     * 
     * @param {*} from - Object that contains the information about the user name.
     * @param {*} document - The information about the document.
     */
    async ReportDocument(from, document) {
        if (this.Enabled === false) {
            return;
        }

        let username = TgHelpers.CleanUpUserName(from);

        let url = await this._tgbot.getFileLink(document.file_id);

        let message = username +
                        " Posted File: " +
                        document.file_name + 
                        " (" + 
                        document.mime_type +
                        ", " + 
                        document.file_size_size +
                        " bytes): " +
                        url;

        this._action(message);
    }
}

module.exports = TgDocumentHandler;