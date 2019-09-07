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
