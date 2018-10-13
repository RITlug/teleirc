'use strict';

const dotenv = require('dotenv').config();
const tg = require("node-telegram-bot-api");
const irc = require("irc");
const TeleIrc = require("./lib/TeleIrc");
const config = require("./config");

// Read in the config file to hide the bot API token
console.log("Reading config file to get bot API token...");

let teleIrc = new TeleIrc(config);
console.log("Reading in IRC Configuration...");
teleIrc.initStage1_ircConfigValidation();
console.log("Reading in Telegram Configuration...");
teleIrc.initStage2_telegramConfigValidation();

// Create the IRC bot side with the settings specified in config object above
console.log("Starting up bot on IRC...");
let ircbot = new irc.Client(config.irc.server, config.irc.botName, {
  channels: [config.irc.channel],
  debug: false,
  username: config.irc.botName,
  autoConnect: false,
  autoRejoin: true
});

// Create the telegram bot side with the settings specified in config object above
console.log("Starting up bot on Telegram...");
let tgbot = new tg(config.token, { polling: true });
// This is super rudimentary error handling. If the Telegram bot
// throws an error, the library we use simply logs it and appears
// to hang. This may be a bug in that library (see
// https://github.com/yagop/node-telegram-bot-api/issues/657).
// This code forces teleirc to fully crash when it encounters a polling
// error so that its parent process/container can restart it.
// It may be possible to handle this in a nicer way, but issues
// open on the bot library indicate that things may not be working right
// as of October 2018.
tgbot.on('polling_error', (error) => {
    console.log("Fatal Telegram polling error: " + error.code);
    console.log(error.stack);
    process.exit(1);
});
tgbot.on('webhook_error', (error) => {
    console.log("Fatal Telegram webhook error: " + error.code);
    console.log(error.stack);
    process.exit(1);
});

teleIrc.initStage3_initBots(ircbot, tgbot);

console.log("Adding IRC Listeners...");
teleIrc.initStage4_addIrcListeners();

console.log("Enabling Telegram message sending...");
teleIrc.initStage5_initTelegramMessageSending();

console.log("Setup complete! Teleirc now running.");
