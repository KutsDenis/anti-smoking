package config

import (
	"github.com/KutsDenis/logpac"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"sync"
)

// Config содержит параметры бота Telegram
type Config struct {
	Bot struct {
		Token string `yaml:"token"`
		Owner int64  `yaml:"owner"`
		Debug bool   `yaml:"debug"`
	} `yaml:"bot"`
}

// Get содержит конфигурацию
var Get *Config

// onceCfg используется для инициализации конфигурации
var onceCfg sync.Once

// GetConfig получает конфигурацию из файла.
// Файл конфигурации имеет формат YAML и содержит параметры бота Telegram
func GetConfig(l *logpac.Logger) {
	configPath := "config.yml"

	onceCfg.Do(func() {
		Get = &Config{}

		_, err := os.Stat(configPath)
		if os.IsNotExist(err) {
			configPath = "../../" + configPath
		}

		if err = cleanenv.ReadConfig(configPath, Get); err != nil {
			l.Fatalf("%s", err)
		}
	})
}
