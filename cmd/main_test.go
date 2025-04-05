package main

import (
	"os"
	"testing"
)

func TestNewConfig(t *testing.T) {

	yamlContent := `
mailscheduler:
  Port: "9090"
`

	tmpFile, err := os.CreateTemp("", "config-*.yaml")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	_, err = tmpFile.Write([]byte(yamlContent))
	if err != nil {
		t.Fatal(err)
	}
	tmpFile.Close()

	c, err := NewConfig(tmpFile.Name())
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	if c.Mailscheduler.Port != ":9090" {
		t.Errorf("Expected port to be ':9090', got: %v", c.Mailscheduler.Port)
	}
}

func TestNewServer(t *testing.T) {
	cfg := &Config{}
	cfg.Mailscheduler.Port = ":8080"

	s := NewServer(cfg)
	if s.router == nil {
		t.Error("Expected gin engine to be initialized")
	}
	if s.config == nil {
		t.Error("Expected config to be set")
	}
	if s.config.Mailscheduler.Port != ":8080" {
		t.Errorf("Expected port ':8080', got %s", s.config.Mailscheduler.Port)
	}
}
