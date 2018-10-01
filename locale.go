package locale

import (
	"os"
	"strings"
)

func DefaultLocale() string {
	locale := os.Getenv("LANG")
	if locale != "" {
		locale = strings.Split(locale, ".")[0]
	}
	return locale
}

var DEFAULT_LOCALE = DefaultLocale()

type Locale struct {
	Default string
}

func (l *Locale) SetDefaults() *Locale {
	if l.Default == "" {
		l.Default = DEFAULT_LOCALE
	}
	return l
}
