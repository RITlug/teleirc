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
