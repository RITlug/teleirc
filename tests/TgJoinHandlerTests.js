'use strict';

const TgJoinHandler = require("../lib/TelegramHandlers/TgJoinHandler");

let fromNoUsername = {
    first_name : "FirstName",
    username : undefined
};

let fromWithUserName = {
    first_name : "FirstName",
    username : "username"
};

/**
 * Ensures nothing happens it the handler is disabled.
 */
exports.TgJoinHandler_DisabledTest = function(assert) {
    var message = undefined;

    let uut = new TgJoinHandler(
        false,
        (msg) => {message = msg;}
    );

    uut.RelayJoinMessage(fromWithUserName);

    assert.strictEqual(undefined, message);
    assert.done();
};

/**
 * Ensures that if the handler is enabled, but the user
 * has no username, we report the user's first name.
 */
exports.TgJoinHandler_EnabledNoUsername = function(assert) {
    var message = undefined;

    let expectedMessage = fromNoUsername.first_name + " has joined the Telegram Group!";

    let uut = new TgJoinHandler(
        true,
        (msg) => {message = msg;}
    );

    uut.RelayJoinMessage(fromNoUsername);

    assert.strictEqual(expectedMessage, message);
    assert.done();
};

/**
 * Ensures that if the handler is enabled, but the user
 * has username, we report BOTH the user's first name and username.
 */
exports.TgJoinHandler_EnabledWithUsername = function(assert) {
    var message = undefined;
    var username = "u" + "\u200B" + "sername";

    let expectedMessage = fromWithUserName.first_name + " (@" + username + ") has joined the Telegram Group!";

    let uut = new TgJoinHandler(
        true,
        (msg) => {message = msg;}
    );

    uut.RelayJoinMessage(fromWithUserName);

    assert.strictEqual(expectedMessage, message);
    assert.done();
};
