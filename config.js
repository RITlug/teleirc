let settings = {
  token: process.env.TELEIRC_TOKEN ||
    "000000000:AAAAAAaAAa2AaAAaoAAAA-a_aaAAaAaaaAA",
  ircBlacklist: process.env.IRC_BLACKLIST ?
    process.env.IRC_BLACKLIST.split(",") : [],
  irc: {
    server: process.env.IRC_SERVER || "irc.freenode.net",
    port: process.env.IRC_PORT || 6667,
    tlsAllowSelfSigned: process.env.IRC_CERT_ALLOW_SELFSIGNED === "true" || false,
    tlsAllowCertExpired: process.env.IRC_CERT_ALLOW_EXPIRED === "true" || false,
    channel: process.env.IRC_CHANNEL || "",
    botName: process.env.IRC_BOT_NAME || "teleirc",
    sendStickerEmoji: process.env.IRC_SEND_STICKER_EMOJI === "true" || true,
    sendDocument: process.env.IRC_SEND_DOCUMENT === "true" || false,
    prefix: process.env.IRC_PREFIX || "<",
    suffix: process.env.IRC_SUFFIX || ">",
    showJoinMessage: process.env.IRC_SHOW_JOIN_MESSAGE === "true" || true,
    showLeaveMessage: process.env.IRC_SHOW_LEAVE_MESSAGE === "true" || true,
    nickservPassword: process.env.IRC_NICKSERV_PASS || "",
    nickservService: process.env.IRC_NICKSERV_SERVICE || "",
    editedPrefix: process.env.IRC_EDITED_PREFIX || "[EDIT]",
    maxMessageLength: Number(process.env.IRC_MAX_MESSAGE_LENGTH) || 400,
  },
  tg: {
    chatId: process.env.TELEGRAM_CHAT_ID || "-000000000",
    showJoinMessage: process.env.SHOW_JOIN_MESSAGE === "true" || false,
    showActionMessage: process.env.SHOW_ACTION_MESSAGE === "true" || true,
    showLeaveMessage: process.env.SHOW_LEAVE_MESSAGE === "true" || false,
    showKickMessage: process.env.SHOW_KICK_MESSAGE === "true" || false,
    maxMessagesPerMinute: Number(process.env.MAX_MESSAGES_PER_MINUTE) || 20,
  },
  imgur: {
    useImgurForImageLinks: process.env.USE_IMGUR_FOR_IMAGES === "true" ||
      false,
    imgurClientId: process.env.IMGUR_CLIENT_ID || "12345",
  }
}

module.exports = settings;
