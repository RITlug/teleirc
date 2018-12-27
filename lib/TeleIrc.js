const IrcActionHandler = require('./IrcHandlers/IrcActionHandler');
const IrcErrorHandler = require('./IrcHandlers/IrcErrorHandler');
const IrcJoinHandler = require('./IrcHandlers/IrcJoinHandler');
const IrcKickHandler = require('./IrcHandlers/IrcKickHandler');
const IrcMessageHandler = require('./IrcHandlers/IrcMessageHandler');
const IrcNickServHandler = require('./IrcHandlers/IrcNickServHandler');
const IrcPartHandler = require('./IrcHandlers/IrcPartHandler');

const TgDocumentHandler = require('./TelegramHandlers/TgDocumentHandler.js');
const TgErrorHandler = require('./TelegramHandlers/TgErrorHandler.js');
const TgEventListener = require('./TelegramHandlers/TgEventListener.js');
const TgJoinHandler = require('./TelegramHandlers/TgJoinHandler.js');
const TgMessageHandler = require('./TelegramHandlers/TgMessageHandler.js');
const TgPartHandler = require('./TelegramHandlers/TgPartHandler.js');
const TgPhotoHandler = require('./TelegramHandlers/TgPhotoHandler.js');
const TgStickerHandler = require('./TelegramHandlers/TgStickerHandler.js');
const TgImgurPhotoHandler = require('./TelegramHandlers/TgImgurPhotoHandler.js');

const MessageRateLimiter = require("./MessageRateLimiter");
const TeleIrcException = require('./TeleIrcException');
const TeleIrcErrorCodes = require('./TeleIrcErrorCodes');

class TeleIrc {

  // ---------------- Constructor ----------------

  /**
   * @param teleIrcConfig - The teleirc config to use.
   */
  constructor(teleIrcConfig) {
    this.config = teleIrcConfig;
    // Workaround since apparently javascript doesn't believe in static
    // variables.  Bleh.
    if (this.errorCodes === undefined) {
      this.errorCodes = new TeleIrcErrorCodes();
    }

    // "Constants:"
    this.defaultSendStickerEmoji = false;
    this.defaultIrcPrefix = "";
    this.defaultIrcSuffix = "";
    this.defaultShowJoinMessage = false;
    this.defaultShowLeaveMessage = false;
  }

  // ---------------- Functions ----------------

  // -------- Init Functions --------

  // ---- Config Validation ----

  /**
   * Stage one of initialization:
   * Make sure the IRC config we passed in is valid.
   * Any option that is missing from the config that is not required
   * gets a default value.
   *
   * If any option that is REQUIRED from the config and is missing results
   * in an exception being thrown.
   * Required settings are:
   *  - server
   *  - channel
   *  - botName
   */
  initStage1_ircConfigValidation() {
    let ircooptions = ["server", "channel", "botName"];
    if (this.config.irc) {
      ircooptions.forEach(option => {
        if (!this.config.irc[option]) {
          let errorCode = 0;
          if (option === "server") {
            errorCode = this.errorCodes.MissingIrcServerConfig;
          } else if (option === "channel") {
            errorCode = this.errorCodes.MissingIrcChannelConfig;
          } else if (option === "botName") {
            errorCode = this.errorCodes.MissingIrcBotNameConfig;
          }
          throw new TeleIrcException(errorCode,
            "Unable to find IRC settings for " + option);
        }
      });

      // The following settings are optional. If there are no
      // options set for them, set default values.
      this.config.irc.sendStickerEmoji = this.checkConfigOption(
        this.config.irc.sendStickerEmoji,
        this.defaultSendStickerEmoji,
        "sendStickerEmoji"
      );

      this.config.irc.prefix = this.checkConfigOption(
        this.config.irc.prefix,
        this.defaultIrcPrefix,
        "prefix"
      );

      this.config.irc.suffix = this.checkConfigOption(
        this.config.irc.suffix,
        this.defaultIrcSuffix,
        "suffix"
      );

      this.config.irc.showJoinMessage = this.checkConfigOption(
        this.config.irc.showJoinMessage,
        this.defaultShowJoinMessage,
        "showJoinMessage"
      );

      this.config.irc.showLeaveMessage = this.checkConfigOption(
        this.config.irc.showLeaveMessage,
        this.defaultShowLeaveMessage,
        "showLeaveMessage"
      );
    } else {
      throw new TeleIrcException(
        this.errorCodes.MissingIrcConfig,
        "Unable to find IRC settings in config.js"
      );
    }
  }

