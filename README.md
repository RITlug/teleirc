# teleirc
This project provides a simple implementation of a [Telegram](https://telegram.org/) to [IRC](https://en.wikipedia.org/wiki/Internet_Relay_Chat) bridge.



# Setup and Usage
In order to use this bridge you will need several pieces of information.

# Telegram Setup
For the Telegram side, you will need to start by creating a new Telegram Bot that will sit inside your group or chat.
Send a '/start' Telegram Message to the @BotFather, the Telegram bot for creating Telegram bots!

You can start creating a bot with the '/newbot' command. You will then be prompted to enter a bot name. Once you have done this, you will get a Bot Token from the BotFather for accessing the Telegram API.

Before this bot can enter any group chats, you will need to configure it to have the correct permissions. Send the '/setprivacy' command to the BotFather, specify which bot this command is for, then 'disable' the privacy so that the bot will receive all messages that are sent in the group chat.
Optionally you can supply a different bot name and picture to make it more obvious what the bot is in the group chat.

Now that you have a bot created and have its associated Telegram API key, you can start using this bridge.

# Teleirc Setup
Clone this repository to get the basic code ready. In this repository, open a command line and type 'npm install' to pull down the required NPM packages for the Telegram and IRC connections.
You will have to create a config.json file that contains your Telegram Bot API key with the following format:

        --- config.json ---
        {
            token: "000000000:AAAAAAaAAa2AaAAaoAAAA-a_aaAAaAaaaAA"
        }

This API key should be kept private. This config file is listed in the .gitignore file so it is not accidentially added to source control.


In the teleirc.js file you can fill in the information for the IRC server, channel, and irc username the bot should use.
You will now need to find the Telegram chatId for your group chat. This is a bit difficult, but you can use the teleirc debug information to help.

If you start up the bot, with no chat Id set, it will just sit there waiting for messages to be sent to it. If you invite the bot to your group chat, you should see a "Debug TG" message with some information about the invite that was sent. One of the fields here will be the chatId. This is the value that needs to be put in the config object.


# IRC Setup
There is not real configuration needed on the IRC side, as IRC is generally very open. But, if you want to reserve a username or something, you can do that so the bot name is not taken.

# Running Teleirc
An easy way to keep this service running in the backgrond on your server is through pm2. pm2 is an npm package that can keep node services running in the background, restart them if they crash, and restarting them if the server reboots.
You can also handle this yourself by using something like 'screen' and a quick shell script and starting the program manually with 'node teleirc.js'

