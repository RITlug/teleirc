'use strict';

const IrcJoinHandler = require("../lib/IrcHandlers/IrcJoinHandler");

class ActionClass {
    constructor(action) {
        this._action = action;
    }

    DoAction(message) {
        this._action(message);
    }
}

/**
 * Ensures that we can pass in a function pointer to the class just fine.
 */
exports.IrcJoinHandler_FunctionPointerClass = function(assert) {
    var message = undefined;

    let actionClass = new ActionClass( (msg) => {message=msg;});
    let uut = new IrcJoinHandler(true, actionClass.DoAction.bind(actionClass));

    uut.ReportJoin("#channel", "user");

    // Disabled, message should remain undefined.
    assert.strictEqual(message, "user has joined #channel.");

    assert.done();
}

/**
 * Ensures that if the handler is disabled,
 * nothing happens.
 */
exports.IrcJoinHandler_DisabledTest = function(assert) {
    var message = undefined;

    let uut = new IrcJoinHandler(false, (msg) => {message = msg;});

    uut.ReportJoin("#channel", "user");

    // Disabled, message should remain undefined.
    assert.strictEqual(message, undefined);

    assert.done();
};

/**
 * Ensures that if the handler is enabled,
 * its callback is activated.
 */
exports.IrcJoinHandler_EnabledTest = function(assert) {
    var message = undefined;

    let uut = new IrcJoinHandler(true, (msg) => {message = msg;});

    uut.ReportJoin("#channel", "user");

    let expectedMessage = "user has joined #channel.";
    assert.strictEqual(message, expectedMessage);

    assert.done();
};
