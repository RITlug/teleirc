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
    botName: "teleirc",
    channel: "#mychannel",
    chatId: -0000000000000,
    // IRC users to not forward messages from
    ircBlacklist: [
        "CowSayBot"
    ]
};

// Create the IRC bot side with the settings specified in config object above
console.log("Starting up bot on IRC...");
var ircbot = new irc.Client(config.server, config.botName, {
    channels: [config.channel],
    debug: false,
    username: config.botName
});

// Create the telegram bot side with the settings specified in config object above
console.log("Starting up bot on Telegram...");
var tgbot = new tg(token, { polling: true });

tgbot.on('message', function (msg) {
    // Only relay messages that come in through the Telegram chat
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

            // Check if this message is a new user joining the group chat:
            if (msg.new_chat_member) {
                var username = msg.new_chat_member.username;
                var first_name = msg.new_chat_member.first_name;
                ircbot.say(config.channel, "New user " + first_name + " ( @" + username + ") has joined the Telegram Group!");
            } else {
                console.log("Ignoring non-text message: " + JSON.stringify(msg));
            }
        } else {
            // Relay all text messages into IRC
            ircbot.say(config.channel, from + ": " + message);
        }
    } else {
        // Messages that are sent to the bot outside of the group chat should just be dumped
        // to the console for potential testing and debugging to do things like check chat IDs
        // and verify the JSON formats of various messages
        console.log("[TG Debug] " + JSON.stringify(msg));
    }
});

// Action to invoke on incoming messages from the IRC side
ircbot.addListener('message', function (from, channel, message) {
    // Anything coming from IRC is going to be valid to display as text
    // in Telegram. Just do a quick passthrough. No checking.
    var matchedNames = config.ircBlacklist.filter(function (name) {
        return from.toLowerCase() === name.toLowerCase();
    });

    if (matchedNames.length <= 0) {
        tgbot.sendMessage(config.chatId, from + ": " + message);
    }

});

ircbot.addListener('action', function (from, channel, message) {
    // Anything coming from IRC is going to be valid to display as text
    // in Telegram. Just do a quick passthrough. No checking.
    var matchedNames = config.ircBlacklist.filter(function (name) {
        return from.toLowerCase() === name.toLowerCase();
    });

    if (matchedNames.length <= 0) {
        tgbot.sendMessage(config.chatId, from + " " + message);
    }

});

ircbot.addListener('error', function (message) {
    console.log("[IRC Debug] " + JSON.stringify(message));
});




// Let the telegram chat know when a user joins the IRC channel
ircbot.addListener('join', function (channel, username) {
    tgbot.sendMessage(config.chatId, username + " has joined " + config.channel + " channel.");
});

/*
// Adding these here for future reference
ircbot.addListener('part', function(channel, who, reason) {
    console.log('%s has left %s: %s', who, channel, reason);
});
ircbot.addListener('kick', function(channel, who, by, reason) {
    console.log('%s was kicked from %s by %s: %s', who, channel, by, reason);
});
*/
