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
 * Handles the event when a user sends a photo to the Telegram Group.
 */
class TgPhotoHandler {

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
     * @param {*} photoList - The information about the photos we can get from Telegram.
     * @param {string} caption - The caption of the photo, if any.
     */
    async RelayPhotoMessage(from, photoList, caption) {
        if (!this.Enabled) {
            return;
        }

        let username = '\x02' + TgHelpers.ResolveUserName(from) + '\x02';
        let photo = this._PickPhoto(photoList);
        if (Helpers.IsNullOrUndefined(photo))
        {
            // If there is no photo, do nothing, and return.
            return;
        }

        // To send a photo, we must first get its URL
        let tgUrl = await this._tgbot.getFileLink(photo.file_id);

        // Then, upload it somewhere, classes that extend this class determine that.
        let url = await this._UploadImage(tgUrl);

        if (Helpers.IsNullOrUndefined(url) === false) {
            let message = this._GetPhotoMessage(caption, username, url);
            this._action(message);
        }
    }

    /**
     * Picks the best photo from the photo array
     * to send to the IRC channel.
     * @param {Array} photoInfo
     * @returns The best photo, or null if there are none (e.g. empty array passed in).
     */
    _PickPhoto(photoInfo) {
        var photo = null;

        // When we upload a photo to telegram,
        // Telegram has several versions of the photos
        // of various sizes.  We'll
        // pick the biggest one (more resolution is always
        // best!).
        for (var i = 0; i < photoInfo.length; ++i) {
            var innerPhoto = photoInfo[i];
            if (photo === null) {
                photo = innerPhoto;
            } else if (photo.file_size < innerPhoto.file_size) {
                photo = innerPhoto;
            }
        }
  
        return photo;
    }

    /**
     * Generates the message to send out.
     * @param {string} caption - The caption of the photo.
     * @param {string} from - Who sent the photo?
     * @param {string} fileUrl - URL to the photo.
     */
    _GetPhotoMessage(caption, from, fileUrl) {
        var message;
        if (Helpers.IsNullOrUndefined(caption)) {
            message = "'Untitled Image'"
        } else {
            message = "'" + caption + "'"
        }
        message += " uploaded by " + from + ": " + fileUrl;
    
        return message;
    }

    /**
     * Uploads the image from Telegram's URL, in case we mirror the photo somewhere else.
     * By default, this just reports Telegram's URL, which will eventually expire.
     * However, classes that extend this class can override this function
     * to upload the image wherever they want.
     */
    async _UploadImage(tgUrl) {
        return Promise.resolve(tgUrl);
    }
}

module.exports = TgPhotoHandler;
