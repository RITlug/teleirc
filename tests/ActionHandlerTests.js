'use strict';

const ActionHandler = require("../lib/IrcHandlers/ActionHandler");

/**
 * Ensures that if the handler is disabled,
 * nothing happens.
 */
exports.ActionHandler_DisabledTest = function(assert) {
    var message = undefined;

    let uut = new ActionHandler(undefined, false, (msg) => {message = msg;});

    uut.ReportAction("User", "#channel", "Did a thing!");

    // Disabled, message should remain undefined.
    assert.strictEqual(message, undefined);

    assert.done();
};

/**
 * Ensures that if the handler is enabled,
 * its callback is activated, even if the black-list is undefined.
 */
exports.ActionHandler_UndefinedBlackListEnabledTest = function(assert) {
    DoSuccessTest(assert, undefined);

    assert.done();
};

/**
 * Ensures that if the handler is enabled,
 * its callback is activated, even if the black-list is null.
 */
exports.ActionHandler_NullBlackListEnabledTest = function(assert) {
    DoSuccessTest(assert, null);

    assert.done();
};

/**
 * Ensures that if the handler is enabled,
 * its callback is activated, if the user is not in the black-list.
 */
exports.ActionHandler_NamesNotMatchBlackListEnabledTest = function(assert) {
    DoSuccessTest(assert, ["Someone"]);

    assert.done();
};

exports.ActionHandler_BlackListNamesMatchTest = function(assert) {
    var message = undefined;

    let uut = new ActionHandler(["user"], true, (msg) => {message = msg;});

    uut.ReportAction("User", "#channel", "Did a thing!");

    // Should be disabled, as a black-list name appeared.
    assert.strictEqual(message, undefined);

    assert.done();
};

function DoSuccessTest(assert, blackList) {
    var message = undefined;

    let uut = new ActionHandler(blackList, true, (msg) => {message = msg;});

    uut.ReportAction("User", "#channel", "Did a thing!");

    let expectedMessage = "User Did a thing!";
    assert.strictEqual(message, expectedMessage);
};
