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
