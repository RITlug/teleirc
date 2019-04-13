'use strict';

const TeleIrc = require("../lib/TeleIrc");
const TgHelper = require("../lib/TelegramHandlers/TgHelpers.js");

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

function createIrcBotMock(assert, expectedMessages) {
  return {
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

// Override default method to not use zero width spaces
TgHelper.ResolveUserName = function(from) {
  if (from.username === undefined) {
    return from.first_name;
  }

  return from.username;
}

exports.TelegramToIrcTests = {
  "Forwarding a simple message": function(assert) {
    const USERNAME = "TEST_USERNAME";
    const EXPECTED_MESSAGE = ["<_\x02TEST_USERNAME\x02>_ TEST_MESSAGE_BODY"];
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
      "<_\x02TEST_USERNAME\x02>_ line 1",
      "<_\x02TEST_USERNAME\x02>_ line 2"
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
      "<_\x02TEST_USERNAME\x02>_ Lorem ipsum dolor sit amet, consectetur \
adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore \
magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation \
ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute \
irure dolor in reprehenderit in voluptate velit esse cillum dolore \
eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, \
sunt in culpa qui officia deserunt mollit anim id est laborum. Lorem ipsum \
dolor HERE IT SPLITS>|",
      "<_\x02TEST_USERNAME\x02>_ |<THIS WAS SPLITTED sed do eiusmod tempor \
incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, \
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo \
consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse \
cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat \
non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."
    ];

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
      text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit,\
 sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.\
 Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris\
 nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in\
 reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla\
 pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa\
 qui officia deserunt mollit anim id est laborum. Lorem ipsum dolor\
 HERE IT SPLITS>||<THIS WAS SPLITTED\
 sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.\
 Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris\
 nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in\
 reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla\
 pariatur. Excepteur sint occaecat cupidatat non proident, sunt in\
 culpa qui officia deserunt mollit anim id est laborum."
    });
  },
};
