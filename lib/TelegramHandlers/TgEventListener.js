const EventEmitter = require('events');

/**
 * This class takes in a Telegram Message object,
 * and fires the corresponding event.
 *
 * It doesn't appear as though the library we are using
 * is able to distinguish the *type* of Telegram message
 * we received, so this class does this for us.
 */
class TgEventListener extends EventEmitter {

    // ---------------- Constructor ----------------

    /**
     *
     * @param {string} chatId - The group ID of the chat.
     */
    constructor({chatId, ircEditedPrefix}) {
        super();
        this._chatId = chatId;
        this._ircEditPrefix = ircEditedPrefix;
    }

    // ---------------- Functions ----------------

    /**
     * Parses the given Telegram message and fires the correct Telegram event
     * that is subscribed to.
     */
    ParseMessage(msg) {
        // If the message's group chat ID does not match
        // the bots, fire a special event.
        if (msg.chat.id != this._chatId) { // Use != instead of !==, as _chatId could be a string from config.js.
            super.emit('bad_chat_id', msg.chat, msg);
        }

        // There are several types of Telegram messages.
        // Let's start with an easy one... if the Message's text is defined,
        // it means it is a message that we received.
        else if (msg.text) {
            super.emit('message', msg.from, msg.text);
        }

        // If there is no text defined, it can be ANY of the following
        // Telegram message types:

        // If new_chat_member is defined, it means someone
        // joined the Telegram group.
        else if (msg.new_chat_member) {
            super.emit('join', msg.new_chat_member);
        }

        // If left_chat_member is defined,
        // it means someone left the Telegram group.
        // The event name is called 'part' to be consistent
        // with the IRC event.
        else if (msg.left_chat_member) {
            super.emit('part', msg.left_chat_member);
        }

        // If there is a sticker, it means we have a sticker message.
        else if (msg.sticker) {
            super.emit('sticker', msg.from, msg.sticker);
        }

        // Having a photo means we have a photo message.
        else if (msg.photo) {
            super.emit('photo', msg.from, msg.photo, msg.caption);
        }

        // And having a document means we have a document message.
        else if (msg.document) {
            super.emit('document', msg.from, msg.document);
        }

        // Everything else is a message type we do not support yet.
        else {
            super.emit('unknown', msg);
        }
    }

    ParseEditedMessage(msg) {
      if (msg.text) {
        return this.ParseMessage(Object.assign({}, msg, {
          text: `${this._ircEditPrefix}${msg.text}`
        }));
      } else {
        return this.ParseMessage(msg);
      }
    }
}

module.exports = TgEventListener;
