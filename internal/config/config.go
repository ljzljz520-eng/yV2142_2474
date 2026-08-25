package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	DBPath, Session string
	PageSize        int
}

func Load() Config {
	p := os.Getenv("IDIOM_DB")
	if p == "" {
		p = "idiomchain.db"
	}
	s := os.Getenv("IDIOM_SESSION")
	if s == "" {
		s = "default"
	}
	return Config{DBPath: p, Session: s, PageSize: 20}
}
func ForPath(path, session string) Config {
	if path == "" {
		path = "idiomchain.db"
	}
	if session == "" {
		session = "default"
	}
	return Config{DBPath: filepath.Clean(path), Session: session, PageSize: 20}
}
