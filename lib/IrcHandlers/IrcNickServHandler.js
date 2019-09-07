/*
The MIT License (MIT)

Copyright (c) 2016 RIT Linux Users Group

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/



/**
 * Handles the event when a user is kicked from the IRC channel.
 */
class IrcNickServHandler {
    
    // ---------------- Constructor ----------------

    /**
     * 
     * @param {*} ircConfig - The bot's IRC Configuration.
     * @param {*} ircBot - The IRC Bot object (needed to join channels).
     */
    constructor(ircConfig, ircBot) {
        this._ircConfig = ircConfig;
        this._ircBot = ircBot;
        
        // If there is no nickserv service to watch, don't
        // even bother with this handler.
        this._enabled = this._ircConfig.nickservService.length > 0;
    }

    // ---------------- Function ----------------

    /**
     * 
     * @param {string} username - The user who sent the message.
     * @param {string} channel - The channel the user has sent the message in.
     * @param {string} userMessage - The message the user sent.
     */
    HandleNickServ(username, channel, userMessage) {
        if (this._enabled == false) {
            return;
        }

        if (username.toLowerCase() === this._ircConfig.nickservService.toLowerCase()) {
            console.log(`[${this._ircConfig.nickservService}] ${userMessage}`);

            if (
                (this._ircBot.chans !== undefined) &&
                (this._ircBot.chans.indexOf(this._ircConfig.channel) == -1)) {
                // we are not in the channel - let's try to join it, hoping that this
                // message was good news about the bot being identified
                console.log("Bot not on the specified channel - trying to join...");
                this._ircBot.join(this._ircConfig.channel);
            }
        }
    }
}

module.exports = IrcNickServHandler;