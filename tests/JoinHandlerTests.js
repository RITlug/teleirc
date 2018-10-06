'use strict';

const JoinHandler = require("../lib/IrcHandlers/JoinHandler");

/**
 * Ensures that if the handler is disabled,
 * nothing happens.
 */
exports.JoinHandler_DisabledTest = function(assert) {
    var message = undefined;

    let uut = new JoinHandler(false, (msg) => {message = msg;});

    uut.ReportJoin("#channel", "user");

    // Disabled, message should remain undefined.
    assert.strictEqual(message, undefined);

    assert.done();
};

/**
 * Ensures that if the handler is enabled,
 * its callback is activated.
 */
exports.JoinHandler_EnabledTest = function(assert) {
    var message = undefined;

    let uut = new JoinHandler(true, (msg) => {message = msg;});

    uut.ReportJoin("#channel", "user");

    let expectedMessage = "user has joined #channel channel.";
    assert.strictEqual(message, expectedMessage);

    assert.done();
};