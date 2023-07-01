package cmd

import (
	"github.com/BudaShest/go-cleanest-template/config"
	"github.com/BudaShest/go-cleanest-template/internal/app"
	"github.com/sirupsen/logrus"
)

const configPath = "./config/config.go"

// main - точка входа.
func main() {
	cfg, err := config.New(configPath)
	if err != nil {
		logrus.Fatalf("Configuration error: %w", err) //todo test w
	}

	application := app.New(cfg)
	application.Run()
}
