package internal

// IRCSettings includes settings related to the IRC bot/message relaying
type IRCSettings struct {
	Server              string
	Port                int
	TLSAllowSelfSigned  bool
	TLSAllowCertExpired bool
	Channel             string
	BotName             string
	SendStickerEmoji    bool
	SendDocument        bool
	Prefix              string
	Suffix              string
	ShowJoinMessage     bool
	ShowLeaveMessage    bool
	NickServPassword    string
	NickServService     string
	EditedPrefix        string
	MaxMessageLength    int
}

// TelegramSettings includes settings related to the Telegram bot/message relaying
type TelegramSettings struct {
	ChatID              string
	ShowJoinMessage     bool
	ShowActionMessage   bool
	ShowLeaveMessage    bool
	ShowKickMessage     bool
	MaxMessagePerMinute int
}

// ImgurSettings includes settings related to Imgur uploading for Telegram photos
type ImgurSettings struct {
	UseImgurForImageLinks bool
	ImgurClientID         string
}

// Settings includes all user-configurable settings for TeleIRC
type Settings struct {
	Token        string
	IRCBlacklist []string
	IRC          IRCSettings
	Telegram     TelegramSettings
	Imgur        ImgurSettings
}
