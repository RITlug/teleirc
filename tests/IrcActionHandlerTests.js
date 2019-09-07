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
