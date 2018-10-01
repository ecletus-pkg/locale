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
	Default   string
	Languages []string
}

func (l *Locale) SetDefaults() *Locale {
	if l.Default == "" {
		l.Default = DEFAULT_LOCALE
	}
	var hasDefault bool

	for _, lang := range l.Languages {
		if lang == l.Default {
			hasDefault = true
			break
		}
	}

	if !hasDefault {
		l.Languages = append(l.Languages, l.Default)
	}
	return l
}
