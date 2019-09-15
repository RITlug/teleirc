package internal

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

type TelegramSettings struct {
	ChatID              string
	ShowJoinMessage     bool
	ShowActionMessage   bool
	ShowLeaveMessage    bool
	ShowKickMessage     bool
	MaxMessagePerMinute int
}

type ImgurSettings struct {
	UseImgurForImageLinks bool
	ImgurClientID         string
}

type Settings struct {
	Token        string
	IRCBlacklist []string
	IRC          IRCSettings
	Telegram     TelegramSettings
	Imgur        ImgurSettings
}
