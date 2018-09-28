#################
Quick Guide
#################

This quick guide will list the most basic steps to get teleirc up and running.
For more complete instructions, as well as various options, take a look at the full installation guide.

These are the general steps, and order, that you should follow in your initial setup.


#. Create a Telegram Bot using the BotFather in Telegram

#. Install Nodejs, git, and Yarn

#. Clone the Telirc Repository to your system

   - run ``git clone https://github.com/RITlug/teleirc.git``

#. Change into the teleirc directory

   - run ``cd teleirc``

#. Run 'yarn' to download dependencies

   - run ``yarn``

#. Copy the example configuration to the real configuration file 

   - run ``cp .env.example .env``

#. Fill in the .env configuration file with settings for your instance

   - Required: IRC_CHANNEL, IRC_SERVER, TELEIRC_TOKEN, TELEGRAM_CHAT_ID


#. Run teleirc 

   - ``node teleirc.js``


