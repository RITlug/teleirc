'use strict';

const TgMessageHandler = require("../lib/TelegramHandlers/TgMessageHandler");

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

/**
 * Ensures nothing happens it the handler is disabled.
 */
exports.TgMessageHandler_DisabledTest = function(assert) {
    var message = undefined;

    let uut = new TgMessageHandler(
        prefixSuffixConfig,
        false,
        (msg) => {message = msg;}
    );

    uut.RelayMessage(fromNoUsername, "My Message");

    assert.strictEqual(undefined, message);
    assert.done();
};

/**
 * Ensures that if the handler is enabled, but the user
 * has no username, we report the user's first name.
 */
exports.TgMessageHandler_EnabledNoUserName = function(assert) {
    var message = undefined;

    let userMessage = "My Message";

    let expectedMessage =
        prefixSuffixConfig.prefix +
        fromNoUsername.first_name +
        prefixSuffixConfig.suffix +
        " " +
        userMessage;

    let uut = new TgMessageHandler(
        prefixSuffixConfig,
        true,
        (msg) => {message = msg;}
    );

    uut.RelayMessage(fromNoUsername, userMessage);

    assert.strictEqual(expectedMessage, message);
    assert.done();
};

/**
 * Ensures that if the handler is enabled, but the user
 * has username, we report just the user's username.
 */
exports.TgMessageHandler_EnabledWithUserName = function(assert) {
    var message = undefined;

    let userMessage = "My Message";

    let expectedMessage =
        prefixSuffixConfig.prefix +
        fromWithUserName.username +
        prefixSuffixConfig.suffix +
        " " +
        userMessage;

    let uut = new TgMessageHandler(
        prefixSuffixConfig,
        true,
        (msg) => {message = msg;}
    );

    uut.RelayMessage(fromWithUserName, userMessage);

    assert.strictEqual(expectedMessage, message);
    assert.done();
};
