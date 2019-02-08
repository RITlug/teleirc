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
        var actualMessage = undefined;
        var actualUsername = undefined;

        let sentMessage = "My Message";
        let expectedUsername = fromWithUserName.first_name;
        let expectedMessage = `<${expectedUsername}> My Message`;

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
        var actualMessage = undefined;
        var actualUsername = undefined;

        let sentMessage = "My Message";
        let expectedUsername = fromWithUserName.username;
        let expectedMessage = `<${expectedUsername}> My Message`;

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
};