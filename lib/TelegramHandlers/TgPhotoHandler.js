const TgHelpers = require('TgHelpers.js');
const imgur = require("imgur");
const Helpers = require('../Helpers.js');

/**
 * Handles the event when a user sends a photo to the Telegram Group.
 */
class TgPhotoHandler {

    // ---------------- Constructor ----------------

    /**
     * 
     * @param {*} _useImgur - Should the photo be uploaded to imgur?
     * @param {boolean} enabled - Is this handler enabled?
     * @param {*} tgbot - Reference to the Telegram Bot API.
     * @param {*} action - The action to take when this handler is fired.
     *                     Only parameter is a string, which is the message
     *                     to send out.
     */
    constructor(useImgur, enabled, tgbot, action) {
        this._useImgur = useImgur;
        this.Enabled = enabled;
        this._tgbot = tgbot;
        this._action = action;
    }

    // ---------------- Functions ----------------

    /**
     * 
     * @param {*} from - Object that contains the information about the user name.
     * @param {*} photo - The information about the photo.
     * @param {string} - The caption of the photo, if any.
     */
    ReportPhoto(from, photo, caption) {
        if (this.Enabled === false) {
            return;
        }

        let username = TgHelpers.CleanUpUserName(from);
        let photo = _PickPhoto(photo);
        if (Helpers.IsNullOrUndefined(photo))
        {
            // If there is no photo, do nothing, and return.
            return;
        }

        let parent = this;

        // To get a photo, we must first get the URL of the photo.

        let tgUrl = await this._tgbot.getFileLink(photo.file_id);

        var url;
        if (this._ImgurUpload) {
            url = this._ImgurUpload(tgUrl);
        }
        else {
            url = tgUrl;
        }

        if (Helpers.IsNullOrUndefined(url) === false) {
            let message = this._GetPhotoMessage(caption, from, url);
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
     * Uploads the photo to imgur
     * @param {string} photoUrl - The URL to the photo to download from Telegram's servers.
     * @returns The imgur URL if successful, null if not.
     */
    _ImgurUpload(photoUrl) {
        try
        {
            let json = await imgur.uploadUrl(photoUrl);
            return json.data.link;
        }
        catch(err)
        {
            console.error("Error when uploading to imgur: " + err.message);
            return null;
        }
    }
}

module.exports = TgPhotoHandler;