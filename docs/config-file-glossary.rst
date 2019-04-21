####################
Config file glossary
####################

This page is a glossary of different settings in the ``env.example`` configuration file.
The values shown for the settings are their defaults.
This glossary is intended for advanced users.


************
IRC settings
************

``IRC_BLACKLIST=""``
    Comma-separated list of IRC nicks to ignore

``IRC_BOT_NAME=teleirc``
    IRC nickname for bot

``IRC_CHANNEL=#channel``
    IRC channel for bot to join

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

``IRC_NICKSERV_SERVICE=NickServ``
    IRC service used for authentication

``IRC_NICKSERV_PASS=""``
    IRC account password to complete IRC authentication

``IRC_EDITED_PREFIX="[EDIT] "``
    Prefix to include when a user edits a Telegram message and it is resent to IRC

``IRC_MAX_MESSAGE_LENGTH=400``
    Maximum length of the message that can be sent to IRC.
    Longer messages will be split into multiple messages.


*****************
Telegram settings
*****************

``TELEGRAM_CHAT_ID=-0000000000000``
    Telegram chat ID of bridged group (`how do I get this? <http://stackoverflow.com/a/32572159>`_)

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

``IMGUR_CLIENT_ID=0000000000``
    Imgur API client ID value to access Imgur API


**********************
Miscellaneous settings
**********************

``NTBA_FIX_319=1``
    Required to fix a bug in a library used by TeleIRC.
    For context, see `yagop/node-telegram-bot-api#319 <https://github.com/yagop/node-telegram-bot-api/issues/319#issuecomment-324963294>`_.
