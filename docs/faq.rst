##########################
Frequently asked questions
##########################

This page collects frequently asked scenarios or problems with TeleIRC.
Did you find something confusing?
Please let us know in our developer chat or open a new pull request with a suggestion!


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
