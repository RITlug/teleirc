const tg = require("node-telegram-bot-api");
const irc = require("irc");

const config = require("./config.js");

// Read in the config file to hide the bot API token
console.log("Reading config.json file to get bot API token...");

// Check for a telegram token
if (!config.token) {
    setupError("Unable to find a telegram bot token in config.js");
}

// Check for required IRC settings:
let ircooptions = ["server", "channel", "botName"];
if (config.irc) {
    ircooptions.forEach(option => {
        if (!config.irc[option]) {
            setupError("Unable to find IRC settings for " + option);
        }
    });

    // The following settings are optional. If there are no 
    // options set for them, set default values.

    if (!config.irc.hasOwnProperty("prefix")) {
        console.log("Using default value for prefix");
        config.irc.prefix = "";
    }

    if (!config.irc.hasOwnProperty("suffix")) {
        console.log("Using default value for suffix");
        config.irc.suffix = "";
    }

} else {
    setupError("Unable to find IRC settings in config.js");
}

// Check for required Telegram settings:
let tgoptions = ["chatId"];
if (config.tg) {
    tgoptions.forEach(option => {
        if (!config.tg[option]) {
            setupError("Unable to find Telegram settings for " + option);
        }
    });

    // The following settings are optional. If there are no 
    // options set for them, default to false.

    // Read config file for showJoinMessage
    if (!config.tg.hasOwnProperty("showJoinMessage")) {
        console.log("Using default 'false' value for showJoinMessage");
        config.tg.showJoinMessage = false;
    }

    // Read config file for showLeaveMessage
    if (!config.tg.hasOwnProperty("showLeaveMessage")) {
        console.log("Using default 'false' value for showLeaveMessage");
        config.tg.showLeaveMessage = false;
    }

    // Read config file for showKickMessage
    if (!config.tg.hasOwnProperty("showKickMessage")) {
        console.log("Using default 'false' value for showKickMessage");
        config.tg.showKickMessage = false;
    }

    if (!config.tg.hasOwnProperty("showActionMessage")) {
        console.log("Using default 'false' value for showActionMessage");
        config.tg.showActionMessage = false;
    }

} else {
    setupError("Unable to find Telegram settings in config.js");
}

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

tgbot.on('message', function (msg) {
    // Only relay messages that come in through the Telegram chat
    if (msg.chat.id == config.tg.chatId) {
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

            // Check if this message is a new user joining the group chat:
            if (msg.new_chat_member) {
                let username = msg.new_chat_member.username;
                let first_name = msg.new_chat_member.first_name;
                ircbot.say(config.channel, "New user " + first_name + " ( @" + username + ") has joined the Telegram Group!");
            } else if (msg.sticker != undefined) {
                // If we have a sticker, we should send it to IRC... sort of...
                // The best way to get a sticker to IRC would be to send a URL to the sticker, but that doesn't
                // seem to exist.  My guess as to how telegram handles stickers is it sends the file ID of the sticker
                // to its servers and downloads it directly from its servers, similar to how photo downloads are done.
                // That won't work in IRC.
                //
                // The only thing we can do if we see a sticker is grab its corresponding emoji and send that to IRC.
                // That is, if the config allows us to do that of course.
                let emoji = msg.sticker.emoji;
                if (emoji != undefined && config.irc.sendStickerEmoji) {
                    ircbot.say(config.irc.channel, config.irc.prefix + from + config.irc.suffix + " " + emoji);
                }
            } else {
                console.log("Ignoring non-text message: " + JSON.stringify(msg));
            }
        } else {
            // Relay all text messages into IRC
            ircbot.say(config.irc.channel, config.irc.prefix + from + config.irc.suffix + " " + message);
        }
    } else {
        // Messages that are sent to the bot outside of the group chat should just be dumped
        // to the console for potential testing and debugging to do things like check chat IDs
        // and verify the JSON formats of various messages
        console.log("[TG Debug] " + JSON.stringify(msg));
    }
});

// Action to invoke on incoming messages from the IRC side
ircbot.addListener('message', (from, channel, message) => {
    let matchedNames = config.ircBlacklist.filter(function (name) {
        return from.toLowerCase() === name.toLowerCase();
    });

    if (matchedNames.length <= 0) {
        tgbot.sendMessage(config.tg.chatId, from + ": " + message);
    }

});

ircbot.addListener('error', (message) => {
    console.log("[IRC Debug] " + JSON.stringify(message));
});

// These additional alerts can be turned on in the config file
if (config.tg.showActionMessage) {
    ircbot.addListener('action', (from, channel, message) => {
        let matchedNames = config.ircBlacklist.filter(function (name) {
            return from.toLowerCase() === name.toLowerCase();
        });

        if (matchedNames.length <= 0) {
            tgbot.sendMessage(config.tg.chatId, from + " " + message);
        }
    });
}

if (config.tg.showJoinMessage) {
    // Let the telegram chat know when a user joins the IRC channel
    ircbot.addListener('join', (channel, username) => {
        tgbot.sendMessage(config.tg.chatId, username + " has joined " + channel + " channel.");
    });
}

if (config.tg.showLeaveMessage) {
    // Let the telegram chat know when a user leaves the IRC channel
    ircbot.addListener('part', (channel, username, reason) => {
        if (typeof reason != "string") {
            reason = "Parting...";
        }
        tgbot.sendMessage(config.tg.chatId, username + " has left " + channel + ": " + reason + ".");
    });
}

if (config.tg.showKickMessage) {
    // Let the telegram chat know when a user is kicked from the IRC channel
    ircbot.addListener('kick', (channel, username, by, reason) => {
        if (typeof reason != "string") {
            reason = "Kicked";
        }
        tgbot.sendMessage(config.tg.chatId, username + " was kicked by " + by + " from " + channel + ": " + reason + ".");
    });
}

// Quick function to print an error message and bail out if the 
// config is missing any required settings in the config file.
function setupError(message) {
    console.warn("[setup] " + message);
    process.exit(1);
}
