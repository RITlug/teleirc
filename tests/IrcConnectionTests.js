'use strict';

const nodeAssert = require('assert');
const TeleIrc = require("../lib/TeleIrc");

const TEST_NICKSERV_PASSWORD = "ThisIsATestPassword";
const TEST_NICKSERV_SERVICE = "TestNickServServiceNick";
const TEST_MESSAGE = "Test message body";
const BLACKLISTED_NICK = "Nasty one!";
const TEST_SETTINGS = {
  token: "EXAMPLE_TOKEN",
  ircBlacklist: [BLACKLISTED_NICK, "some other blacklisted nick"],
  irc: {
    server: "EXAMPLE_IRC_SERVER",
    channel: "#EXAMPLE_IRC_CHANNEL",
    botName: "",
    sendStickerEmoji: "",
    prefix: "<_",
    suffix: ">_",
    showJoinMessage: "",
    showLeaveMessage: "",
    maxMessageLength: 500,
    nickservPassword: "default value to be changed in test",
    nickservService: "default value to be changed in test",
  },
  imgur: {
    seImgurForImageLinks: true,
    imgurClientId: ""
  },
  tg: {
    chatId: "TEST_ID",
    showJoinMessage: "",
    showActionMessage: "",
    showLeaveMessage: "",
    showKickMessage: "",
    maxMessagesPerMinute: 20
  },
};

function createTgBotMock(assert) {
  return {};
}

function createIrcBotMock(assert, expectedMessages, whatSplitShouldReturn) {
  return {
    _splitLongLines: function(line, maxLen, out) {
      return whatSplitShouldReturn || [line];
    },
    connect: function(callback) {callback();},
    say: function(channel, msg) {
      assert.equal(this.expectedMessages[this.messageNo].channel, channel);
      assert.equal(this.expectedMessages[this.messageNo].message, msg);
      this.messageNo += 1;
      if (this.messageNo === this.expectedMessages.length) assert.done();
    },
    addListener: function(name, callback) {this.listeners[name] = callback;},
    messageNo: 0,
    expectedMessages: expectedMessages,
    listeners: {}
  };
}

exports.IrcConnectionTests = {
  setUp: function(callback) {
    callback();
  },
  tearDown: function(callback) {
    nodeAssert.strictEqual(this.ircBotMock.messageNo, this.ircBotMock.expectedMessages.length);
    callback();
  },
  "Connecting and identifying to NickServ as configured": function(assert) {
    const EXPECTED_MESSAGE = [
      {channel: TEST_NICKSERV_SERVICE, message: `IDENTIFY ${TEST_NICKSERV_PASSWORD}`}
    ];
    TEST_SETTINGS.irc.nickservService = TEST_NICKSERV_SERVICE;
    TEST_SETTINGS.irc.nickservPassword = TEST_NICKSERV_PASSWORD;
    let uut = new TeleIrc(TEST_SETTINGS);
    this.ircBotMock = createIrcBotMock(assert, EXPECTED_MESSAGE);
    this.tgBotMock = createTgBotMock(assert);
    uut.initStage3_initBots(this.ircBotMock, this.tgBotMock);
  },
  "Connecting and not identifying to NickServ as no password is provided": function(assert) {
    const EXPECTED_MESSAGE = [];
    TEST_SETTINGS.irc.nickservService = TEST_NICKSERV_SERVICE;
    TEST_SETTINGS.irc.nickservPassword = "";
    let uut = new TeleIrc(TEST_SETTINGS);
    this.ircBotMock = createIrcBotMock(assert, EXPECTED_MESSAGE);
    this.tgBotMock = createTgBotMock(assert);
    uut.initStage3_initBots(this.ircBotMock, this.tgBotMock);
    assert.done();
  },
  "Messages from NickServ will not be forwarded to Telegram": function(assert) {
    const EXPECTED_MESSAGE = [];
    TEST_SETTINGS.irc.nickservService = TEST_NICKSERV_SERVICE;
    TEST_SETTINGS.irc.nickservPassword = "";
    let uut = new TeleIrc(TEST_SETTINGS);
    this.ircBotMock = createIrcBotMock(assert, EXPECTED_MESSAGE);
    this.tgBotMock = createTgBotMock(assert);
    uut.initStage3_initBots(this.ircBotMock, this.tgBotMock);
    uut.initStage4_addIrcListeners();
    this.ircBotMock.listeners.message(TEST_NICKSERV_SERVICE, null, "test message");
    assert.done();
  },
  "Messages from blacklisted nicks will not be forwarded to Telegram": function(assert) {
    const EXPECTED_MESSAGE = [];
    TEST_SETTINGS.irc.nickservService = TEST_NICKSERV_SERVICE;
    TEST_SETTINGS.irc.nickservPassword = "";
    let uut = new TeleIrc(TEST_SETTINGS);
    this.ircBotMock = createIrcBotMock(assert, EXPECTED_MESSAGE);
    this.tgBotMock = createTgBotMock(assert);
    uut.initStage3_initBots(this.ircBotMock, this.tgBotMock);
    uut.initStage4_addIrcListeners();
    this.ircBotMock.listeners.message(BLACKLISTED_NICK, null, "test message");
    assert.done();
  },
  "Messages not from NickServ and not on blacklist will be forwarded to Telegram": function(assert) {
    const EXPECTED_NICK = "some other nick";
    const EXPECTED_MESSAGE = [];
    TEST_SETTINGS.irc.nickservService = TEST_NICKSERV_SERVICE;
    TEST_SETTINGS.irc.nickservPassword = "";
    let uut = new TeleIrc(TEST_SETTINGS);
    this.ircBotMock = createIrcBotMock(assert, EXPECTED_MESSAGE);
    uut.tgRateLimiter = {
      queueMessage: function(text) {
        assert.equal(text, `<${EXPECTED_NICK}> ${TEST_MESSAGE}`)
        assert.done();
      }
    };
    this.tgBotMock = createTgBotMock(assert);
    uut.initStage3_initBots(this.ircBotMock, this.tgBotMock);
    uut.initStage4_addIrcListeners();
    this.ircBotMock.listeners.message(EXPECTED_NICK, null, TEST_MESSAGE);
  },
};