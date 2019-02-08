'use strict';

const dotenv = require('dotenv').config();
const TeleIrc = require("../lib/TeleIrc");
const TeleIrcConfigException = require("../lib/TeleIrcException");
const config = require('../config');

function shouldThrow(assert, block, errorCode) {
  assert.throws(
    block,
    function(exception) {
      let retval = (exception.errorCode === errorCode);
      if (!retval) {
        console.log(JSON.stringify(exception));
        console.trace();
      }
      return retval;
    }
  );
}

// -------- Missing IRC Config Tests --------

exports.ircConfigValidation = {};

/**
 * Ensures if the IRC config is missing completely from our settings, an exception gets thrown.
 */
exports.ircConfigValidation.MissingIrcConfig = function(assert) {
  // Copy our default settings, but ignore the IRC portion.
  let testSettings = {
    token: config.token,
    ircBlacklist: config.ircBlacklist,
    tg: config.tg,
    imgur: config.imgur
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
exports.ircConfigValidation["Missing IRC server config causes an exception"] = function(assert) {
  // Copy our default settings, but ignore the server.
  let testSettings = {
    token: config.token,
    irc: {
      // No Server.
      channel: "test channel",
      botName: "test bot name",
      sendStickerEmoji: config.sendStickerEmoji,
      prefix: config.prefix,
      suffix: config.suffix,
      showJoinMessage: config.showJoinMessage,
      showLeaveMessage: config.showLeaveMessage
    },
    ircBlacklist: config.ircBlacklist,
    tg: config.tg,
    imgur: config.imgur
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

exports.ircConfigValidation["Missing IRC channel config causes an exception"] = function(assert) {
  // Copy our default settings, but ignore the server.
  let testSettings = {
    token: config.token,
    irc: {
      server: "test irc server",
      // No channel
      botName: "test bot name",
      sendStickerEmoji: config.sendStickerEmoji,
      prefix: config.prefix,
      suffix: config.suffix,
      showJoinMessage: config.showJoinMessage,
      showLeaveMessage: config.showLeaveMessage
    },
    ircBlacklist: config.ircBlacklist,
    tg: config.tg,
    imgur: config.imgur
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

exports.ircConfigValidation["Missing IRC bot name causes an exception"] = function(assert) {
  // Copy our default settings, but ignore the server.
  let testSettings = {
    token: config.token,
    irc: {
      server: "test irc server",
      channel: "test channel",
      // No bot name.
      sendStickerEmoji: config.sendStickerEmoji,
      prefix: config.prefix,
      suffix: config.suffix,
      showJoinMessage: config.showJoinMessage,
      showLeaveMessage: config.showLeaveMessage,
    },
    ircBlacklist: config.ircBlacklist,
    tg: config.tg,
    imgur: config.imgur
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
exports.ircConfigValidation.defaultSettingsTest_notDefined = function(assert) {
  // Copy our default settings, but ignore the IRC showEmoji.
  let testSettings = {
    token: config.token,
    irc: {
      // These three options are required, everything else is optional.
      // If missing, we should use a default setting.
      server: "test irc server",
      channel: "test channel",
      botName: "test bot name",
      sendStickerEmoji: undefined,
      prefix: undefined,
      suffix: undefined,
      showJoinMessage: undefined,
      showLeaveMessage: undefined
    },
    ircBlacklist: config.ircBlacklist,
    tg: config.tg,
    imgur: config.imgur
  }

  doDefaultSettingsTest(assert, testSettings);

  assert.done();
};

/**
 * Ensures if the IRC config is missing optional values, it gets
 * added in as a default.
 */
exports.ircConfigValidation.defaultSettingsTest_setToNull = function(assert) {
  // Copy our default settings, but ignore the IRC showEmoji.
  let testSettings = {
    token: config.token,
    irc: {
      // These three options are required, everything else is optional.
      // If missing, we should use a default setting.
      server: "test irc server",
      channel: "test channel",
      botName: "test bot name",
      sendStickerEmoji: null,
      prefix: null,
      suffix: null,
      showJoinMessage: null,
      showLeaveMessage: null
    },
    ircBlacklist: config.ircBlacklist,
    tg: config.tg,
    imgur: config.imgur
  }

  doDefaultSettingsTest(assert, testSettings);

  assert.done();
};

/**
 * Ensures if the IRC config's optional values are set to undefined, it gets
 * set to its default.
 */
exports.ircConfigValidation.defaultSettingsTest_missingSettings = function(
  assert) {
  // Copy our default settings, but ignore the IRC showEmoji.
  let testSettings = {
    token: config.token,
    irc: {
      // These three options are required, everything else is optional.
      // If missing, we should use a default setting.
      server: "test irc server",
      channel: "test channel",
      botName: "test bot name"
    },
    ircBlacklist: config.ircBlacklist,
    tg: config.tg,
    imgur: config.imgur
  }

  doDefaultSettingsTest(assert, testSettings);

  assert.done();
};

exports.ircConfigValidation.checkTypes = (assert) => {
  assert.strictEqual(typeof config.irc.sendStickerEmoji, 'boolean');
  assert.strictEqual(typeof config.irc.showJoinMessage, 'boolean');
  assert.strictEqual(typeof config.irc.showLeaveMessage, 'boolean');
  assert.strictEqual(typeof config.tg.showJoinMessage, 'boolean');
  assert.strictEqual(typeof config.tg.showLeaveMessage, 'boolean');
  assert.strictEqual(typeof config.tg.showKickMessage, 'boolean');
  assert.strictEqual(typeof config.tg.maxMessagesPerMinute, 'number');
  assert.strictEqual(typeof config.imgur.useImgurForImageLinks, 'boolean');
  assert.done();

}
