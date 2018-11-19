'use strict';

const TgDocumentHandler = require("../lib/TelegramHandlers/TgDocumentHandler");

let from = {
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
 * Ensures nothing happens it the handler is disabled.
 */
exports.TgDocumentHandler_DisabledTest = async function(assert) {
    var message = undefined;

    let uut = new TgDocumentHandler(
        false,
        mockTgBot,
        (msg) => {message = msg;}
    );

    await uut.RelayDocumentMessage(from, document);

    assert.strictEqual(undefined, message);
    assert.done();
}

/**
 * Ensures the correct message gets generated if the handler is
 * enabled.
 */
exports.TgDocumentHandler_EnabledTest = async function(assert) {
    var message = undefined;

    let expectedUrl = "https://ritlug.com/test.txt";

    mockTgBot.files[document.file_id] = expectedUrl;

    let uut = new TgDocumentHandler(
        true,
        mockTgBot,
        (msg) => {message = msg;}
    );

    await uut.RelayDocumentMessage(from, document);

    let expectedMessage = 
        from.username + 
        " Posted File: " +
        document.file_name +
        " (" +
        document.mime_type +
        ", " +
        document.file_size +
        " bytes): " +
        expectedUrl;

    assert.strictEqual(expectedMessage, message);
    assert.done();
};
