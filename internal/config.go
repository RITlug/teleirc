package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/caarlos0/env/v6"
	"github.com/go-playground/validator"
	"github.com/joho/godotenv"
)

var validate *validator.Validate

const defaultPath = ".env"

// IRCSettings includes settings related to the IRC bot/message relaying
type IRCSettings struct {
	BindAddress         string   `env:"IRC_HOST_IP" envDefault:""`
	Server              string   `env:"IRC_SERVER,required"`
	ServerPass          string   `env:"IRC_SERVER_PASSWORD" envDefault:""`
	Port                int      `env:"IRC_PORT" envDefault:"6667" validate:"min=0,max=65535"`
	TLSAllowSelfSigned  bool     `env:"IRC_CERT_ALLOW_SELFSIGNED" envDefault:"true"`
	TLSAllowCertExpired bool     `env:"IRC_CERT_ALLOW_EXPIRED" envDefault:"true"`
	Channel             string   `env:"IRC_CHANNEL,required" validate:"notempty"`
	ChannelKey          string   `env:"IRC_CHANNEL_KEY" envDefault:""`
	BotIdent            string   `env:"IRC_BOT_IDENT,required" envDefault:"teleirc"`
	BotName             string   `env:"IRC_BOT_REALNAME" envDefault:"Powered by TeleIRC <github.com/RITlug/teleirc>"`
	BotNick             string   `env:"IRC_BOT_NAME,required" validate:"notempty"`
	SendStickerEmoji    bool     `env:"IRC_SEND_STICKER_EMOJI" envDefault:"true"`
	SendDocument        bool     `env:"IRC_SEND_DOCUMENT" envDefault:"false"`
	SendPhoto           bool     `env:"IRC_SEND_PHOTO" envDefault:"true"`
	Prefix              string   `env:"IRC_PREFIX" envDefault:"<"`
	Suffix              string   `env:"IRC_SUFFIX" envDefault:">"`
	ShowJoinMessage     bool     `env:"IRC_SHOW_JOIN_MESSAGE" envDefault:"true"`
	ShowLeaveMessage    bool     `env:"IRC_SHOW_LEAVE_MESSAGE" envDefault:"true"`
	ShowZWSP            bool     `env:"IRC_SHOW_ZWSP" envDefault:"true"`
	ShowLocationMessage bool     `env:"IRC_SHOW_LOCATION_MESSAGE" envDefault:"false"`
	NickServUser        string   `env:"IRC_NICKSERV_USER" envDefault:""`
	NickServPassword    string   `env:"IRC_NICKSERV_PASS" envDefault:""`
	NickServService     string   `env:"IRC_NICKSERV_SERVICE" envDefault:""`
	EditedPrefix        string   `env:"IRC_EDITED_PREFIX" envDefault:"[EDIT] "`
	MaxMessageLength    int      `env:"IRC_MAX_MESSAGE_LENGTH" envDefault:"400"`
	IRCBlacklist        []string `env:"IRC_BLACKLIST" envDefault:"[]string{}"`
	UseSSL              bool     `env:"IRC_USE_SSL" envDefault:"false"`
	NoForwardPrefix     string   `env:"IRC_NO_FORWARD_PREFIX" envDefault:""`
	QuitMessage         string   `env:"IRC_QUIT_MESSAGE" envDefault:""`
}

// TelegramSettings includes settings related to the Telegram bot/message relaying
type TelegramSettings struct {
	Token                 string   `env:"TELEIRC_TOKEN,required"`
	ChatID                int64    `env:"TELEGRAM_CHAT_ID,required"`
	Prefix                string   `env:"TELEGRAM_MESSAGE_PREFIX" envDefault:"<"`
	Suffix                string   `env:"TELEGRAM_MESSAGE_SUFFIX" envDefault:">"`
	ReplyPrefix           string   `env:"TELEGRAM_MESSAGE_REPLY_PREFIX" envDefault:"["`
	ReplySuffix           string   `env:"TELEGRAM_MESSAGE_REPLY_SUFFIX" envDefault:"]"`
	ReplyLength           int      `env:"TELEGRAM_MESSAGE_REPLY_LENGTH" envDefault:"15"`
	ShowTopicMessage      bool     `env:"SHOW_TOPIC_MESSAGE" envDefault:"false"`
	ShowJoinMessage       bool     `env:"SHOW_JOIN_MESSAGE" envDefault:"false"`
	JoinMessageAllowList  []string `env:"JOIN_MESSAGE_ALLOW_LIST" envDefault:"[]string{}"`
	ShowActionMessage     bool     `env:"SHOW_ACTION_MESSAGE" envDefault:"true"`
	ShowLeaveMessage      bool     `env:"SHOW_LEAVE_MESSAGE" envDefault:"false"`
	LeaveMessageAllowList []string `env:"LEAVE_MESSAGE_ALLOW_LIST" envDefault:"[]string{}"`
	ShowKickMessage       bool     `env:"SHOW_KICK_MESSAGE" envDefault:"false"`
	ShowNickMessage       bool     `env:"SHOW_NICK_MESSAGE" envDefault:"false"`
	ShowDisconnectMessage bool     `env:"SHOW_DISCONNECT_MESSAGE" envDefault:"false"`
	MaxMessagePerMinute   int      `env:"MAX_MESSAGE_PER_MINUTE" envDefault:"20"`
	PreferName            bool     `env:"PREFER_FIRSTNAME" envDefault:"false"`
	QuoteNick             bool     `env:"QUOTE_NICKNAME" envDefault:"false"`
}