  /**
   * Stage two of initialization:
   * Make sure the Telegram config we passed in is valid.
   * Any option that is missing from the config that is not required
   * gets a default value.
   *
   * If any option that is REQUIRED from the config and is missing results
   * in an exception being thrown.
   * Required settings are:
   *  - token
   *  - chatId
   */
  initStage2_telegramConfigValidation() {
    if (!this.config.token) {
      throw "Unable to find a telegram bot token in config.js"
    }

    // Check for required Telegram settings:
    let tgoptions = ["chatId"];
    if (this.config.tg) {
      tgoptions.forEach(option => {
        if (!this.config.tg[option]) {
          throw "Unable to find Telegram settings for " + option;
        }
      });

      // The following settings are optional. If there are no
      // options set for them, default to false.

      // Read config file for showJoinMessage
      if (!this.config.tg.hasOwnProperty("showJoinMessage")) {
        console.log("Using default 'false' value for showJoinMessage");
        this.config.tg.showJoinMessage = false;
      }

      // Read config file for showLeaveMessage
      if (!this.config.tg.hasOwnProperty("showLeaveMessage")) {
        console.log("Using default 'false' value for showLeaveMessage");
        this.config.tg.showLeaveMessage = false;
      }

      // Read config file for showKickMessage
      if (!this.config.tg.hasOwnProperty("showKickMessage")) {
        console.log("Using default 'false' value for showKickMessage");
        this.config.tg.showKickMessage = false;
      }

      if (!this.config.tg.hasOwnProperty("showActionMessage")) {
        console.log("Using default 'false' value for showActionMessage");
        this.config.tg.showActionMessage = false;
      }

      if (!this.config.tg.hasOwnProperty("maxMessagesPerMinute")) {
        console.log("Using default of 20 for maxMessagesPerMinute");
        this.config.tg.maxMessagesPerMinute = 20;
      }

    } else {
      throw "Unable to find Telegram settings in config.js";
    }
  }

  // ---- Telegram/IRC bot construction ----

  /**
   * After initializing the bots outside of this class,
   * this gives this class a reference to those.
   * @param ircbot - The IRC bot to use.
   * @param tgbot - The Telegram bot to use.
   */
  initStage3_initBots(ircbot, tgbot) {
    this.ircbot = ircbot;
    this.tgbot = tgbot;
    // This is super rudimentary error handling. If the Telegram bot
    // throws an error, the library we use simply logs it and appears
    // to hang. This may be a bug in that library (see
    // https://github.com/yagop/node-telegram-bot-api/issues/657).
    // This code forces teleirc to fully crash when it encounters a polling
    // error so that its parent process/container can restart it.
    // It may be possible to handle this in a nicer way, but issues
    // open on the bot library indicate that things may not be working right
    // as of October 2018.
    this.tgbot.on('polling_error', (error) => {
        console.log("Fatal Telegram polling error: " + error.code);
        console.log(error.stack);
        process.exit(1);
    });
    this.tgbot.on('webhook_error', (error) => {
        console.log("Fatal Telegram webhook error: " + error.code);
        console.log(error.stack);
        process.exit(1);
    });

    // Verifies that we can, in fact, reach the Telegram API and
    // are not being rate limited.
    this.tgbot.getMe().then(function(result) {
        console.log("Connected to Telegram. Our bot name is " + result.username);
    })
    .catch(function(result) {
        console.log("Unable to reach Telegram API. Check that you're not being rate limited.");
    });

    this.ircbot.connect(() => {
      this.nickservIdentify();
    });
  }

