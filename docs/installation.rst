############
Installation
############

**NOTE: This page is being reworked into smaller pieces.**
See `this GitHub issue <https://github.com/RITlug/teleirc/issues/118>`_.
Some parts of this page may be confusing or difficult to understand until the refactor is complete.

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


*************************
Configure and run Teleirc
*************************

.. note:: teleirc **DOES NOT** support channels, only groups. Read more about channels vs groups `here <https://telegram.org/faq#q-what-39s-the-difference-between-groups-supergroups-and-channel>`_.

This section explains how to install and configure Teleirc itself.

Requirements
============

- git
- nodejs

  - Node 8 and Node 10 are officially supported

Download
========

#. Clone the repository (``git clone https://github.com/RITlug/teleirc.git``)
#. Install `yarn <https://yarnpkg.com/en/docs/install>`_
#. Install dependencies (``yarn``)

Configuration
=============

Teleirc uses `dotenv <https://www.npmjs.com/package/dotenv>`_ to manage API keys and settings.
All configuration changes live in the ``.env`` file.
This makes it possible to use ``git pull`` to upgrade the bot in-place.

Explaining config file
----------------------

The config file you use is a ``.env`` file.
All configuration values for Teleirc are stored there.
Copy the example file to a production file to get started (``cp env.example .env``)
Edit the ``.env`` file to your preference.

.. seealso::

   See :doc:`config-file-glossary` for detailed information

Imgur support
-------------

Teleirc retrieves images via the Telegram API.
By default, picture messages from Telegram link to the Telegram API URL.
However, links expire and are not reliable.
Optionally, Teleirc can upload an image to Imgur and replace the Telegram API URL with a link to Imgur.
This makes picture messages more durable for logs or someone joining the conversation later.

To add Imgur support, follow these steps:

#. Create an Imgur account, if you do not have one

#. `Register your bot <https://api.imgur.com/oauth2/addclient>`_ with the Imgur API

    - Select *OAuth2 without callback* option

#. Put client ID into ``.env`` file and enable using Imgur


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
