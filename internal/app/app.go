package app

import (
	"github.com/BudaShest/go-cleanest-template/config"
	"github.com/sirupsen/logrus"
)

// App - структура приложения.
type App struct {
	*config.Config
}

// New - создаёт структуру приложения.
func New(conf *config.Config) *App {
	return &App{
		conf,
	}
}

// Run - запускает приложение.
func (app *App) Run() {
	if err := app.initialize(); err != nil {
		logrus.Fatalf("application components initialization error: %w", err) //todo test %s
	}
}

// initialize - инициализириует компоненты приложения.
func (app *App) initialize() error {
	return nil
}
