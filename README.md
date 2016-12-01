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
* [LibreOffice Community](https://www.libreoffice.org/) ([Telegram](https://telegram.me/libreofficecommunity) | [IRC](https://webchat.freenode.net/?channels=libreoffice-telegram))
* [FOSS@MAGIC at RIT](http://foss.rit.edu) ([Telegram](https://telegram.me/fossrit) | [IRC](https://webchat.freenode.net/?channels=rit-foss))
* [Fedora LATAM](http://fedoracommunity.org/latam) ([Telegram](https://telegram.me/fedoralatam) | [IRC](https://webchat.freenode.net/?channels=fedora-latam))
* [MINECON Agents](https://mojang.com/2016/06/calling-all-agents-help-us-run-minecon-2016/) (_Telegram invite-only_ | [IRC](https://webchat.esper.net/?channels=MineconAgents))
* [Sugar Labs](https://sugarlabs.org/) ([Telegram](https://telegram.me/sugarirc) | [IRC](https://webchat.freenode.net/?channels=sugar))
* [Partido Pirata de Argentina](https://github.com/piratas-ar) (_Telegram invite-only_ | [IRC](https://webchat.pirateirc.net/?channels=PPar))


## Installation

In order to use this bridge, you will need several pieces of information. This guide is broken up into Telegram, teleirc, and IRC sections.

On an ArchLinux based distro, teleirc can be installed [from the AUR](https://aur.archlinux.org/packages/teleirc/). You will still need to
follow the configuration steps from below. If you install from the AUR, teleirc's files will be located at `/var/lib/teleirc` and teleirc
can be managed via systemd.


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

#### Run using Docker

To get teleirc working, you will need a server to run it on and a recent version of Docker installed.

##### Which image do I choose?

Fedora, Ubuntu and Alpine Linux images are provided.

Their sizes may be deal breakers (ordered ascending by size):

| **Image**                                                                            | **Size** |
|--------------------------------------------------------------------------------------|----------|
| Alpine Linux(uses [mhart/alpine-node:6](https://hub.docker.com/r/mhart/alpine-node)) | 345 MB   |
| Ubuntu                                                                               | 598.4 MB |
| Fedora                                                                               | 1.13 GB  |

The below example uses alpine, replace `alpine` with `fedora` or `ubuntu` if you so choose.

You will see errors during `npm install`, ignore them. They are not fatal.

```bash
docker build . -f Dockerfile.alpine -t teleirc
docker run -d --name teleirc --restart always \
	-e TELEIRC_TOKEN="000000000:AAAAAAaAAa2AaAAaoAAAA-a_aaAAaAaaaAA" \
	-e IRC_CHANNEL="#channel" \
	-e IRC_BOT_NAME="teleirc" \
	-e IRC_BLACKLIST="CowSayBot,AnotherNickToIgnore" \
	-e TELEGRAM_CHAT_ID="-0000000000000" \
	teleirc
```

Optionally you may use [docker-compose](https://docs.docker.com/compose):

``` yaml
version: '2'
services:
  teleirc:
    build:
      context: .
      dockerfile: Dockerfile.alpine
    env_file: .env
```

We provide an example compose file (`docker-compose.yml.example`). You can optionally, tell Docker Compose to use the Fedora or Ubuntu base by changing dockerfile to use `Dockerfile.ubuntu` for Ubuntu, or `Dockerfile.fedora` for Fedora, but Fedora is a rather large image topping out at over 1GB. We recommend you not use it, but it is provided if you choose to use it.

We ignore the `docker-compose.yml` file in gitignore.

Either of the following will work fine:

- copy `docker-compose.yml.example` to `docker-compose.yml` and do `docker-compose up -d teleirc`
- do `docker-compose -f docker-compose.yml.example up -d teleirc`

##### Configuration Settings

* `TELEIRC_TOKEN`: Private API key for Telegram bot
* `IRC_BLACKLIST`: Comma-separated list of IRC nicks to ignore (default: "")
* `IRC_BOT_NAME`: Nickname for your bot to use on IRC (default: teleirc)
* `IRC_CHANNEL`: IRC channel you want your bot to join (default: #channel)
* `IRC_SEND_STICKER_EMOJI`: Send the emoji associated with a sticker on IRC (default: true)
* `IRC_PREFIX`: Text displayed before Telegram name in IRC (default: <)
* `IRC_SERVER`: IRC server you wish to connect to (default: irc.freenode.net)
* `IRC_SUFFIX`: Text displayed after Telegram name in IRC (default: >)
* `MAX_MESSAGES_PER_MINUTE`: Maximum rate at which to relay messages (default: 20)
* `SHOW_ACTION_MESSAGE`: Relay action messages (default: true)
* `SHOW_JOIN_MESSAGE`: Relay join messages (default: false)
* `SHOW_KICK_MESSAGE`: Relay kick messages (default: false)
* `SHOW_LEAVE_MESSAGE`: Relay leave messages (default: false)
* `TELEGRAM_CHAT_ID`: Telegram chat ID of the group you are bridging ([how do I get this?](http://stackoverflow.com/a/32572159))

Alternatively, if you start up the bot with no Telegram chat ID set, it will sit waiting for messages to be sent to it. If you invite the bot to your group chat, you should see a "Debug TG" message with some information about the invite that was sent. One of the fields here will be the chatId. This is the value that needs to be put in the config object. Be careful not to get the user ID of a specific user when reading these messages.

#### Run natively

To get teleirc working, you will need a server to run it on, git, and the latest version of NodeJS installed.

1. Clone this repository to the server you wish to run teleirc on. `git clone git@github.com:RITlug/teleirc.git`
2. Install dependencies for teleirc with _npm_. This pulls down required NPM packages for Telegram and IRC connections. `npm install`
3. Rename `config.js.example` to `config.js`.
4. **This API key should be kept private!** This config file is listed in the .gitignore file so it is not accidentially added to source control.
5. We use [dotenv](https://www.npmjs.com/package/dotenv) to allow for easy management of api keys and settings. Copy `.env.example` to `.env` -- this is in the gitignore file so it is not accidentally added to git. This file will automatically be loaded by teleirc into environmental variables.
6. Edit `.env` and update `server`, `botName`, `channel`, and `chatId` for your uses. Change the token configuration value to the bot token you received from the BotFather and put it in the value for `TELEIRC_TOKEN`. Other optional values include:
 * _server_: IRC server you wish to connect to (default: irc.freenode.net)
 * _botName_: Nickname for your bot to use on IRC
 * _channel_: IRC channel you want your bot to join
 * _chatId_: Telegram chat ID of the group you are bridging ([how do I get this?](http://stackoverflow.com/a/32572159))
6. Optional settings are available to customize teleirc to respond to extra irc events such as join, leave, kick, action messages. You can also set a custom prefix and suffix that will be inserted when a messages from Telegram is sent to IRC.

Alternatively, if you start up the bot with no Telegram chat ID set, it will sit waiting for messages to be sent to it. If you invite the bot to your group chat, you should see a "Debug TG" message with some information about the invite that was sent. One of the fields here will be the chatId. This is the value that needs to be put in the config object. Be careful not to get the user ID of a specific user when reading these messages.

``` javascript
let settings = {
    token: process.env.TELEIRC_TOKEN || "000000000:AAAAAAaAAa2AaAAaoAAAA-a_aaAAaAaaaAA",
    ircBlacklist: process.env.IRC_BLACKLIST ?
        process.env.IRC_BLACKLIST.split(",") : [],
    irc: {
        server: process.env.IRC_SERVER || "irc.freenode.net",
        channel: process.env.IRC_CHANNEL || "",
        botName: process.env.IRC_BOT_NAME || "teleirc",
        sendStickerEmoji: process.env.IRC_SEND_STICKER_EMOJI || true,
        prefix: process.env.IRC_PREFIX || "<",
        suffix: process.env.IRC_SUFFIX || ">",
        showJoinMessage: process.env.IRC_SHOW_JOIN_MESSAGE || true,
        showLeaveMessage: process.env.IRC_SHOW_LEAVE_MESSAGE || true,
    },
    tg: {
        chatId: process.env.TELEGRAM_CHAT_ID,
        showJoinMessage: process.env.SHOW_JOIN_MESSAGE || false,
        showActionMessage: process.env.SHOW_ACTION_MESSAGE || true,
        showLeaveMessage: process.env.SHOW_LEAVE_MESSAGE || false,
        showKickMessage: process.env.SHOW_KICK_MESSAGE || false,
        maxMessagesPerMinute: process.env.MAX_MESSAGES_PER_MINUTE || 20,
    }
}

module.exports = settings;
```

### IRC

There is not real configuration needed on the IRC side, as IRC is generally very open. It might be a good idea to [register the channel](https://infrastructure.fedoraproject.org/infra/docs/freenode-irc-channel.rst) you are using.


# Running Teleirc

Before running teleirc, you will need to decide how you want to run it persistently. Several options are available:

* pm2: An easy way to keep this service running in the background on your server is through `pm2`. pm2 is an npm package that can keep node services running in the background, restart them if they crash, and restart them if the server reboots.
* systemd: A systemd service file is provided (`teleirc.service`) which can be installed into `/usr/lib/systemd/system` and managed through standard systemd commands. The systemd service assumes you've created a dedicated user called `teleirc` with the home directory `/var/lib/teleirc` which contains teleirc.js, config.js, and node_modules.

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
