'use strict';

const IrcErrorHandler = require("../lib/IrcHandlers/IrcErrorHandler");

/**
 * Ensures that if the handler is disabled,
 * nothing happens.
 */
exports.IrcErrorHandler_DisabledTest = function(assert) {
    var message = undefined;

    let uut = new IrcErrorHandler(false, (msg) => {message = msg;});

    uut.ReportError("My Message");

    // Disabled, message should remain undefined.
    assert.strictEqual(message, undefined);

    assert.done();
};

/**
 * Ensures that if the handler is enabled,
 * its callback is activated.
 */
exports.IrcErrorHandler_EnabledTest = function(assert) {
    var message = undefined;

    let uut = new IrcErrorHandler(true, (msg) => {message = msg;});

    uut.ReportError("My Message");

    let expectedMessage = "[IRC Debug] " + JSON.stringify("My Message");
    assert.strictEqual(message, expectedMessage);

    assert.done();
};
