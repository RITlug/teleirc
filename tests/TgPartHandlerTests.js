'use strict';

const TgPartHandler = require("../lib/TelegramHandlers/TgPartHandler");

let fromNoUsername = {
    first_name : "FirstName",
    username : undefined
};

let fromWithUserName = {
    first_name : "FirstName",
    username : "username"
};

exports.TgPartHandler_DisabledTest = function(assert) {
    var message = undefined;

    let uut = new TgPartHandler(
        false,
        (msg) => {message = msg;}
    );

    uut.RelayPartMessage(fromWithUserName);

    assert.strictEqual(undefined, message);
    assert.done();
};

/**
 * Ensures that if the handler is enabled, but the user
 * has no username, we report the user's first name.
 */
exports.TgPartHandler_EnabledNoUsername = function(assert) {
    var message = undefined;

    let expectedMessage = fromNoUsername.first_name + " has left the Telegram Group.";

    let uut = new TgPartHandler(
        true,
        (msg) => {message = msg;}
    );

    uut.RelayPartMessage(fromNoUsername);

    assert.strictEqual(expectedMessage, message);
    assert.done();
};

/**
 * Ensures that if the handler is enabled, but the user
 * has username, we report BOTH the user's first name and username.
 */
exports.TgPartHandler_EnabledWithUsername = function(assert) {
    var message = undefined;
    var username = "u" + "\u200B" + "sername";

    let expectedMessage = fromWithUserName.first_name + " (@" + username + ") has left the Telegram Group.";

    let uut = new TgPartHandler(
        true,
        (msg) => {message = msg;}
    );

    uut.RelayPartMessage(fromWithUserName);

    assert.strictEqual(expectedMessage, message);
    assert.done();
};
