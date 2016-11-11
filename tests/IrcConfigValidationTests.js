const TeleIrc = require("../lib/libteleirc.js");
const TeleIrcConfigException = require("../lib/libteleirc.js");

let sampleSettings = {
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

function shouldThrow(assert, block, errorCode) {
    assert.throws(
        block,
        function(exception) {
            return exception.errorCode === errorCode
        }
    );
}

// -------- Missing IRC Config Tests --------

/**
 * Ensures if the IRC config is missing completely from our settings, an exception gets thrown.
 */
exports.ircConfigValidation_MissingIrcConfig = function(assert) {
    // Copy our default settings, but ignore the IRC portion.
    let testSettings = {
        token: sampleSettings.token,
        ircBlacklist: sampleSettings.ircBlacklist,
        tg: sampleSettings.tg
    }

    let uut = new TeleIrc(testSettings);

    // Ensure a missing IRC config results in an exception.
    shouldThrow(
        assert,
        uut.initStage1_ircConfigValidation.bind(uut),
        uut.errorCodes.MissingIrcConfig
    );

    // Undefined should also fail.
    testSettings.irc = undefined;
    uut = new TeleIrc(testSettings);
    shouldThrow(
        assert,
        uut.initStage1_ircConfigValidation.bind(uut),
        uut.errorCodes.MissingIrcConfig
    );

    // Ditto for null.
    testSettings.irc = null;
    uut = new TeleIrc(testSettings);
    shouldThrow(
        assert,
        uut.initStage1_ircConfigValidation.bind(uut),
        uut.errorCodes.MissingIrcConfig
    );

    assert.done();
};

// -------- Required Settings Tests --------

function doRequiredSettingsTest(assert, uut, expectedErrorCode) {
    shouldThrow(
        assert,
        uut.initStage1_ircConfigValidation.bind(uut),
        expectedErrorCode
    );
}

/**
 * Ensures if the IRC config is missing its server config, an
 * exception gets thrown.
 */
exports.ircConfigValidation_missingServerConfig = function(assert) {
    // Copy our default settings, but ignore the server.
    let testSettings = {
        token: sampleSettings.token,
        irc: {
            // No Server.
            channel: sampleSettings.irc.channel,
            botName: sampleSettings.irc.botName,
            sendStickerEmoji: sampleSettings.sendStickerEmoji,
            prefix: sampleSettings.prefix,
            suffix: sampleSettings.suffix,
            showJoinMessage: sampleSettings.showJoinMessage,
            showLeaveMessage: sampleSettings.showLeaveMessage
        },
        ircBlacklist: sampleSettings.ircBlacklist,
        tg: sampleSettings.tg
    }

    // Missing server should result in exception.
    let uut = new TeleIrc(testSettings);
    doRequiredSettingsTest(assert, uut, uut.errorCodes.MissingIrcServerConfig);

    // Set server to undefined, should also cause exception.
    testSettings.server = undefined;
    uut = new TeleIrc(testSettings);
    doRequiredSettingsTest(assert, uut, uut.errorCodes.MissingIrcServerConfig);

    // Set server to null, should also cause exception.
    testSettings.server = null;
    uut = new TeleIrc(testSettings);
    doRequiredSettingsTest(assert, uut, uut.errorCodes.MissingIrcServerConfig);

    assert.done();
}

/**
 * Ensures if the IRC config is missing its channel config, an
 * exception gets thrown.
 */
exports.ircConfigValidation_missingChannelConfig = function(assert) {
    // Copy our default settings, but ignore the server.
    let testSettings = {
        token: sampleSettings.token,
        irc: {
            server : sampleSettings.irc.server,
            // No channel
            botName: sampleSettings.irc.botName,
            sendStickerEmoji: sampleSettings.sendStickerEmoji,
            prefix: sampleSettings.prefix,
            suffix: sampleSettings.suffix,
            showJoinMessage: sampleSettings.showJoinMessage,
            showLeaveMessage: sampleSettings.showLeaveMessage
        },
        ircBlacklist: sampleSettings.ircBlacklist,
        tg: sampleSettings.tg
    }

    // Missing channel should result in exception.
    let uut = new TeleIrc(testSettings);
    doRequiredSettingsTest(assert, uut, uut.errorCodes.MissingIrcChannelConfig);

    // Set channel to undefined, should also cause exception.
    testSettings.channel = undefined;
    uut = new TeleIrc(testSettings);
    doRequiredSettingsTest(assert, uut, uut.errorCodes.MissingIrcChannelConfig);

    // Set channel to null, should also cause exception.
    testSettings.channel = null;
    uut = new TeleIrc(testSettings);
    doRequiredSettingsTest(assert, uut, uut.errorCodes.MissingIrcChannelConfig);

    assert.done();
}

/**
 * Ensures if the IRC config is missing its channel config, an
 * exception gets thrown.
 */
exports.ircConfigValidation_missingBotNameConfig = function(assert) {
    // Copy our default settings, but ignore the server.
    let testSettings = {
        token: sampleSettings.token,
        irc: {
            server : sampleSettings.irc.server,
            channel: sampleSettings.irc.channel,
            // No bot name.
            sendStickerEmoji: sampleSettings.sendStickerEmoji,
            prefix: sampleSettings.prefix,
            suffix: sampleSettings.suffix,
            showJoinMessage: sampleSettings.showJoinMessage,
            showLeaveMessage: sampleSettings.showLeaveMessage
        },
        ircBlacklist: sampleSettings.ircBlacklist,
        tg: sampleSettings.tg
    }

    // Missing Bot Name should result in exception.
    let uut = new TeleIrc(testSettings);
    doRequiredSettingsTest(assert, uut, uut.errorCodes.MissingIrcBotNameConfig);

    // Set bot name to undefined, should also cause exception.
    testSettings.botName = undefined;
    uut = new TeleIrc(testSettings);
    doRequiredSettingsTest(assert, uut, uut.errorCodes.MissingIrcBotNameConfig);

    // Set bot name to null, should also cause exception.
    testSettings.botName = null;
    uut = new TeleIrc(testSettings);
    doRequiredSettingsTest(assert, uut, uut.errorCodes.MissingIrcBotNameConfig);

    assert.done();
}

// -------- Default Settings Tests --------

function doDefaultSettingsTest(assert, testSettings) {
    let uut = new TeleIrc(testSettings);

    uut.initStage1_ircConfigValidation();

    // Everything should be set to default values, and NOT undefined.

    // Undefined check:
    assert.notStrictEqual(undefined, testSettings.irc.sendStickerEmoji);
    assert.notStrictEqual(undefined, testSettings.irc.prefix);
    assert.notStrictEqual(undefined, testSettings.irc.suffix);
    assert.notStrictEqual(undefined, testSettings.irc.showJoinMessage);
    assert.notStrictEqual(undefined, testSettings.irc.showLeaveMessage);

    // Default values check:
    assert.strictEqual(uut.defaultSendStickerEmoji, testSettings.irc.sendStickerEmoji);
    assert.strictEqual(uut.defaultIrcPrefix, testSettings.irc.prefix);
    assert.strictEqual(uut.defaultIrcSuffix, testSettings.irc.suffix);
    assert.strictEqual(uut.defaultShowJoinMessage, testSettings.irc.showJoinMessage);
    assert.strictEqual(uut.defaultShowLeaveMessage, testSettings.irc.showLeaveMessage);
}

/**
 * Ensures if the IRC config is missing optional values, it gets
 * added in as a default.
 */
exports.ircConfigValidation_defaultSettingsTest_notDefined = function(assert) {
    // Copy our default settings, but ignore the IRC showEmoji.
    let testSettings = {
        token: sampleSettings.token,
        irc: {
            // These three options are required, everything else is optional.
            // If missing, we should use a default setting.
            server: sampleSettings.irc.server,
            channel: sampleSettings.irc.channel,
            botName: sampleSettings.irc.botName,
            sendStickerEmoji: undefined,
            prefix: undefined,
            suffix: undefined,
            showJoinMessage: undefined,
            showLeaveMessage: undefined
        },
        ircBlacklist: sampleSettings.ircBlacklist,
        tg: sampleSettings.tg
    }

    doDefaultSettingsTest(assert, testSettings);

    assert.done();
};

/**
 * Ensures if the IRC config is missing optional values, it gets
 * added in as a default.
 */
exports.ircConfigValidation_defaultSettingsTest_setToNull = function(assert) {
    // Copy our default settings, but ignore the IRC showEmoji.
    let testSettings = {
        token: sampleSettings.token,
        irc: {
            // These three options are required, everything else is optional.
            // If missing, we should use a default setting.
            server: sampleSettings.irc.server,
            channel: sampleSettings.irc.channel,
            botName: sampleSettings.irc.botName,
            sendStickerEmoji: null,
            prefix: null,
            suffix: null,
            showJoinMessage: null,
            showLeaveMessage: null
        },
        ircBlacklist: sampleSettings.ircBlacklist,
        tg: sampleSettings.tg
    }

    doDefaultSettingsTest(assert, testSettings);

    assert.done();
};

/**
 * Ensures if the IRC config's optional values are set to undefined, it gets
 * set to its default.
 */
exports.ircConfigValidation_defaultSettingsTest_missingSettings = function(assert) {
    // Copy our default settings, but ignore the IRC showEmoji.
    let testSettings = {
        token: sampleSettings.token,
        irc: {
            // These three options are required, everything else is optional.
            // If missing, we should use a default setting.
            server: sampleSettings.irc.server,
            channel: sampleSettings.irc.channel,
            botName: sampleSettings.irc.botName
        },
        ircBlacklist: sampleSettings.ircBlacklist,
        tg: sampleSettings.tg
    }

    doDefaultSettingsTest(assert, testSettings);

    assert.done();
};
