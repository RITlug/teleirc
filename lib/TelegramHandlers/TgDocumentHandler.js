/*
The MIT License (MIT)

Copyright (c) 2016 RIT Linux Users Group

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/



const TgHelpers = require('./TgHelpers.js');
const Helpers = require('../Helpers.js');

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
     * @param {string} caption - The caption of the document, if any
     */
    async RelayDocumentMessage(from, document, caption) {
        let username = '\x02' + TgHelpers.ResolveUserName(from) + '\x02';
        let url = await this._tgbot.getFileLink(document.file_id);

        let message = this._GetDocumentMessage(caption, username, document, url);
        this._action(message);
    }

    /**
     * Generates the message to send out
     * @param {string} caption - The caption of the document
     * @param {string} from - User who sent the photo
     * @param {string} document - Telegram document object
     * @param {string} fileUrl - URL to the document
    */
    _GetDocumentMessage(caption, from, document, fileUrl) {
        var message;

        if (Helpers.IsNullOrUndefined(caption)) {
            caption = "Untitled Document";
        }

        if (!this.Enabled) {
            message = from +
                    " shared a file (" +
                    document.mime_type +
                    ") on Telegram with caption: " +
                    "'" + caption + "'";
        } else {
            message = "'" + caption + "'" +
            " uploaded by " +
            from + ": " +
            fileUrl;
        }
    return message;
    }
}

module.exports = TgDocumentHandler;
