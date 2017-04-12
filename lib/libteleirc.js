'use strict';

const MessageRateLimiter = require("./MessageRateLimiter.js");

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

/**
 * Exception that gets thrown when there's an error.
 */
class TeleIrcException {

    // ---------------- Constructor ----------------

    /**
     * Constructor.
     * @param errorCode - Unique number for this error.
     * @param message - Corresponding Message.
     */
    constructor(errorCode, message) {
        this.errorCode = errorCode;
        this.message = message;
    }

    toString() {
        return "TeleIrc Error " + this.errorCode + ": " + this.message;
    }
}

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
                    }
                    else if (option === "channel") {
                        errorCode = this.errorCodes.MissingIrcChannelConfig;
                    }
                    else if (option === "botName") {
                        errorCode = this.errorCodes.MissingIrcBotNameConfig;
                    }
                    throw new TeleIrcException(errorCode, "Unable to find IRC settings for " + option);
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
    }

    /**
     * Stage 4 of initialization:
     * Add all the IRC bot's listeners based on this object's config.
     */
    initStage4_addIrcListeners() {
        // Action to invoke on incoming messages from the IRC side
        this.ircbot.addListener('message', (from, channel, message) => {
            let matchedNames = this.config.ircBlacklist.filter((name) => {
                return from.toLowerCase() === name.toLowerCase();
            });

            if (matchedNames.length <= 0) {
                this.queueTelegramMessage(this.config.tg.chatId, "<" + from + "> " + message);
            }
        });

        this.ircbot.addListener('error', (message) => {
            console.log("[IRC Debug] " + JSON.stringify(message));
        });

        // These additional alerts can be turned on in the config file
        if (this.config.tg.showActionMessage) {
            this.ircbot.addListener('action', (from, channel, message) => {
                let matchedNames = this.config.ircBlacklist.filter((name) => {
                    return from.toLowerCase() === name.toLowerCase();
                });

                if (matchedNames.length <= 0) {
                    this.queueTelegramMessage(this.config.tg.chatId, from + " " + message);
                }
            });
        }

        if (this.config.tg.showJoinMessage) {
            // Let the telegram chat know when a user joins the IRC channel
            this.ircbot.addListener('join', (channel, username) => {
                this.queueTelegramMessage(this.config.tg.chatId, username + " has joined " + channel + " channel.");
            });
        }

        if (this.config.tg.showLeaveMessage) {
            // Let the telegram chat know when a user leaves the IRC channel
            this.ircbot.addListener('part', (channel, username, reason) => {
                if (typeof reason != "string") {
                    reason = "Parting...";
                }
                this.queueTelegramMessage(this.config.tg.chatId, username + " has left " + channel + ": " + reason + ".");
            });
        }

        if (this.config.tg.showKickMessage) {
            // Let the telegram chat know when a user is kicked from the IRC channel
            this.ircbot.addListener('kick', (channel, username, by, reason) => {
                if (typeof reason != "string") {
                    reason = "Kicked";
                }
                this.queueTelegramMessage(this.config.tg.chatId, username + " was kicked by " + by + " from " + channel + ": " + reason + ".");
            });
        }
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
        let teleirc = this;
        this.tgRateLimiter = new MessageRateLimiter(
            this.config.tg.maxMessagesPerMinute,
            60,
            (message) => {
                teleirc.tgbot.sendMessage(teleirc.config.tg.chatId, message);
            });

        this.tgbot.on('message', this.handleTelegramMsg.bind(teleirc));
        this.tgbot.on('edited_message', this.handleTelegramMsg.bind(teleirc));
    }

    // -------- Send Functions --------

    /**
     * Sends a message to telegram.
     * @param chatId - The telegram chat ID to send the message to.
     * @param messageString - The message to send.
     */
    queueTelegramMessage(chatId, messageString) {
        this.tgRateLimiter.queueMessage(messageString);
    }

    /**
     * Handles a message from telegram.
     */
    handleTelegramMsg(msg) {
        // Only relay messages that come in through the Telegram chat
        if (msg.chat.id == this.config.tg.chatId) {
            let from = msg.from.username;

            // Do some basic cleanup if the user does not have a username
            // on telegram. Replace with first_name instead.
            if (msg.from.username === undefined) {
                from = msg.from.first_name;
            }

            // Check that this message has a text field. If it does not,
            // it is something special to telegram like a file or sticker
            // and should not be passed to IRC
            let message = msg.text;
            if (msg.text === undefined) {
                if (msg.new_chat_member && this.config.irc.showJoinMessage) {
                    // Check if this message is a new user joining the group chat:
                    let username = msg.new_chat_member.username;
                    let first_name = msg.new_chat_member.first_name;
                    this.ircbot.say(this.config.irc.channel, this.getTelegramToIrcJoinLeaveMsg(first_name, username, "has joined the Telegram Group!"));
                } else if (msg.left_chat_member && teleirc.config.irc.showLeaveMessage) {
                    // Check if this message is a user leaving the telegram group chat:
                    let username = msg.left_chat_member.username;
                    let first_name = msg.left_chat_member.first_name;
                    this.ircbot.say(this.config.irc.channel, this.getTelegramToIrcJoinLeaveMsg(first_name, username, "has left the Telegram Group."));
                } else if (msg.sticker !== undefined) {
                    // If we have a sticker, we should send it to IRC... sort of...
                    // The best way to get a sticker to IRC would be to send a URL to the sticker, but that doesn't
                    // seem to exist.  My guess as to how telegram handles stickers is it sends the file ID of the sticker
                    // to its servers and downloads it directly from its servers, similar to how photo downloads are done.
                    // That won't work in IRC.
                    //
                    // The only thing we can do if we see a sticker is grab its corresponding emoji and send that to IRC.
                    // That is, if the config allows us to do that of course.
                    let emoji = msg.sticker.emoji;
                    if (emoji !== undefined && this.config.irc.sendStickerEmoji) {
                        this.ircbot.say(this.config.irc.channel, this.config.irc.prefix + from + this.config.irc.suffix + " " + emoji);
                    }
                } else {
                    // Check if this message contains a document/file to display in irc
                    if (msg.document) {
                        this.tgbot.getFileLink(msg.document.file_id).then((url) => {
                            let newmessage = sg.document.file_name + " (" + msg.document.mime_type + ", " + filemsg.document.file_size_size + " bytes) : " + url;
                            this.ircbot.say(this.config.irc.channel, this.config.irc.prefix + from + this.config.irc.suffix + " " + newmessage);
                        });
                    } else {
                        console.log("Ignoring non-text message: " + JSON.stringify(msg));
                    }
                }
            } else {
                // Relay all text messages into IRC
                this.ircbot.say(this.config.irc.channel, this.config.irc.prefix + from + this.config.irc.suffix + " " + message);
            }
        } else {
            // Messages that are sent to the bot outside of the group chat should just be dumped
            // to the console for potential testing and debugging to do things like check chat IDs
            // and verify the JSON formats of various messages
            console.log("[TG Debug] " + JSON.stringify(msg));
        }
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
     * Generates the "User has joined" or "User has left" message that goes from telegram
     * to IRC in a way that does not cause undefines to appear from an undefined username.
     * The result is if a user has a user name, it will return:
     *     SomeUser (@SomeUserName) suffixHere
     * But if the user does not have a user name, it will return:
     *     SomeUSer suffixHere
     * @param {String} firstName - The telegram user's first name.
     * @param {String} userName - The telegram user's username.
     * @param {String} suffix - The message that appears after the firstname and username.
     */
    getTelegramToIrcJoinLeaveMsg(firstName, userName, suffix) {
        if (userName === undefined) {
            return firstName + " " + suffix;
        } else {
            return firstName + " (@" + userName + ") " + suffix;
        }
    }
}

module.exports = TeleIrc, TeleIrcException;
