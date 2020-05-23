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
    1. [Run binary](#run-binary)
        1. [Requirements](#requirements)
        1. [Install dependencies](#install-dependencies)
        1. [Configuration](#configuration)
        1. [Start bot](#start-bot)
    1. [Run container](#run-container)


## Overview

This section is a written, high-level overview of how to configure and deploy a TeleIRC bot.
The Quick Start Guide will cover these topics:

1. Create a Telegram bot to obtain a Telegram API token
1. Set up an IRC channel for best user experience
1. Configure TeleIRC with your Telegram bot and IRC channel
1. Deploy TeleIRC publicly in one of many officially-supported pathways

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
1. `/query ChanServ ACCESS #channel ADD <NickServ account or hostmask of TeleIRC bot> +V`

### Configure Imgur Image Upload (IIU)

**NOTE**: The Imgur Image Upload (IIU) feature is not yet available in v2.x.x releases.

_By default_, TeleIRC uploads images sent to the Telegram group to [Imgur][7].
Since IRC does not support images, Imgur is an intermediary approach to sending pictures sent on Telegram over to IRC.
Note that images will be publicly visible on the Internet if the URL is known.
[See context][8] for why Imgur is enabled by default.

By default, TeleIRC uses the generic Imgur API key.
Imgur highly recommends registering each bot.

To register your own Imgur API key, follow these steps:

1. Create an Imgur account
1. [Register your bot][9] with Imgur API (select *OAuth2 without callback* option)
1. Add provided Imgur client ID to `.env` file


## Deployment Guide

There are two ways to deploy TeleIRC persistently:

1. Run Go binary
1. Run TeleIRC in a container


### Run binary

This section explains how to configure and install TeleIRC as a simple executable binary.

#### Requirements

- git
- go (v1.13 and v1.14 supported)

#### Install dependencies

1. Clone the repository (`git clone https://github.com/RITlug/teleirc.git`)
1. Install dependencies (`go install`)

#### Configuration

TeleIRC uses [godotenv][10] to manage API keys and settings.
The config file is a `.env` file.
Copy the example file to a production file to get started (`cp env.example .env`).
Edit the `.env` file with your API keys and settings.

See [_Config file glossary_][11] for detailed information.

#### Start bot

<!-- #TODO @jwflory -->


## Run container

Containers are another way to deploy TeleIRC.
Dockerfiles and other deployment resources are available in ``deployments/``.

**Coming soon.**

<!-- #TODO
Needs better documentation by someone with experience running and using the container image for TeleIRC v2.
-->


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
