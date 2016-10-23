teleirc
=======

NodeJS implementation of a [Telegram](https://telegram.org/) <=> [IRC](https://en.wikipedia.org/wiki/Internet_Relay_Chat) bridge for use with any IRC channel and Telegram group


## About

This project was written by [Mark Repka](https://github.com/repkam09) to help bridge communication between [RITlug](http://ritlug.com)'s IRC channel with our Telegram group. The project is a lightweight NodeJS application. This is also being used in various IRC channels and Telegram groups than just RITlug.

### Try it out

A public Telegram supergroup and IRC channel are available for you to see it live in action. This is also where our developers and community hang out as well.

* **[Telegram](https://telegram.me/teleirc)**
* **[IRC](https://webchat.freenode.net/?channels=ritlug-teleirc)** (#ritlug-teleirc @ irc.freenode.net)

### Who's using this app?

Ask any of our successful users!

* [RITlug teleirc](https://github.com/RITlug/teleirc) ([Telegram](https://telegram.me/teleirc) | [IRC](https://webchat.freenode.net/?channels=ritlug-teleirc))
* [RITlug](http://ritlug.com) (_Telegram invite-only_ | [IRC](https://webchat.freenode.net/?channels=ritlug))
* [Fedora Project](https://fedoraproject.org/wiki/Overview) ([Telegram](https://telegram.me/fedora) | [IRC](https://webchat.freenode.net/?channels=fedora-telegram))
* [LibreOffice](https://www.libreoffice.org/) (_Telegram invite-only_ | [IRC](https://webchat.freenode.net/?channels=libreoffice-telegram))
* [FOSS@MAGIC at RIT](http://foss.rit.edu) ([Telegram](https://telegram.me/fossrit) | [IRC](https://webchat.freenode.net/?channels=rit-foss))
* [Fedora LATAM](http://fedoracommunity.org/latam) ([Telegram](https://telegram.me/fedoralatam) | [IRC](https://webchat.freenode.net/?channels=fedora-latam))
* [MINECON Agents](https://mojang.com/2016/06/calling-all-agents-help-us-run-minecon-2016/) (_Telegram invite-only_ | [IRC](https://webchat.esper.net/?channels=MineconAgents))
* [Partido Pirata de Argentina](https://github.com/piratas-ar) (_Telegram invite-only_ | [IRC](https://webchat.pirateirc.net/?channels=PPar))


## Installation

In order to use this bridge, you will need several pieces of information. This guide is broken up into Telegram, teleirc, and IRC sections.

### Telegram

For the Telegram side, you will need to create a new Telegram bot that will sit inside your Telegram group. You will need to do a little configuration and information gathering with the bot.

1. Send a `/start` Telegram message to the @BotFather, the [Telegram bot](https://core.telegram.org/bots) for [creating Telegram bots](https://core.telegram.org/bots#6-botfather).
2. Create a bot with the `/newbot` command in a chat window with BotFather. You will then be prompted to enter a bot name. Once you have done this, you will get a bot token from the BotFather for accessing the Telegram API. Make note of the token for later configuration.
3. Before this bot can enter any group chats, you will need to configure it with correct permissions. Send the `/setprivacy` command to the BotFather, specify which bot this command is for, then **disable** the privacy so the bot receives all messages sent in the group chat.
4. Optionally, you can supply a different bot name, description, and picture to make it more obvious what the bot is in the group chat.
5. This bot supports adhering to Telegram's bot rate limits. You can set messages per minute in the config. As of this writing, the limit
from Telegram is 20, which is the default. If you don't want to use rate limiting, set the option to 0. The bot will bundle messages and
send them together with a delay while rate limiting in hopes of avoiding dropped messages.
6. Add the bot to the Telegram group chat you want to bridge.

Now that you have a bot created, have its Telegram API token, and have it in your group, you can start using this bridge.

### Teleirc

To get teleirc working, you will need a server to run it on, git, and the latest version of NodeJS installed.

1. Clone this repository to the server you wish to run teleirc on. `git clone git@github.com:RITlug/teleirc.git`
2. Install dependencies for teleirc with _npm_. This pulls down required NPM packages for Telegram and IRC connections. `npm install`
3. Rename `config.js.example` to `config.js`. Change the token configuration value to the bot token you received from the BotFather. An example of what that file looks like is farther below.
4. **This API key should be kept private!** This config file is listed in the .gitignore file so it is not accidentially added to source control.
5. Edit `config.js` and update `server`, `botName`, `channel`, and `chatId` for your uses.
 * _server_: IRC server you wish to connect to (default: irc.freenode.net)
 * _botName_: Nickname for your bot to use on IRC
 * _channel_: IRC channel you want your bot to join
 * _chatId_: Telegram chat ID of the group you are bridging ([how do I get this?](http://stackoverflow.com/a/32572159))
6. Optional settings are available to customize teleirc to respond to extra irc events such as join, leave, kick, action messages. You can also set a custom prefix and suffix that will be inserted when a messages from Telegram is sent to IRC.

Alternatively, if you start up the bot with no Telegram chat ID set, it will sit waiting for messages to be sent to it. If you invite the bot to your group chat, you should see a "Debug TG" message with some information about the invite that was sent. One of the fields here will be the chatId. This is the value that needs to be put in the config object. Be careful not to get the user ID of a specific user when reading these messages.

#### Example: config.js

```javascript
{
    token: "000000000:AAAAAAaAAa2AaAAaoAAAA-a_aaAAaAaaaAA",
    ircBlacklist: [
        "CowSayBot"
    ],
    irc: {
        server: "irc.freenode.net",
        channel: "#channel",
        botName: "teleirc",
        prefix: "<",
        suffix: ">"
    },
    tg: {
        chatId: "-0000000000000",
        showJoinMessage: false,
        showActionMessage: true,
        showLeaveMessage: false,
        showKickMessage: false,
        maxMessagesPerMinute: 20
    }
}
```

### IRC

There is not real configuration needed on the IRC side, as IRC is generally very open. It might be a good idea to [register the channel](https://infrastructure.fedoraproject.org/infra/docs/freenode-irc-channel.rst) you are using.


# Running Teleirc

Before running teleirc, you will need to decide how you want to run it persistently. An easy way to keep this service running in the background on your server is through `pm2`. pm2 is an npm package that can keep node services running in the background, restart them if they crash, and restart them if the server reboots.

Alternatively, you can handle this yourself by using something like `screen` or `tmux`, and a quick shell script and starting the program manually with `node teleirc.js`.

## tmux

1. Install tmux to your server.
 * _RHEL / CentOS_: `sudo yum install tmux` ([No package found? Get EPEL](https://fedoraproject.org/wiki/EPEL))
 * _Fedora_: `sudo dnf install tmux`
 * _Debian / Ubuntu_: `sudo apt-get install tmux`
2. Create a new tmux session. `tmux new-session -s teleirc`
3. Navigate to the directory where you have teleirc and start it with node. `node teleirc.js`
4. Exit tmux by typing `CTRL+B`, then the `d` key.


# Hey, I use this project too!

Want to have your name added to the list in this README? Let us know you're using teleirc too! [Submit an issue](https://github.com/RITlug/teleirc/issues/new) against this repo with the following info:

* Organization / group name and website
* Telegram group URL
* Your IRC channel

Please note, to be added, your group must not discuss / contain inappropriate or explicit content.


# License

Licensed under the [MIT License](https://github.com/RITlug/teleirc/blob/master/LICENSE). If you're hacking on teleirc, we'd love to see you bring your improvements back upstream!
