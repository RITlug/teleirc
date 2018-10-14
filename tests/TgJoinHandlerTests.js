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

    uut.ReportJoin(fromWithUserName);

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

    uut.ReportJoin(fromNoUsername);

    assert.strictEqual(expectedMessage, message);
    assert.done();
};

/**
 * Ensures that if the handler is enabled, but the user
 * has username, we report BOTH the user's first name and username.
 */
exports.TgJoinHandler_EnabledWithUsername = function(assert) {
    var message = undefined;

    let expectedMessage = fromWithUserName.first_name + " (@" + fromWithUserName.username + ") has joined the Telegram Group!";

    let uut = new TgJoinHandler(
        true,
        (msg) => {message = msg;}
    );

    uut.ReportJoin(fromWithUserName);

    assert.strictEqual(expectedMessage, message);
    assert.done();
};