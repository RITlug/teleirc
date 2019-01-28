############
Installation
############

This page is an installation guide for a server administrator to install and run Teleirc.
The person who installs Teleirc will have elevated privileges to maintain and administrate the Telegram bot.

There are three parts to installing and configuring Teleirc:

#. Create a Telegram bot
#. Configure and run Teleirc
#. Configure IRC channel


*********************
Create a Telegram bot
*********************

Create a new Telegram bot to act as a bridge from the Telegram side.
From the bot API, you will receive a token key for the bot.
You will then use the bot obtain the unique chat ID of your Telegram group.

Create bot with API
===================

#. Send ``/start`` in a message to @BotFather [#]_ user on Telegram
#. Follow instructions from @BotFather to create new bot (e.g. name, username, description, etc.)
#. Receive **token key** for new bot (used to access Telegram API)
#. (*Optional*) Set description or profile picture for your bot with @BotFather
#. (**Required**) Set ``/setprivacy`` to **DISABLED** (so bot can see messages) [#]_
#. Add bot to Telegram group you plan to bridge
#. (*Optional*) Block your bot from being added to more groups (``/setjoingroups``)

All configuration on Telegram side is complete.


*************************
Configure and run Teleirc
*************************

There are several installation options available.
This guide details the supported method to configure and run Teleirc.
For alternative options, read further below (for Dockerfiles, distro-specific packages, etc.).

Requirements
============

- git
- nodejs

  - Node 8 and Node 10 are officially supported

Download
========

#. Clone the repository to the server

  - ``git clone https://github.com/RITlug/teleirc.git``

#. Change directories to repo

#. Install `yarn <https://yarnpkg.com/en/docs/install>`_ via one of the following methods (we will only Linux, `yarn install documentation <https://yarnpkg.com/en/docs/install>`_ cover the rest)

   - Arch Linux

    .. code-block:: shell

        pacman -S yarn

   - Debian/Ubuntu

    .. code-block:: shell

        curl -sS https://dl.yarnpkg.com/debian/pubkey.gpg | sudo apt-key add -
        echo "deb https://dl.yarnpkg.com/debian/ stable main" | sudo tee /etc/apt/sources.list.d/yarn.list
        sudo apt-get update ; sudo apt-get install --no-install-recommends yarn

   - Fedora/CentOS/RHEL:

    .. code-block:: shell

        curl --silent --location https://dl.yarnpkg.com/rpm/yarn.repo | sudo tee /etc/yum.repos.d/yarn.repo
        yum install yarn or dnf install yarn


#. Install dependencies using yarn

   - ``yarn``

Configuration
=============

Teleirc uses `dotenv <https://www.npmjs.com/package/dotenv>`_ for easy management of API keys and settings.
All your configuration changes will live in the ``.env`` file.
You should not need to change other files.
This makes it possible to use ``git pull`` to upgrade the bot in-place.

Explaining config file
----------------------

The config file you use is a ``.env`` file.
All configuration values for Teleirc are stored here.
Copy the example file to a productive file to get started (``cp .env.example .env``)
Edit the ``.env`` file and configure it your preference.

Explanations of the configuration values are below.

IRC
^^^

- **IRC_BLACKLIST**:

    - Comma-separated list of IRC nicks to ignore (default: ``""``)

- **IRC_BOT_NAME**:

    - Nickname for your bot to use on IRC (default: ``teleirc``)

