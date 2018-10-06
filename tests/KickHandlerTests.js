'use strict';

const KickHandler = require("../lib/IrcHandlers/KickHandler");

/**
 * Ensures that if the handler is disabled,
 * nothing happens.
 */
exports.KickHandler_DisabledTest = function(assert) {
    var message = undefined;

    let uut = new KickHandler(false, (msg) => {message = msg;});

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
exports.KickHandler_EnabledWithReasonTest = function(assert) {
    var message = undefined;

    let uut = new KickHandler(true, (msg) => {message = msg;});

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
exports.KickHandler_EnabledWithOutReasonTest = function(assert) {
    var message = undefined;

    let uut = new KickHandler(true, (msg) => {message = msg;});

    uut.ReportKick("#channel", "kickeduser", "opuser", undefined);

    let expectedMessage = "kickeduser was kicked by opuser from #channel: Kicked.";
    assert.strictEqual(message, expectedMessage);

    assert.done();
};