  /**
   * Identifies to NickServ, if password is provided.
   */
  nickservIdentify() {
    if ((this.config.irc.nickservPassword.length > 0) &&
        (this.config.irc.nickservService.length > 0)) {
      console.log("Identifying to NickServ...");
      var nickservRegisteredNick = this.config.irc.botName;
      if (
        this.config.irc.nickservRegisteredNick &&
        (this.config.irc.nickservRegisteredNick.length > 0)
      ) {
        nickservRegisteredNick = this.config.irc.nickservRegisteredNick;
      }
      this.ircbot.say(
        this.config.irc.nickservService,
        "IDENTIFY " + nickservRegisteredNick + " " + this.config.irc.nickservPassword
      );
    } else {
      console.log("NickServ password not provided - omitting indentification");
    }
  }

  /**
   * Stage 4 of initialization:
   * Add all the IRC bot's listeners based on this object's config.
   */
  initStage4_addIrcListeners() {
    this.messageHandler = new IrcMessageHandler(
      this.config.ircBlacklist,
      this.config.irc,
      true,
      this.queueTelegramMessage.bind(this)
    );
    this.ircbot.addListener('message', this.messageHandler.RelayMessage.bind(this.messageHandler));

    this.nickServHandler = new IrcNickServHandler(this.config.irc, this.ircbot);
    this.ircbot.addListener('message', this.nickServHandler.HandleNickServ.bind(this.nickServHandler));

    this.errorHandler = new IrcErrorHandler(
      true,
      this.status.bind(this)
    );
    this.ircbot.addListener('error', this.errorHandler.ReportError.bind(this.errorHandler));

    this.actionHandler = new IrcActionHandler(
      this.config.ircBlacklist,
      this.config.tg.showActionMessage,
      this.queueTelegramMessage.bind(this)
    );
    this.ircbot.addListener('action', this.actionHandler.ReportAction.bind(this.actionHandler));

    this.joinHandler = new IrcJoinHandler(
      this.config.tg.showJoinMessage,
      this.queueTelegramMessage.bind(this)
    );
    this.ircbot.addListener('join', this.joinHandler.ReportJoin.bind(this.joinHandler));

    this.partHandler = new IrcPartHandler(
      this.config.tg.showLeaveMessage,
      this.queueTelegramMessage.bind(this)
    );
    this.ircbot.addListener('part', this.partHandler.ReportPart.bind(this.partHandler));

    this.kickHandler = new IrcKickHandler(
      this.config.tg.showKickMessage,
      this.queueTelegramMessage.bind(this)
    );
    this.ircbot.addListener('kick', this.kickHandler.ReportKick.bind(this.kickHandler));
  }

