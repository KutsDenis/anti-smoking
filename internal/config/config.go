package config

import (
	"github.com/KutsDenis/logpac"
	"github.com/ilyakaznacheev/cleanenv"
	"sync"
)

// Config содержит параметры бота Telegram
type Config struct {
	Bot struct {
		Token string `yaml:"token"`
		Owner int64  `yaml:"owner"`
		Debug bool   `yaml:"debug"`
	} `yaml:"bot"`
	DB struct {
		Host     string `yaml:"host"`
		Port     uint8  `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	}
}

// Get содержит конфигурацию
var Get *Config

// Путь до конфиг файла
const configPath = "config.yml"

// onceCfg используется для инициализации конфигурации
var onceCfg sync.Once

// GetConfig получает конфигурацию из файла.
// Файл конфигурации имеет формат YAML и содержит параметры бота Telegram
func GetConfig(l *logpac.Logger) {
	onceCfg.Do(func() {
		Get = &Config{}

		if err := cleanenv.ReadConfig(configPath, Get); err != nil {
			l.Fatalf("%s", err)
		}
	})
}
