/*
The MIT License (MIT)

Copyright (c) 2016 RIT Linux Users Group

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/



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
