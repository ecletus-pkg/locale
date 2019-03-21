package locale

import (
	"os"
	"strings"
)

func DefaultLocale() string {
	locale := os.Getenv("LANG")
	if locale != "" {
		locale = strings.Replace(strings.Split(locale, ".")[0], "_", "-", 1)
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

	l.Default = strings.Replace(l.Default, "_", "-", 1)
	for i, v := range l.Languages {
		l.Languages[i] = strings.Replace(v, "_", "-", 1)
	}

	return l
}
