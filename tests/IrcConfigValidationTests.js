const TeleIrc = require("../lib/libteleirc.js");
const TeleIrcConfigException = require("../lib/libteleirc.js");

var sampleSettings = {
    token: "000000000:AAAAAAaAAa2AaAAaoAAAA-a_aaAAaAaaaAA",
    ircBlacklist: [
        "CowSayBot"
    ],
    irc: {
        server: "irc.freenode.net",
        channel: "#channel",
        botName: "teleirc",
        sendStickerEmoji: false,
        prefix: "<",
        suffix: ">",
        showJoinMessage: false,
        showLeaveMessage: false
    },
    tg: {
        chatId: "-0000000000000",
        showJoinMessage: false,
        showActionMessage: true,
        showLeaveMessage: false,
        showKickMessage: false,
        maxMessagesPerMinute: 20
    }
}

/**
 * Ensures if the IRC config is missing completely from our settings, an exception gets thrown.
 */
exports.ircConfigValidation_MissingIrcConfig= function(assert) {

    // Copy our default settings, but ignore the IRC portion.
    var testSettings = {
        token: sampleSettings.token,
        ircBlacklist: sampleSettings.ircBlacklist,
        tg: sampleSettings.tg
    }

    var uut = new TeleIrc(testSettings);

    assert.throws(
        uut.initStage1_ircConfigValidation.bind(uut),
        function(exception) {
            return exception.errorCode === uut.errorCodes.MissingIrcConfig
        }
    );

    assert.done();
};