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

const TgDocumentHandler = require("../lib/TelegramHandlers/TgDocumentHandler");

let fromWithUsername = {
    first_name : "FirstName",
    username: "username"
};

let document = {
    file_name : "filename.txt",
    mime_type : "text",
    file_size : 256,
    file_id : 13
};

let mockTgBot = {
    files : {},
    async getFileLink(id) {
        return Promise.resolve(this.files[id]);
    }
};

 /**
  * Ensures that if disabled, and no caption exists, we use the
  * default caption.
  */
exports.TgDocumentHandler_DisabledNoCaptionTest = async function(assert) {
    var message = undefined;

    let uut = new TgDocumentHandler(
        false,
        mockTgBot,
        (msg) => {message = msg;}
    );

    await uut.RelayDocumentMessage(fromWithUsername, document, null);

    var expectedMessage = '\x02' + fromWithUsername.username + '\x02' +
                " shared a file (" +
                document.mime_type +
                ") on Telegram with caption: 'Untitled Document'";

    assert.strictEqual(expectedMessage, message);
    assert.done();
}

 /**
  * Ensures that if disabled, and a caption exist, we grab the 
  * correct caption.
  */
exports.TgDocumentHandler_DisabledCaptionTest = async function(assert) {
    var message = undefined;

    let uut = new TgDocumentHandler(
        false,
        mockTgBot,
        (msg) => {message = msg;}
    );

    let caption = "My caption";

    let expectedMessage = '\x02' + fromWithUsername.username + '\x02' + 
                " shared a file (" +
                document.mime_type +
                ") on Telegram with caption: " +
                "'" + caption + "'";

    await uut.RelayDocumentMessage(fromWithUsername, document, caption);

    assert.strictEqual(expectedMessage, message);
    assert.done();
}

/**
 * Ensures that if we are enabled, but a caption does not exist, we
 * use the default caption.
 */
exports.TgDocumentHandler_EnabledNoCaptionTest = async function(assert) {
    var message = undefined;

    let expectedUrl = "https://ritlug.com/test.txt";

    mockTgBot.files[document.file_id] = expectedUrl;

    let uut = new TgDocumentHandler(
        true,
        mockTgBot,
        (msg) => {message = msg;}
    );

    await uut.RelayDocumentMessage(fromWithUsername, document, null);

    let expectedMessage = 
        "'Untitled Document' uploaded by " + 
        '\x02' + fromWithUsername.username + '\x02' + 
        ": " + 
        expectedUrl;
    
    assert.strictEqual(expectedMessage, message);
    assert.done();
};

/**
 * Ensures that if we are enabled, and a caption exists, we caption
 * the image correctly.
 */
exports.TgDocumentHandler_EnabledCaptionTest = async function(assert) {
    var message = undefined;

    let expectedUrl = "https://ritlug.com/test.txt";

    mockTgBot.files[document.file_id] = expectedUrl;

    let uut = new TgDocumentHandler(
        true,
        mockTgBot,
        (msg) => {message = msg;}
    );

    let caption = "My Caption";

    await uut.RelayDocumentMessage(fromWithUsername, document, caption);

    let expectedMessage = 
        "'" + 
        caption + 
        "' uploaded by " +
        '\x02' + fromWithUsername.username + '\x02' +
        ": " +
        expectedUrl;
    
    assert.strictEqual(expectedMessage, message);
    assert.done();
};
