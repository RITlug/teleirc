Contributing to TeleIRC
=======================

<!--
    Style rule: one sentence per line please!
    This makes git diffs easier to read.
-->

This is a guide on how to contribute to the TeleIRC project.
It explicitly defines working practices of the development team.
The goal of this document is to help new contributors get up to speed with working on the project.
It is a living document and may change.
If you think something could be better, please [open an issue](https://github.com/RITlug/teleirc/issues/new) with your feedback.


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

* [Nodejs](https://nodejs.org/en/) (v10+ preferred)
* Telegram account
* IRC client ([HexChat](https://hexchat.github.io/) recommended)
* For docs: [Python 3](https://www.python.org/downloads/) (3.6+ preferred)

### Create Telegram bot

Create a Telegram bot using the Telegram [BotFather](https://t.me/botfather).
See [TeleIRC documentation](https://docs.teleirc.com/en/latest/quick-install/#create-a-telegram-bot) for more instructions on how to do this.

### Create Telegram group

Create a new Telegram group for testing.
Invite the bot user as another member to the group.
Configure the Telegram bot to TeleIRC specifications before adding it to the group.

### Register IRC channel

Registering an IRC channel is encouraged, but optional.
At the least, you need an unused IRC channel to use for testing.
Registering the channel gives you additional privileges as a channel operator (e.g. testing NickServ authentication to join private IRC channels).
See your IRC network's documentation on registering a channel.

### Configure and run TeleIRC

Change the `env.example` file to `.env`.
Change the configuration values to the Telegram bot's tokens.
For more help with configuration, see the [TeleIRC documentation](https://docs.teleirc.com/en/latest/quick-install/#configure-and-run-teleirc).


## Open a new pull request

These guidelines help maintainers review new pull requests.
Stick to the guidelines for quicker and easier pull request reviews.

1. Prefer gradual small changes than sudden big changes
2. Write a helpful title for your pull request (if someone reads only one sentence, will they understand your change?)
3. Address the following questions in your pull request:
    1. What is a summary of your change?
    2. Why is this change helpful?
    3. Any specific details to consider?
    4. What do you think is the outcome of this change?


## Maintainer response time

Project maintainers are committed to **no more than 10 days for a reply** to a new ticket.
Current maintainers are volunteers working on the project, so we try to keep up with the project as best we can.
If more than 10 days passed and you have not received a reply, follow up in [Telegram](https://t.me/teleirc) or [IRC](https://webchat.freenode.net/?channels=ritlug-teleirc) (`#ritlug-teleirc` on irc.freenode.net).
Someone may have missed your comment – we are not intentionally ignoring anyone.

_Remember_, using issue templates and answering the above questions in new pull requests likely reduces response time from a maintainer to your ticket / PR.
