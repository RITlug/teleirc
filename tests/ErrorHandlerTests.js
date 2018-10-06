'use strict';

const ErrorHandler = require("../lib/IrcHandlers/ErrorHandler");

/**
 * Ensures that if the handler is disabled,
 * nothing happens.
 */
exports.ErrorHandler_DisabledTest = function(assert) {
    var message = undefined;

    let uut = new ErrorHandler(false, (msg) => {message = msg;});

    uut.ReportError("My Message");

    // Disabled, message should remain undefined.
    assert.strictEqual(message, undefined);

    assert.done();
};

/**
 * Ensures that if the handler is enabled,
 * its callback is activated.
 */
exports.ErrorHandler_EnabledTest = function(assert) {
    var message = undefined;

    let uut = new ErrorHandler(true, (msg) => {message = msg;});

    uut.ReportError("My Message");

    let expectedMessage = "[IRC Debug] " + JSON.stringify("My Message");
    assert.strictEqual(message, expectedMessage);

    assert.done();
};