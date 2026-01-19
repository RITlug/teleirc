package mediashare

import "fmt"

// Language represents a supported language code.
type Language string

const (
	LangPolish  Language = "pl"
	LangEnglish Language = "en"
)

// Translations holds all translatable strings.
var translations = map[Language]map[string]string{
	LangPolish: {
		// Landing page
		"uploaded_by":        "Wysłane przez",
		"download":           "Pobierz",
		"not_found":          "Plik nie został znaleziony",
		"not_found_desc":     "Ten plik mógł wygasnąć lub został usunięty.",
		"unsupported":        "Nieobsługiwany typ pliku",
		"unsupported_desc":   "Ten typ pliku nie może być wyświetlony w przeglądarce.",
		"powered_by":         "Obsługiwane przez",
		"uploaded_at":        "Wysłano",
		"file_info":          "Informacje o pliku",
		"direct_link":        "Link bezpośredni",
		"copy_link":          "Kopiuj link",
		"copied":             "Skopiowano!",
		"loading":            "Ładowanie...",
		"error_loading":      "Błąd ładowania pliku",
		"open_in_new_tab":    "Otwórz w nowej karcie",

		// List page
		"recent_files":       "Ostatnie pliki",
		"no_files":           "Brak plików",
		"no_files_desc":      "Nie ma jeszcze żadnych przesłanych plików.",
		"table_date":         "Data",
		"table_user":         "Użytkownik",
		"table_file":         "Plik",
		"table_link":         "Link",
		"table_last_opened":  "Ostatnio otwarte",
		"never":              "nigdy",
		"anonymous":          "anonim",

		// IRC messages (format: "user udostępnił X: URL")
		"shared_video": "%s udostępnił wideo",
		"shared_voice": "%s udostępnił wiadomość głosową",
		"shared_photo": "%s udostępnił zdjęcie",
		"shared_file":  "%s udostępnił plik",
		"with_caption": " z opisem '%s'",
	},
	LangEnglish: {
		// Landing page
		"uploaded_by":        "Uploaded by",
		"download":           "Download",
		"not_found":          "File not found",
		"not_found_desc":     "This file may have expired or been deleted.",
		"unsupported":        "Unsupported file type",
		"unsupported_desc":   "This file type cannot be displayed in the browser.",
		"powered_by":         "Powered by",
		"uploaded_at":        "Uploaded",
		"file_info":          "File information",
		"direct_link":        "Direct link",
		"copy_link":          "Copy link",
		"copied":             "Copied!",
		"loading":            "Loading...",
		"error_loading":      "Error loading file",
		"open_in_new_tab":    "Open in new tab",

		// List page
		"recent_files":       "Recent files",
		"no_files":           "No files",
		"no_files_desc":      "There are no uploaded files yet.",
		"table_date":         "Date",
		"table_user":         "User",
		"table_file":         "File",
		"table_link":         "Link",
		"table_last_opened":  "Last opened",
		"never":              "never",
		"anonymous":          "anonymous",

		// IRC messages
		"shared_video": "%s shared a video",
		"shared_voice": "%s shared a voice message",
		"shared_photo": "%s shared a photo",
		"shared_file":  "%s shared a file",
		"with_caption": " with caption '%s'",
	},
}

// I18n provides internationalization functionality.
type I18n struct {
	lang Language
}

// NewI18n creates a new I18n instance with the specified language.
// Falls back to Polish if the language is not supported.
func NewI18n(lang string) *I18n {
	l := Language(lang)
	if _, ok := translations[l]; !ok {
		l = LangPolish
	}
	return &I18n{lang: l}
}

// T returns the translation for the given key.
// Returns the key itself if translation is not found.
func (i *I18n) T(key string) string {
	if trans, ok := translations[i.lang][key]; ok {
		return trans
	}
	// Fallback to Polish
	if trans, ok := translations[LangPolish][key]; ok {
		return trans
	}
	return key
}

// Tf returns a formatted translation for the given key.
func (i *I18n) Tf(key string, args ...interface{}) string {
	format := i.T(key)
	return fmt.Sprintf(format, args...)
}

// Lang returns the current language code.
func (i *I18n) Lang() string {
	return string(i.lang)
}

// GetTranslations returns all translations for the current language.
// Useful for passing to templates.
func (i *I18n) GetTranslations() map[string]string {
	return translations[i.lang]
}

// FormatSharedMessage formats an IRC message for shared media.
func (i *I18n) FormatSharedMessage(username, mediaType, caption, url string) string {
	var key string
	switch mediaType {
	case "video":
		key = "shared_video"
	case "voice", "voice message":
		key = "shared_voice"
	case "photo":
		key = "shared_photo"
	default:
		key = "shared_file"
	}

	msg := i.Tf(key, username)
	if caption != "" {
		msg += i.Tf("with_caption", caption)
	}
	if url != "" {
		msg += ": " + url
	}
	return msg
}
