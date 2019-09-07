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



/**
 * Wrapper Class for constants that are error codes.
 *
 * Since, you know, enums and static classes have been around for decades,
 * but Javascript doesn't believe in such things >_>.
 *
 * Main reason for dealing with error codes is for unit testing.  Why did an exception
 * get thrown?  What was the reason?  We could just ensure that when we call a method,
 * we catch the TeleIrcException... but what if something else we were NOT testing
 * threw that exception?  Our test passes since it got the correct exception,
 * but the test didn't do what it was designed to do.  With error codes, that wouldn't
 * be an issue since, yeah, we caught the exception, but the error code we
 * got was not what we expected, so we'll still fail the test.
 *
 * We could compare strings of error messages instead of codes in the unit test,
 * but what if we want to support multiple languages?
 * Numbers are universal, error messages are not.
 */
class TeleIrcErrorCodes {
  constructor() {
    // Let 1XXX be irc configuration errors.

    this.MissingIrcConfig = 1000; // IRC config is missing from the config, is undefined, or is null.
    this.MissingIrcServerConfig = 1001; // IRC config is missing its server information.
    this.MissingIrcChannelConfig = 1002; // IRC config is missing its channel information.
    this.MissingIrcBotNameConfig = 1003; // IRC config is missing its bot name.

    // Let 2XXX be Telegram Configuration errors.
  }
}
module.exports = TeleIrcErrorCodes;