  /**
   * Stage 5 of initialization:
   * The telegram rate-limiter and telegram bot gets intialized.
   */
  initStage5_initTelegramMessageSending() {
    // In node, inline functions are a tad different from other languages such
    // as C#.  If we were to pass in this.tgbot.sendMessage instead of
    // teleirc.tgbot.sendMessage, tgbot would be called on the instance of the telegram
    // bot, not in this class.  Therefore, copy a reference to this class
    // to a local variable, and call that local variable in the inline function
    // to the telegram bot calls the right class.
    this.tgEventListener = new TgEventListener({
      chatId: this.config.tg.chatId,
      ircEditedPrefix: this.config.irc.editedPrefix,
    });

    let teleirc = this;
    this.tgRateLimiter = new MessageRateLimiter(
      this.config.tg.maxMessagesPerMinute,
      60,
      (message) => {
        teleirc.tgbot.sendMessage(teleirc.config.tg.chatId, message);
      });

    this.tgbot.on('message', teleirc.tgEventListener.ParseMessage.bind(teleirc.tgEventListener));
    this.tgbot.on('edited_message', teleirc.tgEventListener.ParseEditedMessage.bind(teleirc.tgEventListener));

    // ---- Setup Telegram Handlers ----

    this.tgDocumentHandler = new TgDocumentHandler(
      true,
      this.tgbot,
      this.sendMessageToIrc.bind(this)
    );
    this.tgEventListener.addListener(
      'document',
      this.tgDocumentHandler.RelayDocumentMessage.bind(this.tgDocumentHandler)
    );

    this.tgErrorHandler = new TgErrorHandler(
      true,
      this.status.bind(this)
    );
    this.tgEventListener.addListener(
      'bad_chat_id',
      this.tgErrorHandler.HandleBadChatId.bind(this.tgErrorHandler)
    );
    this.tgEventListener.addListener(
      'unknown',
      this.tgErrorHandler.HandleUnknownMessageType.bind(this.tgErrorHandler)
    );

    this.tgJoinHandler = new TgJoinHandler(
      this.config.irc.showJoinMessage,
      this.sendMessageToIrc.bind(this)
    );
    this.tgEventListener.addListener(
      'join',
      this.tgJoinHandler.RelayJoinMessage.bind(this.tgJoinHandler)
    );

    this.tgMsgHandler = new TgMessageHandler(
      this.config.irc,
      true,
      this.sendMessageToIrc.bind(this)
    );
    this.tgEventListener.addListener(
      'message',
      this.tgMsgHandler.RelayMessage.bind(this.tgMsgHandler)
    );

    this.tgPartHandler = new TgPartHandler(
      this.config.irc.showLeaveMessage,
      this.sendMessageToIrc.bind(this)
    );
    this.tgEventListener.addListener(
      'part',
      this.tgPartHandler.RelayPartMessage.bind(this.tgPartHandler)
    );

    if (this.config.imgur.useImgurForImageLinks) {
      this.tgPhotoHandler = new TgImgurPhotoHandler(
        require('imgur'),
        true,
        this.tgbot,
        this.sendMessageToIrc.bind(this)
      );
    }
    else {
      this.tgPhotoHandler = new TgPhotoHandler(
        true,
        this.tgbot,
        this.sendMessageToIrc.bind(this)
      );
    }
    this.tgEventListener.addListener(
      'photo',
      this.tgPhotoHandler.RelayPhotoMessage.bind(this.tgPhotoHandler)
    );

    this.tgStickerHandler = new TgStickerHandler(
      this.config.irc,
      this.config.irc.sendStickerEmoji,
      this.sendMessageToIrc.bind(this)
    );
    this.tgEventListener.addListener(
      'sticker',
      this.tgStickerHandler.RelayStickerMessage.bind(this.tgStickerHandler)
    );
  }

  // -------- Send Functions --------

  /**
   * Sends a message to telegram.
   * @param messageString - The message to send.
   */
  queueTelegramMessage(messageString) {
    this.tgRateLimiter.queueMessage(messageString);
  }

  // -------- Helper Functions --------

  /**
   * Checks the given configuration option and it makes sure
   * it exists, is not undefined, and is not null.  If it is any
   * of these things, it gets set to the default value.
   *
   * @param option - Reference to the option that needs to be checked.
   * @param defaultValue - The default value to set the option to if the option
   *                       is null or undefined.
   * @param optionName - Name of the option (Would be nice if JS had C#'s nameof...)
   * @returns The default option if undefined or null was passed in for
   *          the option parameter, otherwise just the option parameter.
   */
  checkConfigOption(option, defaultValue, optionName) {
    if ((option === undefined) || (option === null)) {
      option = defaultValue;
      console.log("Using default value for " + optionName);
    }

    return option;
  }

  /**
   * Sends a message to the IRC channel.
   * @param {String} message - The message to send to the IRC channel.
   */
  sendMessageToIrc(message) {
    this.ircbot.say(
      this.config.irc.channel,
      message
    );
  }

  /**
   * Single choke-point for reporting status to the console.
   * @param {String} message - The message to log.
   */
  status(message) {
    console.log(message);
  }
}

module.exports = TeleIrc;
