####################
Config file glossary
####################

This page is a glossary of different settings in the ``env.example`` configuration file.
All values shown are the default settings.
This glossary is intended for advanced users.


************
IRC settings
************

Host connection settings
========================

``IRC_HOST_IP=""``
    Specify IP address to use for IRC connection.
    Useful for hosts with multiple IP addresses.

Server connection settings
==============================

``IRC_SERVER=chat.freenode.net``
    IRC server to connect to

``IRC_SERVER_PASSWORD=""``
    IRC server password

``IRC_PORT=6697``
    IRC server port

Encryption options
------------------

``IRC_USE_SSL=true``
    Connect to the IRC server with SSL

``IRC_CERT_ALLOW_EXPIRED=false``
    Allow connecting to IRC server with an expired TLS/SSL certificate

``IRC_CERT_ALLOW_SELFSIGNED=false``
    Allows TeleIRC to accept TLS/SSL certificates from non-trusted/unknown Certificate Authorities (CA)

Channel settings
================

``IRC_CHANNEL="#channel"``
    IRC channel for bot to join

``IRC_CHANNEL_KEY=""``
    IRC channel key, for password-protected channels

``IRC_BLACKLIST=""``
    Comma-separated list of IRC nicks to ignore

Bot settings
============

``IRC_BOT_NAME=teleirc``
    IRC nickname for bot.
    Most IRC clients and bridges show this nickname.

``IRC_BOT_REALNAME="Powered by TeleIRC <github.com/RITlug/teleirc>"``
    IRC ``REALNAME`` for bot.
    Often visible in IRC ``/whois`` reports, but most clients do not show this information by default.

``IRC_BOT_IDENT=teleirc``
    Identification metadata for bot connection.
    This is rarely used or shown in most IRC clients and bridges, so only change this if you know what you are doing.
    Often visible in IRC ``/whois`` reports, but most clients do not show this information by default.

NickServ options
----------------

``IRC_NICKSERV_SERVICE=NickServ``
    IRC service used for authentication.

``IRC_NICKSERV_USER=""``
    IRC NickServ username.

``IRC_NICKSERV_PASS=""``
    IRC NickServ password.

Message settings
================

``IRC_PREFIX="<"``
    Text displayed before Telegram name in IRC

``IRC_SUFFIX=">"``
    Text displayed after Telegram name in IRC

``IRC_SEND_STICKER_EMOJI=true``
    Send emojis associated with a sticker to IRC (when a Telegram user sends a sticker)

``IRC_SEND_DOCUMENT=false``
    Send documents and files from Telegram to IRC (`why is this false by default? <https://github.com/RITlug/teleirc/issues/115>`_)

``IRC_EDITED_PREFIX="(edited) "``
    Prefix to prepend to messages when a user edits a Telegram message and it is resent to IRC

``IRC_MAX_MESSAGE_LENGTH=400``
    Maximum length of the message that can be sent to IRC.
    Longer messages are split into multiple messages.

``IRC_SHOW_ZWSP=true``
    Prevents users with the same Telegram and IRC username from pinging themselves across platforms.

``IRC_NO_FORWARD_PREFIX="[off]"``
    A string users can prefix their message with to prevent it from being relayed across the bridge.
    Removing this option or setting it to "" disables it.

``IRC_QUIT_MESSAGE="TeleIRC bridge stopped."``
    A string that TeleIRC sends to the IRC channel when the application exits using IRC's "QUIT" command.
    If not specified, it simply closes the connection without providing a reason.
    The bot must be connected to a server for a certain amount of time for the server to send the quit message to the channel.


*****************
Telegram settings
*****************

``TELEGRAM_CHAT_ID=-0000000000000``
    Telegram chat ID of bridged group (:ref:`how do I get this? <chat-id>`).

``TELEIRC_TOKEN=000000000:AAAAAAaAAa2AaAAaoAAAA-a_aaAAaAaaaAA``
    Private API token for Telegram bot.

``MAX_MESSAGES_PER_MINUTE=20``
    Maximum number of messages sent to Telegram from IRC per minute.

``TELEGRAM_MESSAGE_REPLY_PREFIX="["
    Prefix separator for Telegram reply

``TELEGRAM_MESSAGE_REPLY_SUFFIX="]"
    Suffix separator for Telegram reply

``TELEGRAM_MESSAGE_REPLY_LENGTH=15
    Length of quoted reply message before truncation

``SHOW_TOPIC_MESSAGE=true``
    Send Telegram message when the topic in the IRC channel is changed.

``SHOW_ACTION_MESSAGE=true``
    Relay action messages (e.g. ``/me thinks TeleIRC is cool!``).

``SHOW_JOIN_MESSAGE=false``
    Send Telegram message when someone joins IRC channel.

``JOIN_MESSAGE_ALLOW_LIST=""``
    List of users (separated by a space character) whose IRC leave messages will be sent to Telegram, even if SHOW_JOIN_MESSAGE is false.
    This is ignored if SHOW_JOIN_MESSAGE is set to true.

``SHOW_KICK_MESSAGE=true``
    Send Telegram message when someone is kicked from IRC channel.

``SHOW_NICK_MESSAGE=false``
    Send Telegram message when someone changes their nickname in the IRC channel.

``SHOW_LEAVE_MESSAGE=false``
    Send Telegram message when someone leaves IRC channel either by quitting or parting.

``LEAVE_MESSAGE_ALLOW_LIST=""``
    List of users (separated by a space character) whose IRC leave messages will be sent to Telegram, even if SHOW_LEAVE_MESSAGE is false.
    This is ignored if SHOW_LEAVE_MESSAGE is set to true.

``SHOW_DISCONNECT_MESSAGE=true``
    Sends a message to Telegram when the bot disconnects from the IRC side.

**************
Imgur settings
**************

``IMGUR_CLIENT_ID=7d6b00b87043f58``
    Imgur API client ID value to access Imgur API.
    Uses a default API key.
    If you are bridging to a very active Telegram group, *please register your own API key*.
