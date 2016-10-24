const MessageRateLimiter = require("./MessageRateLimiter.js");

class TeleIrc {

    // ---------------- Constructor ----------------

    /**
     * @param teleIrcConfig - The teleirc config to use.
     */
    constructor(teleIrcConfig){
        this.config = teleIrcConfig;
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
    initStage1_ircConfigValidation(){
        let ircooptions = ["server", "channel", "botName"];
        if (this.config.irc) {
            ircooptions.forEach(option => {
                if (!this.config.irc[option]) {
                    throw "Unable to find IRC settings for " + option;
                }
            });

            // The following settings are optional. If there are no
            // options set for them, set default values.

            if (!this.config.irc.hasOwnProperty("prefix")) {
                console.log("Using default value for prefix");
                this.config.irc.prefix = "";
            }

            if (!this.config.irc.hasOwnProperty("suffix")) {
                console.log("Using default value for suffix");
                this.config.irc.suffix = "";
            }

        } else {
            throw "Unable to find IRC settings in config.js";
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
    initStage2_telegramConfigValidation(){
        if (!this.config.token){
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

            if (!this.config.tg.hasOwnProperty("maxMessagesPerMinue")) {
                console.log("Using default of 20 for maxMessagesPerMinute");
                this.config.tg.maxMessagesPerMinute = 20;
            }

        } else {
            throw "Unable to find Telegram settings in config.js";
        }
    }

    // ---- Telegram/IRC bot construction ----

    /**
     * After initializting the bots outside of this class,
     * this gives this class a reference to those.
     * @param ircbot - The IRC bot to use.
     * @param tgbot - The Telegram bot to use.
     */
    initStage3_initBots(ircbot, tgbot){
        this.ircbot = ircbot;
        this.tgbot = tgbot;
    }

    /**
     * Stage 4 of initialization:
     * Add all the IRC bot's listeners based on this object's config.
     */
    initStage4_addIrcListeners(){
        // Action to invoke on incoming messages from the IRC side
        this.ircbot.addListener('message', (from, channel, message) => {
            let matchedNames = this.config.ircBlacklist.filter(function (name) {
                return from.toLowerCase() === name.toLowerCase();
            });

            if (matchedNames.length <= 0) {
                this.queueTelegramMessage(this.config.tg.chatId, from + ": " + message);
            }
        });

        this.ircbot.addListener('error', (message) => {
            console.log("[IRC Debug] " + JSON.stringify(message));
        });

        // These additional alerts can be turned on in the config file
        if (this.config.tg.showActionMessage) {
            this.ircbot.addListener('action', (from, channel, message) => {
                let matchedNames = this.config.ircBlacklist.filter(function (name) {
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
        let teleirc = this;
        this.tgRateLimiter = new MessageRateLimiter(
            this.config.tg.maxMessagesPerMinute,
            60,
            function(message){
                teleirc.tgbot.sendMessage(teleirc.config.tg.chatId, message);
            });

        this.tgbot.on('message', this.handleTelegramMsg.bind(teleirc));
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
                    console.log("Ignoring non-text message: " + JSON.stringify(msg));
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

module.exports = TeleIrc;