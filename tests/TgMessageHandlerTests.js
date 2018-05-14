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

        let expectedMessage = "My Message";
        let expectedUsername = fromWithUserName.first_name;

        let uut = new TgMessageHandler(
            prefixSuffixConfig,
            true,
            (actualUsername, actualMessage) => {
                assert.strictEqual(expectedMessage, actualMessage);
                assert.strictEqual(expectedUsername, actualUsername);
                assert.done();
            }
        );

        uut.RelayMessage(fromNoUsername, expectedMessage);
    },
    "Username is reported if it is available": function(assert) {
        var actualMessage = undefined;
        var actualUsername = undefined;

        let expectedMessage = "My Message";
        let expectedUsername = fromWithUserName.username;

        let uut = new TgMessageHandler(
            prefixSuffixConfig,
            true,
            (actualUsername, actualMessage) => {
                assert.strictEqual(expectedMessage, actualMessage);
                assert.strictEqual(expectedUsername, actualUsername);
                assert.done();
            }
        );

        uut.RelayMessage(fromWithUserName, expectedMessage);
    },
};