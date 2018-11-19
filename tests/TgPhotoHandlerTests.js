'use strict';

const TgPhotoHandler = require('../lib/TelegramHandlers/TgPhotoHandler.js');
const TgImgurPhotoHandler = require("../lib/TelegramHandlers/TgImgurPhotoHandler.js");

let fromNoUsername = {
    first_name : "FirstName",
    username : undefined
};

let fromWithUserName = {
    first_name : "FirstName",
    username : "username"
};

let photoSmall = {
    file_id : 1,
    file_size : 1,
    url : "https://ritlug.com/test/1.png",
    imgurUrl : "https://imgur.com/small.png"
};

let photoMed = {
    file_id : 2,
    file_size : 2,
    url : "https://ritlug.com/test/2.png",
    imgurUrl : "https://imgur.com/med.png"
};

let photoLarge = {
    file_id : 3,
    file_size : 3,
    url : "https://ritlug.com/test/3.png",
    imgurUrl : "https://imgur.com/large.png"
};

// This photo is not in the imgur search,
// so if we pass in this photo in the handler,
// our mock imgur object will throw an exception.
let photoNotInImgur = {
    file_id : 4,
    file_size : 4,
    url : "https://ritlug.com/test/4.png"
}

let photos = [photoSmall, photoMed, photoLarge];

let mockTgBot = {
    files : {
        1 : photoSmall.url,
        2 : photoMed.url,
        3 : photoLarge.url
    },
    async getFileLink(id) {
        return Promise.resolve(this.files[id]);
    }
};

let mockImgur = {
    uploadUrl(tgUrl) {
        for (var i = 0; i < photos.length; ++i) {
            if (photos[i].url === tgUrl) {
                return Promise.resolve( 
                    {
                        data : {
                            link: photos[i].imgurUrl
                        }
                    }
                );
            }
        }

        return Promise.reject("Could not upload image " + tgUrl);
    }
};

/**
 * Ensures nothing happens it the handler is disabled.
 */
exports.TgPhotoHandler_DisabledTest = async function(assert) {
    var message = undefined;

    let uut = new TgPhotoHandler(
        false,
        mockTgBot,
        (msg) => {message = msg;}
    );

    await uut.RelayPhotoMessage(fromNoUsername, photos, "My Caption");

    assert.strictEqual(undefined, message);
    assert.done();
};

/**
 * Ensures that if we are enabled, but no caption exists, we
 * label the image as 'untitled'.
 */
exports.TgPhotoHandler_EnabledNoCaptionTest = async function(assert) {
    var message = undefined;

    let uut = new TgPhotoHandler(
        true,
        mockTgBot,
        (msg) => {message = msg;}
    );

    let expectedMessage =
        "'Untitled Image' uploaded by " + 
        fromNoUsername.first_name +
        ": " +
        photoLarge.url; // Should always pick large photo's URL.
        
    await uut.RelayPhotoMessage(fromNoUsername, photos, null);

    assert.strictEqual(expectedMessage, message);
    assert.done();
};

/**
 * Ensures that if we are enabled, and a caption exists, we
 * caption the image correctly.
 */
exports.TgPhotoHandler_EnabledCaptionTest = async function(assert) {
    var message = undefined;

    let uut = new TgPhotoHandler(
        true,
        mockTgBot,
        (msg) => {message = msg;}
    );

    let caption = "My caption";

    let expectedMessage =
        "'" +
        caption +
        "' uploaded by " + 
        fromWithUserName.username +
        ": " +
        photoLarge.url; // Should always pick large photo's URL.
        
    await uut.RelayPhotoMessage(fromWithUserName, photos, caption);

    assert.strictEqual(expectedMessage, message);
    assert.done();
};

/**
 * Tests the happy path to ensure we upload to imgur correctly.
 */
exports.ImgurPhotoHandler_SuccessTest = async function(assert) {
    var message = undefined;

    let uut = new TgImgurPhotoHandler(
        mockImgur,
        true,
        mockTgBot,
        (msg) => {message = msg;}
    );

    let caption = "My Caption";

    let expectedMessage =
        "'" +
        caption +
        "' uploaded by " + 
        fromWithUserName.username +
        ": " +
        photoLarge.imgurUrl; // Should always pick large photo's URL.

    await uut.RelayPhotoMessage(fromWithUserName, photos, caption);

    assert.strictEqual(expectedMessage, message);
    assert.done();
}

/**
 * Tests the failure path to ensure we don't crash if we can't
 * upload to imgur.
 */
exports.ImgurPhotoHandler_SuccessTest = async function(assert) {
    var message = undefined;

    let uut = new TgImgurPhotoHandler(
        mockImgur,
        true,
        mockTgBot,
        (msg) => {message = msg;}
    );

    let caption = "My Caption";

    // Pass in an empty photo array,
    await uut.RelayPhotoMessage(fromWithUserName, [photoNotInImgur], caption);

    assert.strictEqual(undefined, message);
    assert.done();
}
