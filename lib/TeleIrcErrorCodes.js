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
