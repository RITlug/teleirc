####################
Config file glossary
####################

This page is a glossary of different settings in the ``env.example`` configuration file.
All values shown are the default settings.
This glossary is intended for advanced users.


************
IRC settings
************

``IRC_BLACKLIST=""``
    Comma-separated list of IRC nicks to ignore

``IRC_BOT_NAME=teleirc``
    IRC nickname for bot

``IRC_CHANNEL="#channel"``
    IRC channel for bot to join

``IRC_CHANNEL_KEY=""``
    IRC channel key, for password-protected channels

``IRC_SEND_STICKER_EMOJI=true``
    Send emojis associated with a sticker to IRC (when a Telegram user sends a sticker)

``IRC_SEND_DOCUMENT=false``
    Send documents and files from Telegram to IRC (`why is this false by default? <https://github.com/RITlug/teleirc/issues/115>`_)

``IRC_PREFIX="<"``
    Text displayed before Telegram name in IRC

``IRC_SUFFIX=">"``
    Text displayed after Telegram name in IRC

``IRC_SERVER=chat.freenode.net``
    IRC server to connect to

``IRC_SERVER_PASSWORD=""``
    IRC server password

``IRC_SERVER_PORT=6697``
    IRC server port

``IRC_CERT_ALLOW_SELFSIGNED=false``
    Allows TeleIRC to accept TLS/SSL certificates from non-trusted/unknown Certificate Authorities (CA)

``IRC_CERT_ALLOW_EXPIRED=false``
    Allow connecting to IRC server with an expired TLS/SSL certificate

``IRC_NICKSERV_SERVICE=NickServ``
    IRC service used for authentication

``IRC_NICKSERV_PASS=""``
    IRC account password to complete IRC authentication

``IRC_EDITED_PREFIX="[EDIT] "``
    Prefix to prepend to messages when a user edits a Telegram message and it is resent to IRC

``IRC_MAX_MESSAGE_LENGTH=400``
    Maximum length of the message that can be sent to IRC.
    Longer messages are split into multiple messages.


*****************
Telegram settings
*****************

``TELEGRAM_CHAT_ID=-0000000000000``
    Telegram chat ID of bridged group (:ref:`how do I get this? <chat-id>`)

``TELEIRC_TOKEN=000000000:AAAAAAaAAa2AaAAaoAAAA-a_aaAAaAaaaAA``
    Private API token for Telegram bot

``MAX_MESSAGES_PER_MINUTE=20``
    Maximum number of messages sent to Telegram from IRC per minute

``SHOW_ACTION_MESSAGE=true``
    Relay action messages (e.g. ``/me thinks TeleIRC is cool!``)

``SHOW_JOIN_MESSAGE=false``
    Send Telegram message when someone joins IRC channel

``SHOW_KICK_MESSAGE=true``
    Send Telegram message when someone is kicked from IRC channel

``SHOW_LEAVE_MESSAGE=false``
    Send Telegram message when someone leaves IRC channel


**************
Imgur settings
**************

``USE_IMGUR_FOR_IMAGES=true``
    Upload picture messages from Telegram to Imgur, send Imgur link to IRC

``IMGUR_CLIENT_ID=7d6b00b87043f58``
    Imgur API client ID value to access Imgur API.
    Uses a default API key.
    If you are bridging to a very active Telegram group, *please register your own API key*.
