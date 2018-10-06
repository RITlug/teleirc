'use strict';

const IrcActionHandler = require("../lib/IrcHandlers/IrcActionHandler");

/**
 * Ensures that if the handler is disabled,
 * nothing happens.
 */
exports.IrcActionHandler_DisabledTest = function(assert) {
    var message = undefined;

    let uut = new IrcActionHandler(undefined, false, (msg) => {message = msg;});

    uut.ReportAction("User", "#channel", "Did a thing!");

    // Disabled, message should remain undefined.
    assert.strictEqual(message, undefined);

    assert.done();
};

/**
 * Ensures that if the handler is enabled,
 * its callback is activated, even if the black-list is undefined.
 */
exports.IrcActionHandler_UndefinedBlackListEnabledTest = function(assert) {
    DoSuccessTest(assert, undefined);

    assert.done();
};

/**
 * Ensures that if the handler is enabled,
 * its callback is activated, even if the black-list is null.
 */
exports.IrcActionHandler_NullBlackListEnabledTest = function(assert) {
    DoSuccessTest(assert, null);

    assert.done();
};

/**
 * Ensures that if the handler is enabled,
 * its callback is activated, if the user is not in the black-list.
 */
exports.IrcActionHandler_NamesNotMatchBlackListEnabledTest = function(assert) {
    DoSuccessTest(assert, ["Someone"]);

    assert.done();
};

exports.IrcActionHandler_BlackListNamesMatchTest = function(assert) {
    var message = undefined;

    let uut = new IrcActionHandler(["user"], true, (msg) => {message = msg;});

    uut.ReportAction("User", "#channel", "Did a thing!");

    // Should be disabled, as a black-list name appeared.
    assert.strictEqual(message, undefined);

    assert.done();
};

function DoSuccessTest(assert, blackList) {
    var message = undefined;

    let uut = new IrcActionHandler(blackList, true, (msg) => {message = msg;});

    uut.ReportAction("User", "#channel", "Did a thing!");

    let expectedMessage = "User Did a thing!";
    assert.strictEqual(message, expectedMessage);
};
