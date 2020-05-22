Contributing guidelines
=======================

<!--
    Style rule: one sentence per line please!
    This makes git diffs easier to read.
-->

This guide explains how to contribute to the TeleIRC project.
It defines working practices of the development team.
This document helps new contributors start working on the project.
It is a living document and will change.
If you think something could be better, please [open an issue](https://github.com/RITlug/teleirc/issues/new/choose) with your feedback.


## Table of contents

1. [Set up a development environment](#set-up-a-development-environment)
2. [Open a new pull request](#open-a-new-pull-request)
3. [Maintainer response time](#maintainer-response-time)


## Set up a development environment

**Contents**:

1. [Requirements](#requirements)
2. [Create Telegram bot](#create-telegram-bot)
3. [Create Telegram group](#create-telegram-group)
4. [Configure and run TeleIRC](#configure-and-run-teleirc)

### Requirements

To set up a TeleIRC development environment, you need the following:

* [Go](https://golang.org/dl/) (v1.13.x or later)
* Telegram account
* IRC client ([HexChat](https://hexchat.github.io/) recommended)
* For docs: [Python 3](https://www.python.org/downloads/) (3.6 or later)

### Create Telegram bot

Create a Telegram bot using the Telegram [BotFather](https://t.me/botfather).
See the [TeleIRC Quick Start](/user/quick-start#create-a-telegram-bot) for more instructions on how to do this.

### Create Telegram group

Create a new Telegram group for testing.
Invite the bot user as another member to the group.
Configure the Telegram bot to TeleIRC specifications before adding it to the group.

### Register IRC channel

Registering an IRC channel is encouraged, but optional.
At minimum, you need an unused IRC channel to use for testing.
Registering the channel gives you additional privileges as a channel operator (e.g. testing NickServ authentication to join private IRC channels).
See your IRC network's documentation on registering a channel.

### Configure and run TeleIRC

Change the `env.example` file to `.env`.
Change the configuration values to the Telegram bot's tokens.
For more help with configuration, see the [_Config file glossary_](/user/config-file-glossary).


## Open a pull request

These guidelines help maintainers review new pull requests.
Stick to the guidelines for quicker and easier pull request reviews.

1. We prefer gradual, small changes over sudden, big changes
1. Write a helpful title for your pull request (if someone reads only one sentence, will they understand your change?)
1. Address the following questions in your pull request:
    1. What is a short summary of your change?
    1. Why is this change helpful?
    1. Any specific details to consider?
    1. What is the desired outcome of your change?


## Maintainer response time

Project maintainers make a best effort to respond in **10 days or less** to new issues.
Current maintainers are volunteers working on the project, so we try to keep up as best we can.
If more than 10 days passed and you have not received a reply, follow up in [Telegram](https://t.me/teleirc) or [IRC](https://webchat.freenode.net/?channels=rit-lug-teleirc) (`#rit-lug-teleirc` on irc.freenode.net).
Someone may have missed your comment – we are not intentionally ignoring anyone.

_Remember_, using issue templates and answering the above questions in new pull requests reduces the response time from a maintainer to your issue or pull request.
