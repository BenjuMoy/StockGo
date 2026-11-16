package config

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
)

type Config struct {
	DataDir string
	Addr    string
}

func Load() (Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Config{}, fmt.Errorf("get home directory: %w", err)
	}

	dataDir := filepath.Join(home, ".stockgo")

	if err := os.MkdirAll(dataDir, 0o755); err != nil {
		return Config{}, fmt.Errorf("create data directory: %w", err)
	}

	slog.SetDefault(
		slog.New(
			slog.NewTextHandler(
				os.Stderr,
				&slog.HandlerOptions{
					Level: slog.LevelInfo,
				},
			),
		),
	)

	return Config{
		DataDir: dataDir,
		Addr:    ":8080",
	}, nil
}

func (c Config) DBPath() string {
	return filepath.Join(c.DataDir, "main.db")
}
