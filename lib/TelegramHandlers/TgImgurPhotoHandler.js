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



const TgPhotoHandler = require("./TgPhotoHandler");

/**
 * Handles the event when a user
 * sends a photo to the Telegram group.
 * This handler responds by uploading the
 * photo to imgur.
 */
class TgImgurPhotoHandler extends TgPhotoHandler {

    // ---------------- Constructor ----------------

    constructor(imgur, enabled, tgbot, action) {
        super(enabled, tgbot, action);
        this._imgur = imgur;
    }

    // ---------------- Functions ----------------

   /**
     * Uploads the photo to imgur
     * @param {string} photoUrl - The URL to the photo to download from Telegram's servers.
     * @returns The imgur URL if successful, null if not.
     */
    async _UploadImage(photoUrl) {
        try
        {
            let json = await this._imgur.uploadUrl(photoUrl);
            return json.data.link;
        }
        catch(err)
        {
            console.error("Error when uploading to imgur: " + err.message);
            return null;
        }
    }
}

module.exports = TgImgurPhotoHandler;
