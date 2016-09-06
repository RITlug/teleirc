var tg = require("node-telegram-bot-api");
var irc = require("irc");
var fs = require("fs");

// If you want to change the default config file name, do that here.
var configfilename = 'config.json';

// This is the base config object that will store values from the settings file.
var config = {};

// Read in the config file to hide the bot API token
console.log("Reading config.json file to get bot API token...");
var settings = JSON.parse(fs.readFileSync(configfilename, 'utf8'));

// Set up IRC settings
if (settings.irc) {
    // Get the IRC server name
    if (settings.irc.server) {
        config.server = settings.irc.server;
    } else {
        setupError("Unable to find IRC settings for server address");
    }

    // Get the IRC channel name
    if (settings.irc.channel) {
        config.channel = settings.irc.channel;
    } else {
        setupError("Unable to find IRC settings for channel name");
    }

    // Get the IRC bot account name
    if (settings.irc.botName) {
        config.botName = settings.irc.botName;
    } else {
        setupError("Unable to find IRC settings for bot name");
    }
} else {
    setupError("Unable to find IRC settings in " + configfilename);
}

if (settings.tg) {
    // Get the Telegram group chat id
    if (settings.tg.chatId) {
        config.chatId = settings.tg.chatId;
    } else {
        setupError("Unable to find Telegram settings for chatId");
    }

    // Read config file to determine various ettings
    if (settings.tg.spacer) {
        console.log("Setting text spacer to '" + settings.tg.spacer + "'")
        config.spacer = settings.tg.spacer;
    } else {
        config.spacer = ":";
    }

    // Read config file for showJoinMessage
    if (settings.tg.showJoinMessage) {
        config.showJoinMessage = settings.tg.showJoinMessage;
    } else {
        config.showJoinMessage = false;
    }

    // Read config file for showLeaveMessage
    if (settings.tg.showLeaveMessage) {
        config.showLeaveMessage = settings.tg.showLeaveMessage;
    } else {
        config.showLeaveMessage = false;
    }

    // Read config file for showKickMessage
    if (settings.tg.showKickMessage) {
        config.showKickMessage = settings.tg.showKickMessage;
    } else {
        config.showKickMessage = false;
    }

    if (settings.tg.showActionMessage) {
         config.showActionMessage = settings.tg.showActionMessage;
    } else {
        config.showActionMessage = false;
    }

} else {
    setupError("Unable to find Telegram settings in " + configfilename);
}

var token = null;
if (settings.token) {
    token = settings.token;
} else {
    setupError("Unable to find a telegram bot token in " + configfilename);
}

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
            ircbot.say(config.channel, from + config.spacer + " " + message);
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
    var matchedNames = config.ircBlacklist.filter(function (name) {
        return from.toLowerCase() === name.toLowerCase();
    });

    if (matchedNames.length <= 0) {
        tgbot.sendMessage(config.chatId, from + ": " + message);
    }

});

ircbot.addListener('error', function (message) {
    console.log("[IRC Debug] " + JSON.stringify(message));
});




// These additional alerts can be turned on in the config file
if (config.showActionMessage) {
    ircbot.addListener('action', function (from, channel, message) {
        var matchedNames = config.ircBlacklist.filter(function (name) {
            return from.toLowerCase() === name.toLowerCase();
        });

        if (matchedNames.length <= 0) {
            tgbot.sendMessage(config.chatId, from + " " + message);
        }
    });
}

if (config.showJoinMessage) {
    // Let the telegram chat know when a user joins the IRC channel
    ircbot.addListener('join', function (channel, username) {
        tgbot.sendMessage(config.chatId, username + " has joined " + config.channel + " channel.");
    });
}

if (config.showLeaveMessage) {
    // Let the telegram chat know when a user leaves the IRC channel
    ircbot.addListener('part', function (channel, username, reason) {
        tgbot.sendMessage(config.chatId, username + " has left " + channel + ": " + reason + ".");
    });
}

if (config.showKickMessage) {
    // Let the telegram chat know when a user is kicked from the IRC channel
    ircbot.addListener('kick', function (channel, username, by, reason) {
        tgbot.sendMessage(config.chatId, username + " was kicked by " + by + " from " + channel + ": " + reason + ".");
        console.log('%s was kicked from %s by %s: %s', who, channel, by, reason);
    });
}

// Quick function to print an error message and bail out if the 
// config is missing any required settings in the config file.
function setupError(message) {
    console.warn("[setup] " + message);
    process.exit(1);
}
