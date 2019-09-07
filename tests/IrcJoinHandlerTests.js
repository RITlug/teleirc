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
