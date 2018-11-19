/**
 * Class that handles errors that can occur when using Telegram.
 */
class TgErrorHandler {

    // ----------------- Constructor -----------------

    constructor(enabled, action) {
        this.Enabled = enabled;
        this._action = action;
    }

    // ----------------- Functions -----------------

    /**
     * Use this handler when we get a chat id we were not expecting.
     * @param {*} chat - The chat object whose ID is invalid.
     * @param {*} msg - The message object that we received.
     */
    HandleBadChatId(chat, msg) {
        if (!this.Enabled) {
            return;
        }

        // Messages that are sent to the bot outside of the group chat should just be dumped
        // to the console for potential testing and debugging to do things like check chat IDs
        // and verify the JSON formats of various messages
        let message = "[TG Debug] - Unexpected chat ID '" + chat.id + "' from message: " + JSON.stringify(msg);
        this._action(message);
    }

    /**
     * Use this handler when we get a message whose type (e.g. sticker, photo, etc)
     * is unknown.
     * @param {*} msg - The message object that we received.
     */
    HandleUnknownMessageType(msg) {
        if (!this.Enabled) {
            return;
        }

        let message = "[TG Debug] Ignoring non-text message: " + JSON.stringify(msg);
        this._action(message);
    }
}

module.exports = TgErrorHandler;
