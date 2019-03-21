package locale

import (
	"os"

	"github.com/ecletus/ecletus"
	"github.com/ecletus/plug"
)

type Plugin struct {
	ConfigDirKey  string
	LocaleKey     string
	DefaultLocale string
}

func (p *Plugin) ProvideOptions() []string {
	return []string{p.LocaleKey}
}

func (p *Plugin) RequireOptions() []string {
	return []string{p.ConfigDirKey}
}

func (p *Plugin) Init(options *plug.Options) error {
	var l = &Locale{}
	config_dir := options.GetInterface(p.ConfigDirKey).(*ecletus.ConfigDir)
	if err := config_dir.Load(l, "locale.yaml"); err != nil {
		if !os.IsNotExist(err) {
			return err
		}
	}
	options.Set(p.LocaleKey, l.SetDefaults())
	return nil
}
