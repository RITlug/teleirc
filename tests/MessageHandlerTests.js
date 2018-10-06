'use strict';

const MessageHandler = require("../lib/IrcHandlers/MessageHandler");

/**
 * Ensures that if the handler is disabled,
 * nothing happens.
 */
exports.MessageHandler_DisabledTest = function(assert) {
    var message = undefined;

    let uut = new MessageHandler(undefined, false, (msg) => {message = msg;});

    uut.ReportMessage("User", "#channel", "Hello, World!");

    // Disabled, message should remain undefined.
    assert.strictEqual(message, undefined);

    assert.done();
};

/**
 * Ensures that if the handler is enabled,
 * its callback is activated, even if the black-list is undefined.
 */
exports.MessageHandler_UndefinedBlackListEnabledTest = function(assert) {
    DoSuccessTest(assert, undefined);

    assert.done();
};

/**
 * Ensures that if the handler is enabled,
 * its callback is activated, even if the black-list is null.
 */
exports.MessageHandler_NullBlackListEnabledTest = function(assert) {
    DoSuccessTest(assert, null);

    assert.done();
};

/**
 * Ensures that if the handler is enabled,
 * its callback is activated, if the user is not in the black-list.
 */
exports.MessageHandler_NamesNotMatchBlackListEnabledTest = function(assert) {
    DoSuccessTest(assert, ["Someone"]);

    assert.done();
};

exports.MessageHandler_BlackListNamesMatchTest = function(assert) {
    var message = undefined;

    let uut = new MessageHandler(["user"], true, (msg) => {message = msg;});

    uut.ReportMessage("User", "#channel", "Hello, World!");

    // Should be disabled, as a black-list name appeared.
    assert.strictEqual(message, undefined);

    assert.done();
};

function DoSuccessTest(assert, blackList) {
    var message = undefined;

    let uut = new MessageHandler(blackList, true, (msg) => {message = msg;});

    uut.ReportMessage("User", "#channel", "Hello, World!");

    let expectedMessage = "<User> Hello, World!";
    assert.strictEqual(message, expectedMessage);
};
