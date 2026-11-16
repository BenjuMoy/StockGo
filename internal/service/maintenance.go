package service

import (
	"io"
	"os"
	"path/filepath"
	"time"
)

func Create(dbPath string) (string, error) {
	if err := os.MkdirAll("backups", 0755); err != nil {
		return "", err
	}

	name := time.Now().Format("2006-01-02-150405")

	dst := filepath.Join(
		"backups",
		"stocks-"+name+".db",
	)

	src, err := os.Open(dbPath)
	if err != nil {
		return "", err
	}
	defer src.Close()

	out, err := os.Create(dst)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, src)

	return dst, err
}
