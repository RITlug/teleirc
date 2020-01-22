package internal

import (
	"os"

	"github.com/caarlos0/env/v6"
	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
)

var validate *validator.Validate

const defaultPath = ".env"

// IRCSettings includes settings related to the IRC bot/message relaying
type IRCSettings struct {
	Server              string   `env:"IRC_SERVER,required"`
	Port                int      `env:"IRC_PORT" envDefault:"6667" validate:"min=0,max=65535"`
	TLSAllowSelfSigned  bool     `env:"IRC_CERT_ALLOW_SELFSIGNED" envDefault:"true"`
	TLSAllowCertExpired bool     `env:"IRC_CERT_ALLOW_EXPIRED" envDefault:"true"`
	Channel             string   `env:"IRC_CHANNEL,required"`
	BotName             string   `env:"IRC_BOT_NAME" envDefault:"teleirc"`
	SendStickerEmoji    bool     `env:"IRC_SEND_STICKER_EMOJI" envDefault:"true"`
	SendDocument        bool     `env:"IRC_SEND_DOCUMENT" envDefault:"true"`
	Prefix              string   `env:"IRC_PREFIX" envDefault:"<"`
	Suffix              string   `env:"IRC_SUFFIX" envDefault:">"`
	ShowJoinMessage     bool     `env:"IRC_SHOW_JOIN_MESSAGE" envDefault:"true"`
	ShowLeaveMessage    bool     `env:"IRC_SHOW_LEAVE_MESSAGE" envDefault:"true"`
	NickServPassword    string   `env:"IRC_NICKSERV_PASS" envDefault:""`
	NickServService     string   `env:"IRC_NICKSERV_SERVICE" envDefault:""`
	EditedPrefix        string   `env:"IRC_EDITED_PREFIX" envDefault:"[EDIT] "`
	MaxMessageLength    int      `env:"IRC_MAX_MESSAGE_LENGTH" envDefault:"400"`
	IRCBlacklist        []string `env:"IRC_BLACKLIST" envDefault:"[]string{}"`
}

// TelegramSettings includes settings related to the Telegram bot/message relaying
type TelegramSettings struct {
	Token               string `env:"TELEIRC_TOKEN,required"`
	ChatID              int64  `env:"TELEGRAM_CHAT_ID,required"`
	ShowJoinMessage     bool   `env:"SHOW_JOIN_MESSAGE" envDefault:"false"`
	ShowActionMessage   bool   `env:"SHOW_ACTION_MESSAGE" envDefault:"false"`
	ShowLeaveMessage    bool   `env:"SHOW_LEAVE_MESSAGE" envDefault:"false"`
	ShowKickMessage     bool   `env:"SHOW_KICK_MESSAGE" envDefault:"false"`
	MaxMessagePerMinute int    `env:"MAX_MESSAGE_PER_MINUTE" envDefault:"20"`
}

// ImgurSettings includes settings related to Imgur uploading for Telegram photos
type ImgurSettings struct {
	UseImgurForImageLinks bool   `env:"USE_IMGUR_FOR_IMAGE" envDefault:"true"`
	ImgurClientID         string `env:"IMGUR_CLIENT_ID" envDefault:"7d6b00b87043f58"`
}

// Settings includes all user-configurable settings for TeleIRC
type Settings struct {
	IRC      IRCSettings
	Telegram TelegramSettings
	Imgur    ImgurSettings
}

/*
LoadConfig loads in the .env file in the provided path (or ".env" by default)
If the user-provided config is valid, return a new Settings struct that contains these settings.
Otherwise, return the error that caused the failure.
*/
func LoadConfig(path string) (*Settings, error) {
	validate = validator.New()
	// Attempt to load environment variables from path if path was provided
	if path != "" {
		if err := godotenv.Load(path); err != nil {
			return nil, err
		}
	} else if _, err := os.Stat(defaultPath); !os.IsNotExist(err) {
		// Attempt to load from defaultPath is defaultPath exists
		if err := godotenv.Load(defaultPath); err != nil {
			return nil, err
		}
	}
	settings := Settings{}
	if err := env.Parse(&settings); err != nil {
		return nil, err
	}
	if err := validate.Struct(&settings); err != nil {
		return nil, err.(validator.ValidationErrors)
	}
	return &settings, nil
}
