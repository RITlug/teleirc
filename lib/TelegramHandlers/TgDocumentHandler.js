const TgHelpers = require('./TgHelpers.js');

/**
 * Handles the event when a user sends a document to the Telegram Group.
 */
class TgDocumentHandler {

    // ---------------- Constructor ----------------

    /**
     * 
     * @param {*} prefixSuffixConfig - Configuration that specifies what the
     *                                 username's prefix and suffix
     *                                 shall be.
     * @param {boolean} enabled - Is this handler enabled?
     * @param {*} tgbot - Reference to the Telegram Bot API.
     * @param {*} action - The action to take when this handler is fired.
     *                     Only parameter is a string, which is the message
     *                     to send out.
     */
    constructor(prefixSuffixConfig, enabled, tgbot, action) {
        this._prefixSuffixConfig = prefixSuffixConfig;
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
    ReportDocument(from, document) {
        if (this.Enabled === false) {
            return;
        }

        let username = TgHelpers.CleanUpUserName(from);

        this._tgbot.getFileLink(document.file_id).then(
            (url) => {
                let message = this._prefixSuffixConfig.prefix +
                              username +
                              this._prefixSuffixConfig.suffix +
                              " Posted File: "
                              document.file_name + 
                              " (" + 
                              document.mime_type +
                              ", " + 
                              document.file_size_size +
                              " bytes): " +
                              url;

                this._action(message);
            }
        );
    }
}

module.exports = TgDocumentHandler;