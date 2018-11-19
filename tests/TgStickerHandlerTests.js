'use strict';

const TgStickerHandler = require("../lib/TelegramHandlers/TgStickerHandler");

let fromNoUsername = {
    first_name : "FirstName",
    username : undefined
};

let fromWithUserName = {
    first_name : "FirstName",
    username : "username"
};

// Subset of the IRC config.
let prefixSuffixConfig = {
    prefix : "<",
    suffix : ">",
};

let stickerWithEmoji = {
    emoji : "🌈"
};

let stickerWithNoEmoji = {
    emoji : undefined
};

/**
 * Ensures nothing happens it the handler is disabled.
 */
exports.TgStickerHandler_DisabledTest = function(assert) {
    let message = undefined;

    let uut = new TgStickerHandler(
        prefixSuffixConfig,
        false,
        (msg) => {message = msg;}
    );

    uut.RelayStickerMessage(fromNoUsername, stickerWithEmoji);

    assert.strictEqual(undefined, message);
    assert.done();
};

/**
 * Ensures nothing happens if there is no emoji to go out.
 */
exports.TgStickerHandler_NoEmojiTest = function(assert) {
    let message = undefined;

    let uut = new TgStickerHandler(
        prefixSuffixConfig,
        false,
        (msg) => {message = msg;}
    );

    uut.RelayStickerMessage(fromNoUsername, stickerWithNoEmoji);

    assert.strictEqual(undefined, message);
    assert.done();
};

exports.TgStickerHandler_EnabledNoUserName = function(assert) {
    var message = undefined;

    let expectedMessage =
        prefixSuffixConfig.prefix +
        fromNoUsername.first_name +
        prefixSuffixConfig.suffix +
        " " +
        stickerWithEmoji.emoji;

    let uut = new TgStickerHandler(
        prefixSuffixConfig,
        true,
        (msg) => {message = msg;}
    );

    uut.RelayStickerMessage(fromNoUsername, stickerWithEmoji);

    assert.strictEqual(expectedMessage, message);
    assert.done();
};

exports.TgStickerHandler_EnabledWithUserName = function(assert) {
    var message = undefined;

    let expectedMessage =
        prefixSuffixConfig.prefix +
        fromWithUserName.username +
        prefixSuffixConfig.suffix +
        " " +
        stickerWithEmoji.emoji;

    let uut = new TgStickerHandler(
        prefixSuffixConfig,
        true,
        (msg) => {message = msg;}
    );

    uut.RelayStickerMessage(fromWithUserName, stickerWithEmoji);

    assert.strictEqual(expectedMessage, message);
    assert.done();
};
