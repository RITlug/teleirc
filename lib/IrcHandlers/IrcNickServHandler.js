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