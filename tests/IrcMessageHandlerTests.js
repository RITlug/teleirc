'use strict';

const IrcMessageHandler = require("../lib/IrcHandlers/IrcMessageHandler");

/**
 * Ensures that if the handler is disabled,
 * nothing happens.
 */
exports.IrcMessageHandler = {
    DisabledTest: function(assert) {
        var message = undefined;

        let uut = new IrcMessageHandler(
            undefined,
            GetMockIrcConfig(),
            false,
            (msg) => {message = msg;}
        );

        uut.RelayMessage("User", "#channel", "Hello, World!");

        // Disabled, message should remain undefined.
        assert.strictEqual(message, undefined);

        assert.done();
    },

    /**
     * Ensures that if we get a message from a channel
     * we are not in, we ignore the message.
     */
    IgnoreChannelTest: function(assert) {
        var message = undefined;

        let uut = new IrcMessageHandler(
            undefined,
            GetMockIrcConfig(),
            true,
            (msg) => {message = msg;}
        );

        uut.RelayMessage("User", "#badchannel", "Hello, World!");

        // Not the channel we are in, ignore.
        assert.strictEqual(message, undefined);

        assert.done();
    },

    /**
     * Ensures that if the handler is enabled,
     * its callback is activated, even if the black-list is undefined.
     */
    UndefinedBlackListEnabledTest: function(assert) {
        DoSuccessTest(assert, undefined);

        assert.done();
    },

    /**
     * Ensures that if the handler is enabled,
     * its callback is activated, even if the black-list is null.
     */
    NullBlackListEnabledTest: function(assert) {
        DoSuccessTest(assert, null);

        assert.done();
    },

    /**
     * Ensures that if the handler is enabled,
     * its callback is activated, if the user is not in the black-list.
     */
    NamesNotMatchBlackListEnabledTest: function(assert) {
        DoSuccessTest(assert, ["Someone"]);

        assert.done();
    },

    BlackListNamesMatchTest: function(assert) {
        var message = undefined;

        let uut = new IrcMessageHandler(
            ["user"],
            GetMockIrcConfig(),
            true,
            (msg) => {message = msg;}
        );

        uut.RelayMessage("User", "#channel", "Hello, World!");

        // Should be disabled, as a black-list name appeared.
        assert.strictEqual(message, undefined);

        assert.done();
    }
};

function DoSuccessTest(assert, blackList) {
    var message = undefined;

    let uut = new IrcMessageHandler(
        blackList,
        GetMockIrcConfig(),
        true,
        (msg) => {message = msg;}
    );

    uut.RelayMessage("User", "#channel", "Hello, World!");

    let expectedMessage = "<*User*> Hello, World!";
    assert.strictEqual(message, expectedMessage);
};

function GetMockIrcConfig() {
    return {
        channel : "#channel"
    };
}
