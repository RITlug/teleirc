var tg = require("node-telegram-bot-api");
var irc = require("irc");
var fs = require("fs");

// Read in the config file to hide the bot API token
console.log("Reading config.json file to get bot API token...");
var settings = JSON.parse(fs.readFileSync('config.json', 'utf8'));
if (settings.token) {
    console.error("Token value not found in config.json file");
}

var token = settings.token;

var config = {
    server: "irc.freenode.net",
    botName: "ritlugtg",
    channel: "#ritlug",
    chatId: -13280454,
    ircBlacklist: [
        "CowSayBot"
        ]
};

// Create the IRC bot side with the settings specified in config object above
console.log("Starting up bot on irc...");
var ircbot = new irc.Client(config.server, config.botName, {
    channels: [config.channel],
    debug: false,
    username: config.botName
});

// Create the telegram bot side with the settings specified in config object above
console.log("Starting up bot on telegram...");
var tgbot = new tg(token, { polling: true });

tgbot.on('message', function (msg) {
    // Only relay messages that come in through the RITLug telegram chat
    if (msg.chat.id === config.chatId) {
        var from = msg.from.username;

        // Do some basic cleanup if the user does not have a username
        // on telegram. Replace with first_name instead.
        if (msg.from.username === undefined) {
            from = msg.from.first_name;
        }

        // Check that this message has a text field. If it does not,
        // it is something special to telegram like a file or sticker
        // and should not be passed to IRC
        var message = msg.text;
        if (msg.text === undefined) {
            console.log("Ignoring non-text message: " + JSON.stringify(msg));
        } else {
            // Relay all text messages into IRC
            ircbot.say(config.channel, from + ": " + message);
        }
    } else {
        // Messages that are sent to the bot outside of the RITLug chat should just be dumped
        // to the console for potential testing and debugging to do things like check chat IDs
        // and verify the JSON formats of various messages
        console.log("Debug: " + JSON.stringify(msg));
    }
});

// Action to invoke on incoming messages from the IRC side
ircbot.addListener('message', function (from, channel, message) {
    // Anything coming from IRC is going to be valid to display as text
    // in Telegram. Just do a quick passthrough. No checking. 
    if (config.ircBlacklist.indexOf(from) === -1) {
        tgbot.sendMessage(config.chatId, from + ": " + message);
    }
    
});
