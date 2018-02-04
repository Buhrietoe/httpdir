package main

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	var testListenString = ":8080"
	var testServeDir = "/test/path"

	err := os.Setenv("HTTP_DIR", testServeDir)
	if err != nil {
		t.Error(err)
	}

	config, err := LoadConfig()
	if err != nil {
		t.Error(err)
	}

	if config.ListenString != testListenString {
		t.Errorf("expected :8080, got %v", config.ListenString)
	}

	if config.ServeDir != testServeDir {
		t.Errorf("expected %s, got %v", testServeDir, config.ServeDir)
	}
}
