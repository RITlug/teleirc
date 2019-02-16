#############
Quick install
#############

This is a quick installation guide for an administrator to configure and deploy Teleirc.
Teleirc configuration is divided into these steps:

#. Create a Telegram bot
#. Configure IRC channel
#. Configure and run Teleirc


*********************
Create a Telegram bot
*********************

.. note:: Teleirc **DOES NOT** support channels, only groups.
          Read more about channels vs groups `here <https://telegram.org/faq#q-what-39s-the-difference-between-groups-supergroups-and-channel>`_.

Create a new Telegram bot to act as a bridge from the Telegram side.
The bot API provides a token key for the bot.
Use the bot to discover the unique chat ID of your Telegram group.

Create bot with BotFather
=========================

#. Send ``/start`` to @BotFather [#]_ user on Telegram
#. Follow instructions to create new bot (e.g. name, username, description, etc.)
#. Receive **token key** for new bot (used to access Telegram API)
#. (**REQUIRED**) Set ``/setprivacy`` to **DISABLED** (so bot can see messages) [#]_
#. Add bot to Telegram group you plan to bridge

Optional configuration changes
==============================

#. Set description or profile picture for your bot with @BotFather
#. Block your bot from being added to more groups (``/setjoingroups``)


*********************
Configure IRC channel
*********************

There is no required configuration for an IRC channel.
However, there are recommendations for best practices:

#. `Register your channel <https://docs.pagure.org/infra-docs/sysadmin-guide/sops/freenode-irc-channel.html#adding-new-channel>`_
#. Give permanent voice to your bridge bot via **ChanServ** (most networks use the ``+V`` flag)
    - *Example*: On freenode, ``/query ChanServ ACCESS #channel ADD my-teleirc-bot +V``


*************************
Configure and run Teleirc
*************************

This section explains how to configure and install Teleirc itself.

Requirements
============

- git
- nodejs (v8 and v10 supported)
- `yarn <https://yarnpkg.com/en/docs/install>`_

Install dependencies
====================

#. Clone the repository (``git clone https://github.com/RITlug/teleirc.git``)
#. Install dependencies (``yarn``)

Configuration
=============

Teleirc uses `dotenv <https://www.npmjs.com/package/dotenv>`_ to manage API keys and settings.
The config file you use is a ``.env`` file.
Copy the example file to a production file to get started (``cp env.example .env``).
Edit the ``.env`` file with your API keys and settings.

.. seealso::

   See :doc:`config-file-glossary` for detailed information.

Relay Telegram picture messages via Imgur
-----------------------------------------

Teleirc retrieves picture messages via the Telegram API.
By default, picture messages from Telegram are sent to IRC through Imgur.
`See context <https://github.com/RITlug/teleirc/issues/115>`_ for why Imgur is enabled by default.

.. note:: By default, Teleirc uses the generic Imgur API key.
          Imgur highly recommends registering each Teleirc bot.

To add Imgur support, follow these steps:

#. Create an Imgur account
#. `Register your bot <https://api.imgur.com/oauth2/addclient>`_ with the Imgur API
    - Select *OAuth2 without callback* option
#. Put client ID into ``.env`` file


.. [#] @BotFather is the `Telegram bot <https://core.telegram.org/bots>`_ for `creating Telegram bots <https://core.telegram.org/bots#6-botfather>`_
.. [#] Privacy setting must be disabled for Teleirc bot to see messages in the Telegram group.
       By default, bots cannot see messages unless a person uses a command to interact with the bot.
       Since Teleirc forwards all messages, it needs to see all messages.
       Messages are not stored or tracked by Teleirc.
