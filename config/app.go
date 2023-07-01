package config

// App - структура конфигурации приложения.
type App struct {
	Name    string `yaml:"name" env:"NAME"` //todo test prefix
	Version string `yaml:"version" env:"VERSION" env-default:"1.0.0"`
}
