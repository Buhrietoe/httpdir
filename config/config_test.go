package config

import (
	"os"
	"testing"
)

func TestLoad(t *testing.T) {
	var testListenString string = ":8080"
	var testServeDir string = "/test/path"

	err := os.Setenv("HTTP_DIR", testServeDir)
	if err != nil {
		t.Error(err)
	}

	config := Load()

	if config.ListenString != testListenString {
		t.Errorf("expected :8080, got %v", config.ListenString)
	}

	if config.ServeDir != testServeDir {
		t.Errorf("expected %s, got %v", testServeDir, config.ServeDir)
	}
}
