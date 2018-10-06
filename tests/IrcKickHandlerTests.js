'use strict';

const IrcKickHandler = require("../lib/IrcHandlers/IrcKickHandler");

/**
 * Ensures that if the handler is disabled,
 * nothing happens.
 */
exports.IrcKickHandler_DisabledTest = function(assert) {
    var message = undefined;

    let uut = new IrcKickHandler(false, (msg) => {message = msg;});

    uut.ReportKick("#channel", "kickeduser", "opuser", "reason");

    // Disabled, message should remain undefined.
    assert.strictEqual(message, undefined);

    assert.done();
};

/**
 * Ensures that if the handler is enabled,
 * its callback is activated.
 * This test has a reason defined.
 */
exports.IrcKickHandler_EnabledWithReasonTest = function(assert) {
    var message = undefined;

    let uut = new IrcKickHandler(true, (msg) => {message = msg;});

    uut.ReportKick("#channel", "kickeduser", "opuser", "reason");

    let expectedMessage = "kickeduser was kicked by opuser from #channel: reason.";
    assert.strictEqual(message, expectedMessage);

    assert.done();
};

/**
 * Ensures that if the handler is enabled,
 * its callback is activated.
 * This test does not have a reason defined, so it should default.
 */
exports.IrcKickHandler_EnabledWithOutReasonTest = function(assert) {
    var message = undefined;

    let uut = new IrcKickHandler(true, (msg) => {message = msg;});

    uut.ReportKick("#channel", "kickeduser", "opuser", undefined);

    let expectedMessage = "kickeduser was kicked by opuser from #channel: Kicked.";
    assert.strictEqual(message, expectedMessage);

    assert.done();
};
