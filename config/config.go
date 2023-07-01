package config

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/sirupsen/logrus"
)

// Config - родительская структура конфигурации.
type Config struct {
	App `yaml:"app" env-prefix:"APP"` //todo test-prefix
}

// New - создаёт и вохвращает структуру конфигурации.
// Пытается считать конфигурационный файл, при его отсутсвии считаывает переменные окружения.
func New(configPath string) (*Config, error) {
	cfg := &Config{}

	if err := cleanenv.ReadConfig(configPath, cfg); err != nil {
		logrus.Errorf("read configuration file error: %#v", err)
		if err = cleanenv.ReadEnv(cfg); err != nil {
			return nil, fmt.Errorf("read env. variables error: %#v", err)
		}
	}

	return cfg, nil
}
