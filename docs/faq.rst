##########################
Frequently asked questions
##########################

This page collects frequently asked scenarios or problems with TeleIRC.
Did you find something confusing?
Please let us know in our developer chat or open a new pull request with a suggestion!


********
Telegram
********

**How do I find a chat ID for a Telegram group?**
    The chat ID is found when viewing results from the Telegram API from a web browser.
    First, add your bot to the group.
    Then, open a browser and enter the Telegram API URL with your API token, as explained in `this post <https://stackoverflow.com/questions/32423837/telegram-bot-how-to-get-a-group-chat-id/32572159#32572159>`_.
    Next, send a message in the group with the bot username and refresh the browser window.
    You will see the chat ID for the Telegram group along with other information.

**I reinstalled TeleIRC after it was inactive for a while. But the bot doesn't work. Why?**
    If a Telegram bot is not used for a while, it "goes to sleep".
    Even if TeleIRC is configured and installed correctly, you need to "wake up" the bot.
    To fix this, *remove the bot from the group and add it again*.
    Restart TeleIRC and it should work again.