- **IRC_CHANNEL**:

    - IRC channel you want your bot to join (default: ``#channel``)

- **IRC_SEND_STICKER_EMOJI**:

    - Send the emoji associated with a sticker on IRC (default: ``true``)

- **IRC_PREFIX**:

    - Text displayed before Telegram name in IRC (default: ``"<"``)

- **IRC_SUFFIX**:

    - Text displayed after Telegram name in IRC (default: ``">"``)

- **IRC_SERVER**:

    - IRC server you wish to connect to (default: ``chat.freenode.net``)

- **IRC_NICKSERV_SERVICE**:
    - IRC service you would like to use to authenticate with IRC (default: ``NickServ``)

- **IRC_NICKSERV_PASS**:
    - IRC password for your bot to use in order to complete IRC authentication (default: ``""``)

- **IRC_MAX_MESSAGE_LENGTH**:
    - Maximum length of the message that can be sent to IRC. Longer messages
      will be split into multiple messages. (default: ``400``)

Telegram
^^^^^^^^

.. note:: teleirc **DOES NOT** support channels, only groups. Read more about channels vs groups `here <https://telegram.org/faq#q-what-39s-the-difference-between-groups-supergroups-and-channel>`_.


- **TELEIRC_TOKEN**:

    - Private API token for Telegram bot

- **MAX_MESSAGES_PER_MINUTE**:

    - Maximum rate at which to relay messages (default: ``20``)

- **SHOW_ACTION_MESSAGE**:

    - Relay action messages (default: ``true``)

- **SHOW_JOIN_MESSAGE**:

    - Relay join messages (default: ``false``)

- **SHOW_KICK_MESSAGE**:

    - Relay kick messages (default: ``false``)

- **SHOW_LEAVE_MESSAGE**:

    - Relay leave messages (default: ``false``)

- **TELEGRAM_CHAT_ID**:

    - Telegram chat ID of the group you are bridging (`how do I get this? <http://stackoverflow.com/a/32572159>`_)

Imgur
^^^^^

- **USE_IMGUR_FOR_IMAGES**:

    - Upload picture messages from Telegram to Imgur, convert picture to Imgur link in IRC (default: ``false``)

- **IMGUR_CLIENT_ID**:

    - Imgur API client ID value to access Imgur API

Imgur support
-------------

Teleirc retrieves images via the Telegram API.
By default, picture messages from Telegram will link to the Telegram API URL.
However, the links expire and are not reliable.
Optionally, Teleirc can upload an image to Imgur and replace the Telegram API URL with a link to Imgur.
This makes picture messages more durable for logs or someone joining the conversation later.

To add Imgur support, follow these steps:

#. Create an Imgur account, if you do not have one

#. `Register your bot <https://api.imgur.com/oauth2/addclient>`_ with the Imgur API

    - Select *OAuth2 without callback* option

#. Put client ID into ``.env`` file and enable using Imgur

Usage
=====

Choose how you want to run Teleirc persistently.
Teleirc officially supports three methods.

pm2
---

`pm2 <http://pm2.keymetrics.io/>`_ is a NPM package that keeps NodeJS running in the background.
If you run an application and it crashes, pm2 restarts the process.
pm2 also restarts processes if the server reboots.

Read the `pm2 documentation <http://pm2.keymetrics.io/docs/usage/quick-start/>`_ for more information.

After pm2 is installed, follow these steps to start Teleirc::

    cd teleirc/
    pm2 start -n teleirc-channel teleirc.js

systemd
-------

systemd is an option to run the bot persistently.
A provided systemd service file is available (``misc/teleirc.service``)
Move the provided file to ``/usr/lib/systemd/system/`` to activate it.
Now, you can manage Teleirc through standard ``systemctl`` commands.

Note that the provided file makes two assumptions:

- Using a dedicated system user (e.g. ``teleirc``)
- Home directory located at ``/usr/lib/teleirc/`` (i.e. files inside Teleirc repository)

screen / tmux
-------------

Terminal multiplexers like `GNU screen <https://www.gnu.org/software/screen/>`_ and `tmux <https://en.wikipedia.org/wiki/Tmux>`_ let you run Teleirc persistently.
If you are not familiar with a multiplexer, read more about tmux `here <https://hackernoon.com/a-gentle-introduction-to-tmux-8d784c404340>`_.

Inside of your persistent window, follow these steps to start Teleirc::

    cd teleirc/
    node teleirc.js

ArchLinux
=========

On ArchLinux, teleirc is available `in the AUR <https://aur.archlinux.org/packages/teleirc/>`_.
The AUR package uses the systemd method for running Teleirc.
Configure the bot as detailed above in the ``/usr/lib/teleirc/`` directory.

Docker
======

.. seealso::

   See :doc:`Using Docker <using-docker>` for more information


*********************
Configure IRC channel
*********************

Depending on the IRC network you use, no configuration in IRC is required.
However, there are recommendations for best practices to follow.

- `Register your channel <https://docs.pagure.org/infra-docs/sysadmin-guide/sops/freenode-irc-channel.html>`_
- Give permanent voice to your bridge bot via **NickServ** (for most networks, the ``+Vv`` flags)
    - *Example*: For freenode, ``/query NickServ ACCESS #channel ADD my-teleirc-bot +Vv``


.. [#] @BotFather is the `Telegram bot <https://core.telegram.org/bots>`_ for `creating Telegram bots <https://core.telegram.org/bots#6-botfather>`_
.. [#] Privacy setting must be changed for the bot to see messages in the Telegram group.
       By default, bots cannot see messages unless they use a command to interact with the bot.
       Since Teleirc forwards all messages, it needs to see all messages.
       This is why this setting must be changed.
