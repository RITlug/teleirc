'use strict';

const TeleIrc = require("../lib/TeleIrc");

const TEST_SETTINGS = {
  token: "EXAMPLE_TOKEN",
  ircBlacklist: "",
  irc: {
    server: "EXAMPLE_IRC_SERVER",
    channel: "#EXAMPLE_IRC_CHANNEL",
    botName: "",
    sendStickerEmoji: "",
    prefix: "<_",
    suffix: ">_",
    showJoinMessage: "",
    showLeaveMessage: "",
    maxMessageLength: 500
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

function createTgBotMock() {
  return {
    on: function(messageName, callback) {
      this.callbacks[messageName] = callback;
    },
    callbacks: {},
    getMe: async function() {
      return Promise.resolve({ username:"test" });
    }
  };
}

function createIrcBotMock(assert, expectedMessages, whatSplitShouldReturn) {
  return {
    _splitLongLines: function(line, maxLen, out) {
      return whatSplitShouldReturn || [line];
    },
    say: function(channel, msg) {
      assert.equal(TEST_SETTINGS.irc.channel, channel);
      assert.equal(this.expectedMessages[this.messageNo], msg);
      this.messageNo += 1;
      if (this.messageNo === this.expectedMessages.length) assert.done();
    },
    messageNo: 0,
    expectedMessages: expectedMessages,
    connect: function() {}
  };
}

exports.TelegramToIrcTests = {
  "Forwarding a simple message": function(assert) {
    const USERNAME = "TEST_USERNAME";
    const EXPECTED_MESSAGE = ["<_TEST_USERNAME>_ TEST_MESSAGE_BODY"];
    const MESSAGE_FROM_TELEGRAM = "TEST_MESSAGE_BODY";

    let uut = new TeleIrc(TEST_SETTINGS);
    let ircBotMock = createIrcBotMock(assert, EXPECTED_MESSAGE);
    let tgBotMock = createTgBotMock();
    uut.initStage3_initBots(ircBotMock, tgBotMock);
    uut.initStage5_initTelegramMessageSending();

    tgBotMock.callbacks["message"]({
      chat: {
        id: TEST_SETTINGS.tg.chatId
      },
      from: {
        username: USERNAME
      },
      text: MESSAGE_FROM_TELEGRAM
    });
  },
  "Forwarding a multiline message": function(assert) {
    const USERNAME = "TEST_USERNAME";
    const EXPECTED_MESSAGES = [
      "<_TEST_USERNAME>_ line 1",
      "<_TEST_USERNAME>_ line 2"
    ];
    const MESSAGE_FROM_TELEGRAM = "line 1\n\nline 2";

    let uut = new TeleIrc(TEST_SETTINGS);
    let ircBotMock = createIrcBotMock(assert, EXPECTED_MESSAGES);
    let tgBotMock = createTgBotMock();
    uut.initStage3_initBots(ircBotMock, tgBotMock);
    uut.initStage5_initTelegramMessageSending();

    tgBotMock.callbacks["message"]({
      chat: {
        id: TEST_SETTINGS.tg.chatId
      },
      from: {
        username: USERNAME
      },
      text: MESSAGE_FROM_TELEGRAM
    });
  },
  "Forwarding a long message": function(assert) {
    const USERNAME = "TEST_USERNAME";
    const EXPECTED_MESSAGES = [
      "<_TEST_USERNAME>_ line 1",
      "<_TEST_USERNAME>_ line 2"
    ];
    const SPLITTED_MESSAGE = ["line 1", "line 2"];

    let uut = new TeleIrc(TEST_SETTINGS);
    let ircBotMock = createIrcBotMock(assert, EXPECTED_MESSAGES, SPLITTED_MESSAGE);
    let tgBotMock = createTgBotMock();
    uut.initStage3_initBots(ircBotMock, tgBotMock);
    uut.initStage5_initTelegramMessageSending();

    tgBotMock.callbacks["message"]({
      chat: {
        id: TEST_SETTINGS.tg.chatId
      },
      from: {
        username: USERNAME
      },
      text: "EXAMPLE"
    });
  },
};