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
  port: config.irc.port,
  selfSigned: config.irc.tlsAllowSelfSigned,
  certExpired: config.irc.tlsAllowCertExpired,
  channels: [config.irc.channel],
  debug: false,
  username: config.irc.botName,
  autoConnect: false,
  autoRejoin: true
});

// Create the telegram bot side with the settings specified in config object above
console.log("Starting up bot on Telegram...");
let tgbot = new tg(config.token, { polling: true });
teleIrc.initStage3_initBots(ircbot, tgbot);

console.log("Adding IRC Listeners...");
teleIrc.initStage4_addIrcListeners();

console.log("Enabling Telegram message sending...");
teleIrc.initStage5_initTelegramMessageSending();

console.log("Setup complete! TeleIRC now running.");
