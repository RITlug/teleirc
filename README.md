Teleirc [![Build Status](https://travis-ci.org/RITlug/teleirc.svg?branch=devel)](https://travis-ci.org/RITlug/teleirc)
=======

NodeJS implementation of a [Telegram](https://telegram.org/) <=>
[IRC](https://en.wikipedia.org/wiki/Internet_Relay_Chat) bridge for use with
any IRC channel and Telegram group


## About

**Teleirc** helps bridge communication between a Telegram group and IRC
channels. The project is a lightweight NodeJS application.

This bot was originally written for [RITlug](http://ritlug.com) and
our own Telegram group and IRC channel. Today, it is used by various
communities other than just RITlug.

### Live demo

A public Telegram supergroup and IRC channel (on freenode) are available for
you to test. Our (small) development community is also found in these channels
too.

* **[Telegram](https://telegram.me/teleirc)**
* **[IRC](https://webchat.freenode.net/?channels=ritlug-teleirc)** (#ritlug-teleirc @ irc.freenode.net)

### Who uses Teleirc?

See what projects and communities use Teleirc on the
[wiki page](https://github.com/RITlug/teleirc/wiki/Who-uses-Teleirc%3F "Who uses Teleirc?").


## Installation

There are three parts to configuring Teleirc.

1. Creating a Telegram bot
2. Configuring Teleirc
3. Setting up the IRC channel

### Telegram

Create a new Telegram bot to act as the bridge from the Telegram side. You will
need a token key for the bot and you will also use the bot to discover the
unique chat ID of the Telegram you are adding it to.

#### Create the bot

1. Send `/start` in a message to the @BotFather¹ user on Telegram
2. Follow instructions to create a new bot (e.g. name, username, etc.)
3. Receive token key for new bot (used for accessing Telegram API)
4. _Optional_: Add descriptions / profile picture for your bot on Telegram
5. **Required**: Change `/setprivacy` to _Disabled_ so bot can see messages²
6. Add bot to Telegram group you want to bridge
7. _Optional_: After adding, block bot from being added to more groups
(`/setjoingroups`)

You have now set up the Telegram side for Teleirc.

¹ @BotFather is the [Telegram bot](https://core.telegram.org/bots) for
[creating Telegram bots](https://core.telegram.org/bots#6-botfather).

² Privacy setting must be changed for the bot to see messages in the Telegram
group. By default, bots can't see messages unless they use a command to
interact with the bot. Since Teleirc is forwarding all messages, it needs to
see all the messages, which is why this setting has to be changed.

### Teleirc

On an ArchLinux based distro, teleirc can be installed [from the AUR](https://aur.archlinux.org/packages/teleirc/). You will still need to
follow the configuration steps from below. If you install from the AUR, teleirc's files will be located at `/var/lib/teleirc` and teleirc
can be managed via systemd.

#### Run using Docker

To get teleirc working, you will need a server to run it on and a recent version of Docker installed.

##### Which image do I choose?

Official Node 6 Slim, Official Node 6 Alpine Linux, and Fedora and images are provided.

Their sizes may be deal breakers (ordered ascending by size):

| **Image**                                                                                                    | **Size** |
|--------------------------------------------------------------------------------------------------------------|----------|
| [Official Node 6 Alpine Linux](https://hub.docker.com/r/_/node/)(`node:6-alpine`) -- See `Dockerfile.alpine` | 66.9 MB  |
| [Offical Node 6 Slim](https://hub.docker.com/r/_/node/) (`node:6-slim`) -- See `Dockerfile.slim`             | 256 MB   |
| [Official Fedora Latest](https://hub.docker.com/r/_/fedora/) -- See `Dockerfile.fedora`                      | 569 MB   |

The below example uses alpine, replace `alpine` with `fedora`, or `slim` for the `node:6-slim` base, if you so choose.

You will see errors during `npm install`, ignore them. They are not fatal.

```bash
docker build . -f Dockerfile.alpine -t teleirc
docker run -d -u teleirc --name teleirc --restart always \
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
    user: teleirc
    build:
      context: .
      dockerfile: Dockerfile.alpine
    env_file: .env
```

We provide an example compose file (`docker-compose.yml.example`).

You can optionally, tell Docker Compose to use the Fedora or the official `node:6-slim` image as well by simply replacing the `Dockerfile.alpine` with `Dockerfile.fedora` or `Dockerfile.slim`. It is worth noting that the Fedora image is rather large at over **560 MB** and thus we recommend you use either `Dockerfile.alpine` or `Dockerfile.slim`.

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

Depending on your network, there is no real configuration needed on the IRC
side. There are a few good ideas you can use to set up your IRC channel:

* [Register your channel](https://infrastructure.fedoraproject.org/infra/docs/freenode-irc-channel.rst)
* Give permanent voice to your bridge bot (for most networks, the `+Vv` flags)


## Usage

First, choose how you want to run Teleirc persistently. There are several options.

* **pm2**: `pm2` is a NPM package that…
    * Keeps NodeJS services running in background
    * Restarts them if they crash
    * Restarts them if server reboots
* **systemd**: Use provided systemd service file (`teleirc.service`) to run
    * Install provided file into `/usr/lib/systemd/system/`
    * Manage Teleirc through standard `systemctl` commands
    * `teleirc.service` assumes…
        * Using a dedicated user (`teleirc`)
        * Home directory at `/var/lib/teleirc/` with `teleirc.js`, `config.js`,
`node_modules`, etc.
* **screen / tmux**: Manually run or use a shell script to start Teleirc in a
persistent window


## License

Teleirc is provided under the
[MIT License](https://github.com/RITlug/teleirc/blob/master/LICENSE). If you're
hacking on teleirc, we'd love to see you submit improvements back upstream!

