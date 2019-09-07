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
        prefixSuffixConfig.prefix + '\x02' +
        fromNoUsername.first_name +
        '\x02' + prefixSuffixConfig.suffix +
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
        prefixSuffixConfig.prefix + '\x02' +
        fromWithUserName.username +
        '\x02' + prefixSuffixConfig.suffix +
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
