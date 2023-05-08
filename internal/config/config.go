package config

import (
	"github.com/KutsDenis/logpac"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

// Config contains the Telegram bot parameters
type Config struct {
	Bot struct {
		Token    string `yaml:"token"`
		Owner    int64  `yaml:"owner"`
		TimeZone int8   `yaml:"time_zone"`
		Debug    bool   `yaml:"debug"`
	} `yaml:"bot"`
}

// Get contains the configuration
var Get *Config

// onceCfg is used for initializing the configuration
var onceCfg sync.Once

const configPath = "configs/config.yml"

// GetConfig get the configuration from a file.
// The configuration file has a YAML format and contains the Telegram bot parameters,
// such as the token, administrator and group identifiers, as well as the timezone
func GetConfig(l *logpac.Logger) {
	onceCfg.Do(func() {
		Get = &Config{}

		if err := cleanenv.ReadConfig(configPath, Get); err != nil {
			l.Fatalf("%s", err)
		}
	})
}
