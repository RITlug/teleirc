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

const TgMessageHandler = require("../lib/TelegramHandlers/TgMessageHandler");

let fromNoUsername = {
    first_name : "FirstName",
    username : undefined
};

let fromWithUserName = {
    first_name : "FirstName",
    username : "username"
};

let fromWithZWP = {
    first_name : "First Name",
    username : "u" + "\u200B" + "sername"
}

// Subset of the IRC config.
let prefixSuffixConfig = {
    prefix : "<",
    suffix : ">",
    maxMessageLength: 400,
};

/**
 * Ensures nothing happens it the handler is disabled.
 */
exports.TgMessageHandler = {
    "Nothing happens when disabled": function(assert) {
        var message = undefined;

        let uut = new TgMessageHandler(
            prefixSuffixConfig,
            false,
            (uname, msg) => {
                assert.fail("Messages should not be forwarded when disabled!");
            }
        );

        uut.RelayMessage(fromNoUsername, "My Message");

        assert.strictEqual(undefined, message);
        assert.done();
    },
    "If user has no username, first name is reported": function(assert) {

        let sentMessage = "My Message";
        let expectedUsername = fromWithUserName.first_name;
        let expectedMessage = `<\x02${expectedUsername}\x02> My Message`;

        let uut = new TgMessageHandler(
            prefixSuffixConfig,
            true,
            (input) => {
                assert.strictEqual(expectedMessage, input);
                assert.done();
            }
        );

        uut.RelayMessage(fromNoUsername, sentMessage);
    },
    "Username is reported if it is available": function(assert) {

        let sentMessage = "My Message";
        let expectedUsername = fromWithUserName.username;
        let expectedMessage = `<\x02${expectedUsername}\x02> My Message`;

        let uut = new TgMessageHandler(
            prefixSuffixConfig,
            true,
            (input) => {
                assert.strictEqual(expectedMessage, input);
                assert.done();
            }
        );

        uut.RelayMessage(fromWithUserName, sentMessage);
    },
    "Zero-width space is inserted into username": function(assert) {
        let sentMessage = "My Message";
        let expectedUsername = fromWithZWP.username;
        let expectedMessage = `<\x02${expectedUsername}\x02> My Message`;

        let uut = new TgMessageHandler (
            prefixSuffixConfig,
            true,
            (input) => {
                assert.strictEqual(expectedMessage, input);
                assert.done();
            }
        );

        uut.RelayMessage(fromWithZWP, sentMessage);
    }
};
