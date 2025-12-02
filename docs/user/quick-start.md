# Quick Start Guide

This page is a quick start guide to using TeleIRC.
It is an overview of how to install, set up, and deploy TeleIRC v2.x.x releases.
Note this does not apply to v1.x.x releases; see the [v1.3.4 documentation][1].


## Contents

1. [Overview](#overview)
    1. [Telegram Channels](#telegram-channels)
1. [Create a Telegram bot](#create-a-telegram-bot)
    1. [Create bot with BotFather](#create-bot-with-botfather)
    1. [Optional BotFather tweaks](#optional-botfather-tweaks)
1. [Configure IRC channel](#configure-irc-channel)
    1. [IRC channel overview](#irc-channel-overview)
    1. [Configure a Freenode IRC channel](#configure-a-freenode-irc-channel)
    1. [Configure Imgur Image Upload (IIU)](#configure-imgur-image-upload-iiu)
1. [Deployment Guide](#deployment-guide)
    1. [Run container](#run-container)
    1. [Run binary](#run-binary)
        1. [Pre-requirements](#pre-requirements)
        1. [Build TeleIRC](#build-teleirc)
        1. [Configuration](#configuration)
        1. [Start bot](#start-bot)
            1. [Example Linux setup](#example-linux-setup)


## Overview

This section is a written, high-level overview of how to configure and deploy a TeleIRC bot.
The Quick Start Guide will cover these topics:

1. [Create a Telegram bot to obtain a Telegram API token](#create-a-telegram-bot)
1. [Set up an IRC channel for best user experience](#configure-irc-channel)
1. [Deploy TeleIRC to your system](#deployment-guide)

**It is important each step is followed exactly and in order.**
Missing a step or skipping a section often results in common frustrations, such as one-way relay of chat messages.

### Telegram Channels

TeleIRC **DOES NOT** support Telegram Channels, only groups.
Read more on the differences between channels and groups in the [Telegram FAQ][2].


## Create a Telegram bot

TeleIRC requires a Telegram API token in order to access messages in a Telegram group.
To obtain a token, someone must register a new Telegram bot.
The Telegram bot will appear as the Sending User in Telegram for all IRC messages.

### Create bot with BotFather

BotFather is the Telegram bot for [creating Telegram bots][3].
[**See the official Telegram documentation for how to create a new bot.**][4]

Once you create a new bot, you _must_ follow these additional steps (**IN EXACT ORDER**) for TeleIRC:

1. Send `/setprivacy` to @BotFather, change to **Disable** [_why?_][5]
1. Add bot to Telegram group to be bridged
1. Send `/setjoingroups` to @BotFather, change to **Disable** [_why?_][6]

### Optional BotFather tweaks

1. Set a description or add profile picture for your bot
1. Block your bot from being added to more groups (``/setjoingroups``)


## Configure IRC channel

This section explains best practices for configuring IRC channels for TeleIRC.
Because IRC networks can run different software, exact instructions may differ depending on your IRC network.
So, this section is divided in two ways:

1. High-level overview of how to set up your IRC channel
1. How to actually do it on Freenode IRC network

### IRC channel overview

No matter what IRC network you use, TeleIRC developers recommend this IRC channel configuration:

* Register your IRC channel with the IRC network (i.e. `ChanServ`).
* Unauthenticated users may join the channel.
* Only authenticated users may write in the channel (i.e. `NickServ`).
* Any user connecting from a network-recognized gateway (e.g. web chat) with an assigned hostmask automatically receives voice on join (and thus, does not need to authenticate to write in channel).
* TeleIRC bot hostmask automatically receives voice on join (and thus, does not need to authenticate to write in channel).
* IRC channel operators automatically receive operator privilege on join.

### Configure a Freenode IRC channel

If your IRC channel is on the Freenode IRC network, use these exact commands to create a channel policy as described above:

1. `/join #channel`
1. `/query ChanServ REGISTER #channel`
1. `/query ChanServ SET #channel GUARD on`
1. `/query ChanServ ACCESS #channel ADD <NickServ account> +AORfiorstv` (_repeat for each IRC user who needs admin access_) ([_what do these mean?_][6])
1. `/query ChanServ SET mlock #channel +Ccnt`
1. `/mode #channel +q $~a`
1. `/query ChanServ ACCESS #channel ADD *!*@gateway/* +V`
1. `/query ChanServ ACCESS #channel ADD *!*@freenode/staff/* +Aiotv`
1. `/query ChanServ ACCESS #channel ADD <bot NickServ account or hostmask> +V`

### Adjust default Imgur Image Upload (IIU)

Since IRC does not support images, Imgur is an intermediary approach to sending pictures sent on Telegram over to IRC.

> [!IMPORTANT]
> _By default_, all images the Telegram Bot reads are uploaded by TeleIRC to [Imgur][7].
> [See context][8] for why Imgur upload is enabled by default.

By default, TeleIRC uses the TeleIRC-registered Imgur API key.
We highly recommend registering your own API key in high-traffic channels.
Otherwise, API rate limiting can occur.

#### Optionally disable all Imgur image uploads from Telegram

* Set `IRC_SEND_PHOTO` to `false` in your `.env` file

#### Alternatively use your own Imgur API details
To register your own Imgur API key, follow these steps:

1. Create an Imgur account
1. [Register your bot][9] with Imgur API (select *OAuth2 without callback* option)
1. Add provided Imgur client ID to `.env` file


## Deployment Guide

There are two ways to deploy TeleIRC persistently:

1. Run TeleIRC in a container
1. Run Go binary


### Run container

Containers are the easiest way to deploy TeleIRC.
Dockerfiles and other deployment resources are available in ``deployments/``.

Ensure you have [docker](https://www.docker.com/) installed.

#### Build TeleIRC docker image

1. Enter container deployment directory (`cd deployments/container`)
1. Build image (`./build_image.sh`)
1. Run container (`docker run teleirc:latest`)

> [!NOTE]
> This deployment can optionally copy a standalone .env file


#### Run TeleIRC using Docker compose

1. Enter container deployment directory (`cd deployments/container`)
1. Run service using `docker compose`:

```bash
IRC_SERVER=chat.freenode.net \
IRC_CHANNEL='#channelname' \
IRC_BOT_NAME='teleirc' \
TELEIRC_TOKEN='000000000:AAAAAAaAAa2AaAAaoAAAA-a_aaAAaAaaaAA' \
TELEGRAM_CHAT_ID='-0000000000000' \
docker compose up -d teleirc
```

> [!TIP]
> Instead you can also add `environment:` entries via `docker-compose.yml`, or pass a standalone `.env` file using the CLI:
> `docker compose --env-file ../../.env up --build -d teleirc`


### Run binary

This section explains how to configure and install TeleIRC as a simple executable binary.

**NOTE**:
**This assumes you are building from source.**
If you use a pre-built binary from a [GitHub Release][14], skip to [_Configuration_](#configuration).

#### Pre-requirements

- git
- go (v1.15 through v1.22 supported)

Packages for these pre-requirements are available on most `*NIX` distributions.
Check your distribution documentation for more info on how to install these packages.

#### Build TeleIRC

This section is only required if you are building a binary from source:

1. Clone repository (`git clone https://github.com/RITlug/teleirc.git`)
1. Enter repository (`cd teleirc/`)
1. Build binary (`./build_binary.sh`)

#### Configuration

TeleIRC uses [godotenv][10] to manage API keys and settings.
The config file is a `.env` file.
Copy the example file to a production file to get started (`cp env.example .env`).
Edit the `.env` file with your API keys and settings.

See [_Config file glossary_][11] for detailed information.

#### Start bot

**NOTE**:
This section is one opinionated way to start and configure TeleIRC.
Experienced system administrators may have other preferences and slight deviation is permittable.
However _upstream only offers free support for installations that follow our documentation_.

To start the bot, you need to consider the following factors:

1. Where will the binary go?
1. Where is your config file on the system?
1. How will you automate the bot to start-up automatically after a system reboot?


##### Example Linux setup

**NOTE**:
_Looking for an easier way?_
_Check out the [TeleIRC Ansible Role][17] for an automated installation of the following steps._

Upstream offers a [systemd unit file][15] to automate TeleIRC on a Linux system that uses [systemd][16].
This example uses the upstream systemd unit file to automatically run TeleIRC on a Linux system.

This example was tested on a CentOS 8 system and is easily adaptable for other `*NIX` distributions.
It uses `v2.2.1` as a default:

```sh
# Download TeleIRC deployment assets from GitHub.
$ curl --location --output ~/teleirc https://github.com/RITlug/teleirc/releases/download/v2.2.1/teleirc-2.2.1-linux-x86_64
$ curl --location --output ~/teleirc.sysusers https://raw.githubusercontent.com/RITlug/teleirc/v2.2.1/deployments/systemd/teleirc.sysusers
$ curl --location --output ~/teleirc.tmpfiles https://raw.githubusercontent.com/RITlug/teleirc/v2.2.1/deployments/systemd/teleirc.tmpfiles
$ curl --location --output ~/teleirc@.service https://raw.githubusercontent.com/RITlug/teleirc/v2.2.1/deployments/systemd/teleirc@.service
$ curl --location --output ~/teleirc.env https://raw.githubusercontent.com/RITlug/teleirc/v2.2.1/env.example

# Install TeleIRC files and user
$ sudo install -Dm755 -o root -g root ~/teleirc /usr/local/bin/teleirc
$ sudo install -Dm644 -o root -g root ~/teleirc.sysusers /etc/sysusers.d/teleirc.conf
$ sudo install -Dm644 -o root -g root ~/teleirc.tmpfiles /etc/tmpfiles.d/teleirc.conf
$ sudo install -Dm644 -o root -g root ~/teleirc@.service /etc/systemd/system/teleirc@.service
$ sudo systemd-sysusers /etc/sysusers.d/teleirc.conf
$ sudo systemd-tmpfiles --create /etc/tmpfiles.d/teleirc.conf
$ sudo install -Dm644 -o root -g root ~/teleirc.env /etc/teleirc/example
$ rm ~/teleirc ~/teleirc.sysusers ~/teleirc.tmpfiles ~/teleirc@.service ~/teleirc.env

# Systems with SELinux ONLY.
$ sudo chcon --type bin_t --user system_u /usr/local/bin/teleirc
$ sudo chcon --type etc_t --user system_u -R /etc/teleirc/
$ sudo chcon --type etc_t --user system_u /etc/sysusers.d/teleirc.conf
$ sudo chcon --type etc_t --user system_u /etc/tmpfiles.d/teleirc.conf
$ sudo chcon --type systemd_unit_file_t --user system_u /etc/systemd/system/teleirc@.service

# Start and enable TeleIRC.
$ sudo systemctl enable --now teleirc@example.service
```

To run multiple instances, create other files in /etc/teleirc/ and enable the
service named teleirc@FILENAME.service.


[1]: /en/v1.3.4/
[2]: https://telegram.org/faq#q-what-39s-the-difference-between-groups-supergroups-and-channel
[3]: https://core.telegram.org/bots#6-botfather
[4]: https://core.telegram.org/bots#creating-a-new-bot
[5]: /en/latest/user/faq/
[6]: https://web.archive.org/web/20200515153553/http://toxin.jottit.com/freenode_chanserv_commands
[7]: https://imgur.com/
[8]: https://github.com/RITlug/teleirc/issues/115
[9]: https://api.imgur.com/oauth2/addclient
[10]: https://github.com/joho/godotenv
[11]: config-file-glossary
[12]: https://github.com/RITlug/teleirc/issues/new/choose
[13]: https://raw.githubusercontent.com/RITlug/teleirc/master/deployments/container/Dockerfile
[14]: https://github.com/RITlug/teleirc/releases
[15]: https://github.com/RITlug/teleirc/blob/master/deployments/systemd/teleirc@.service
[16]: https://systemd.io/
[17]: https://github.com/jwflory/ansible-role-teleirc
