'use strict';

const PartHandler = require("../lib/IrcHandlers/PartHandler");

/**
 * Ensures that if the handler is disabled,
 * nothing happens.
 */
exports.PartHandler_DisabledTest = function(assert) {
    var message = undefined;

    let uut = new PartHandler(false, (msg) => {message = msg;});

    uut.ReportPart("#channel", "user", "reason");

    // Disabled, message should remain undefined.
    assert.strictEqual(message, undefined);

    assert.done();
};

/**
 * Ensures that if the handler is enabled,
 * its callback is activated.
 * This test has a reason defined.
 */
exports.PartHandler_EnabledWithReasonTest = function(assert) {
    var message = undefined;

    let uut = new PartHandler(true, (msg) => {message = msg;});

    uut.ReportPart("#channel", "user", "reason");

    let expectedMessage = "user has left #channel: reason.";
    assert.strictEqual(message, expectedMessage);

    assert.done();
};

/**
 * Ensures that if the handler is enabled,
 * its callback is activated.
 * This test does not have a reason defined, so it should default.
 */
exports.PartHandler_EnabledWithOutReasonTest = function(assert) {
    var message = undefined;

    let uut = new PartHandler(true, (msg) => {message = msg;});

    uut.ReportPart("#channel", "user", undefined);

    let expectedMessage = "user has left #channel: Parting...";
    assert.strictEqual(message, expectedMessage);

    assert.done();
};