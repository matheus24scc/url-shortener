package config

import (
	"os"
	"testing"
)

func TestLoadDefaults(t *testing.T) {
	// Ensure a clean env so defaults apply.
	for _, k := range []string{"SERVER_PORT", "DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME", "REDIS_ADDR", "REDIS_PASSWORD", "REDIS_DB"} {
		os.Unsetenv(k)
	}

	cfg := Load()
	if cfg == nil {
		t.Fatal("Load() returned nil")
	}
	if cfg.ServerPort != "8080" {
		t.Fatalf("expected default ServerPort %q, got %q", "8080", cfg.ServerPort)
	}
	if cfg.DBPort != "5432" {
		t.Fatalf("expected default DBPort %q, got %q", "5432", cfg.DBPort)
	}
	if cfg.RedisDB != 0 {
		t.Fatalf("expected default RedisDB 0, got %d", cfg.RedisDB)
	}
}

func TestLoadOverrides(t *testing.T) {
	os.Setenv("SERVER_PORT", "9090")
	os.Setenv("REDIS_DB", "3")
	defer func() {
		os.Unsetenv("SERVER_PORT")
		os.Unsetenv("REDIS_DB")
	}()

	cfg := Load()
	if cfg.ServerPort != "9090" {
		t.Fatalf("expected SERVER_PORT override %q, got %q", "9090", cfg.ServerPort)
	}
	if cfg.RedisDB != 3 {
		t.Fatalf("expected REDIS_DB override 3, got %d", cfg.RedisDB)
	}
}