// ImgurSettings includes settings related to Imgur uploading for Telegram photos
type ImgurSettings struct {
	ImgurClientID     string `env:"IMGUR_CLIENT_ID" envDefault:"7d6b00b87043f58"`
	ImgurClientSecret string `env:"IMGUR_CLIENT_SECRET" envDefault:""`
	ImgurRefreshToken string `env:"IMGUR_REFRESH_TOKEN" envDefault:""`
	ImgurAccessToken  string ``
	ImgurAlbumHash    string `env:"IMGUR_ALBUM_HASH" envDefault:""`
}

// Settings includes all user-configurable settings for TeleIRC
type Settings struct {
	IRC      IRCSettings
	Telegram TelegramSettings
	Imgur    ImgurSettings
}

func validateEmptyString(fl validator.FieldLevel) bool {
	return fl.Field().String() != ""
}

// ConfigErrors lets us wrap the validator errors in a type we can return
type ConfigErrors []validator.FieldError

func (ce ConfigErrors) Error() string {
	finalStr := ""
	for _, err := range ce {
		switch err.Tag() {
		case "notempty":
			finalStr += fmt.Sprintf("Field %s was an empty string. "+
				"Perhaps you had a # and need to surround the value with \"\"?\n", err.Namespace())
		case "min":
			finalStr += fmt.Sprintf("Field %s failed to validate: %s too small.\n",
				err.Namespace(), err.Param())
		case "max":
			finalStr += fmt.Sprintf("Field %s failed to validate: %s too large.\n",
				err.Namespace(), err.Param())
		default:
			finalStr += fmt.Sprintf("Field %s failed to validate: %s failed %s.\n",
				err.Namespace(), err.Param(), err.Tag())
		}
	}
	return finalStr
}

/*
For environment variables that are meant to be a list separated by whitespace,
we need to do that manually.   This function does just that.
*/
func splitEnvVar(list []string) []string {
	if list == nil {
		return list
	} else if len(list) == 0 {
		return list
	}

	return strings.Split(list[0], " ")
}

/*
LoadConfig loads in the .env file in the provided path (or ".env" by default)
If the user-provided config is valid, return a new Settings struct that contains these settings.
Otherwise, return the error that caused the failure.
*/
func LoadConfig(path string) (*Settings, error) {
	validate = validator.New()
	if err := validate.RegisterValidation("notempty", validateEmptyString); err != nil {
		return nil, err
	}
	// If a path was provided, try to load it.
	if path != "" {
		if info, err := os.Stat(path); err == nil {
			// If the path is a directory, look for <path>/.env.
			if info.IsDir() {
				envFile := filepath.Join(path, defaultPath)
				if _, err := os.Stat(envFile); err == nil {
					if err := godotenv.Load(envFile); err != nil {
						return nil, err
					}
				}
			} else {
				// path exists and is a file — attempt to load it
				if err := godotenv.Load(path); err != nil {
					return nil, err
				}
			}
		} else {
			// If the provided path does not exist, continue and rely on passed process ENV variables
			if os.IsNotExist(err) {
				warning.Printf("config path %q not provided or does not exist; continuing and using process environment variables", path)
			} else {
				return nil, err
			}
		}
	} else if _, err := os.Stat(defaultPath); !os.IsNotExist(err) {
		// Attempt to load from defaultPath if it exists
		if err := godotenv.Load(defaultPath); err != nil {
			return nil, err
		}
	}
	settings := &Settings{}
	if err := env.Parse(settings); err != nil {
		return nil, err
	}
	if err := validate.Struct(settings); err != nil {
		fieldErrs := ConfigErrors{}
		for _, errs := range err.(validator.ValidationErrors) {
			fieldErrs = append(fieldErrs, errs)
		}
		return nil, fieldErrs
	}

	settings.IRC.IRCBlacklist = splitEnvVar(settings.IRC.IRCBlacklist)
	settings.Telegram.JoinMessageAllowList = splitEnvVar(settings.Telegram.JoinMessageAllowList)
	settings.Telegram.LeaveMessageAllowList = splitEnvVar(settings.Telegram.LeaveMessageAllowList)

	return settings, nil
}
