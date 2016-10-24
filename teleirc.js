const tg = require("node-telegram-bot-api");
const irc = require("irc");
const TeleIrc = require("./lib/libteleirc.js");
const config = require("./config.js");

// Read in the config file to hide the bot API token
console.log("Reading config.json file to get bot API token...");

let teleIrc = new TeleIrc(config);
teleIrc.initStage1_ircConfigValidation();
teleIrc.initStage2_telegramConfigValidation();

// Create the IRC bot side with the settings specified in config object above
console.log("Starting up bot on IRC...");
let ircbot = new irc.Client(config.irc.server, config.irc.botName, {
    channels: [config.irc.channel],
    debug: false,
    username: config.irc.botName
});

// Create the telegram bot side with the settings specified in config object above
console.log("Starting up bot on Telegram...");
let tgbot = new tg(config.token, { polling: true });

teleIrc.initStage3_initBots(ircbot, tgbot);
teleIrc.initStage4_addIrcListeners();
teleIrc.initStage5_initTelegramMessageSending();
