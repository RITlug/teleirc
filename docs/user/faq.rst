##########################
Frequently asked questions
##########################

This page collects frequently asked scenarios or problems with TeleIRC.
Did you find something confusing?
Please let us know in our developer chat or open a new pull request with a suggestion!


***
IRC
***

Messages do not appear in the IRC channel. Why?
===============================================

There are a lot of things that *could* be the cause.
However, make sure the **IRC channel is surrounded by quotes in the `.env` file**:

```
IRC_CHANNEL="#my-cool-channel"
```

If there are no quotes, the value is interpreted as a comment, or the same thing as if it were an empty string.


********
Telegram
********

.. _chat-id:

How do I find a chat ID for a Telegram group?
=============================================

There are two ways we suggest finding the chat ID of a Telegram group.

The easiest way is to add the `@getidsbot <https://t.me/getidsbot>`_ to the group.
As soon as the bot joins the group, it will print a message with the group chat ID.
You can remove the bot once you get the chat ID.

.. image:: /_static/about/faq-getidsbot.png
   :alt: Screenshot of sample message when adding @getidsbot to a group

Another way to get the chat ID is from the Telegram API via a web browser.
First, add your bot to the group.
Then, open a browser and enter the Telegram API URL with your API token, as explained in `this post <https://stackoverflow.com/questions/32423837/telegram-bot-how-to-get-a-group-chat-id/32572159#32572159>`_.
Next, send a message in the group that @tags the bot username, and refresh the browser window.
You will see the chat ID for the Telegram group along with other information.

I reinstalled TeleIRC after it was inactive for a while. But the bot doesn't work. Why?
=======================================================================================

If a Telegram bot is not used for a while, it "goes to sleep".
Even if TeleIRC is configured and installed correctly, you need to "wake up" the bot.
To fix this, *remove the bot from the group and add it again*.
Restart TeleIRC and it should work again.

.. _disable-privacy:

Why do I have to disable privacy on the Telegram bot during setup?
==================================================================

The privacy setting must be disabled for TeleIRC bot to "see" messages in the Telegram group.
By default, bots cannot see messages unless a person uses a command to interact directly with a bot.
Since TeleIRC forwards all sent messages from Telegram to IRC, it must see all messages to work.

Messages are not stored or tracked by TeleIRC (but may optionally be logged by an administrator).


*****
Imgur
*****

.. _imgur-setup:

How do I get an Imgur API Client ID?
====================================

If you are bridging to a busy telegram group, please register your own application.

Visit Imgur's `application registration page <https://api.imgur.com/oauth2/addclient>`
to create a new application. Choose a name, set the type to "OAuth 2
authorization without a callback URL", and enter your email address and a
description for the app. From the next page, copy the Client ID into your
configuration file. If you intend to upload to an account (rather than
anonymously), also copy your client secret.

These can be retrieved later from your `applications page <https://imgur.com/account/settings/apps>`

.. _imgur-login:

How do I upload to my Imgur account, instead of anonymously?
============================================================

If you only configure a Client ID, uploaded images will not be linked with any
account. If you want them to belong to an account, in addition to the Client ID
above you will need to put your app's **client secret** into the configuration
file and connect TeleIRC to your Imgur account using OAuth2. Replace the
`client_id` parameter in this link with yours and then visit it
to authorize the app.

```
https://api.imgur.com/oauth2/authorize?response_type=token&client_id=YOUR_CLIENT_ID
```

You will then be redirected to the imgur homepage, with some additional
parameters in the URL:

```
https://imgur.com/#access_token=e3b0...c442&expires_in=315360000&token_type=bearer&refresh_token=98fc...1c14&account_username=...
```

From these parameters copy the string of characters after `refresh_token=`, in
this case `98fc...1c14`, into the `IMGUR_REFRESH_TOKEN` variable in your
configuration file, then restart the bridge.

.. _imgur-album:

How do I find the album hash?
=============================

The easiest way is by signing in to Imgur, clicking your name at the top right,
clicking "Images", opening your browser's developer tools, switching to the
"Network" tab, and then choosing the album you want from the dropdown on the
page.

In the network inspector will be a request for a URL like the following:

```
https://yourname.imgur.com/ajax/images?sort=3&order=0&album=r7bbAyI&page=1&perPage=60
```

In this example, the album hash is `r7bbAyI`.
