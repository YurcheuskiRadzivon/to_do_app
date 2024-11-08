package config

import (
	"os"
	"testing"
)

const testYAML = `
database:
  user: "testuser"
  password: "testpassword"
  host: "testhost"
  port: "1234"
  name: "testdb"
`

func createTestConfigFile() (string, error) {
	f, err := os.CreateTemp("", "test_config*.yaml")
	if err != nil {
		return "", err
	}
	defer f.Close()

	if _, err := f.WriteString(testYAML); err != nil {
		return "", err
	}

	return f.Name(), nil
}

func TestLoadConfig(t *testing.T) {
	fileName, err := createTestConfigFile()
	if err != nil {
		t.Fatalf("Could not create test config file: %v", err)
	}
	defer os.Remove(fileName)

	cfg, err := loadConfig(fileName)
	if err != nil {
		t.Fatalf("loadConfig() error: %v", err)
	}

	if cfg.User != "testuser" {
		t.Errorf("Expected User to be 'testuser', got %s", cfg.User)
	}
	if cfg.Password != "testpassword" {
		t.Errorf("Expected Password to be 'testpassword', got %s", cfg.Password)
	}
	if cfg.Host != "testhost" {
		t.Errorf("Expected Host to be 'testhost', got %s", cfg.Host)
	}
	if cfg.Port != "1234" {
		t.Errorf("Expected Port to be '1234', got %s", cfg.Port)
	}
	if cfg.Name != "testdb" {
		t.Errorf("Expected Name to be 'testdb', got %s", cfg.Name)
	}
}

func TestConnectionString(t *testing.T) {
	cfg := &DBConfig{
		User:     "testuser",
		Password: "testpassword",
		Host:     "testhost",
		Port:     "1234",
		Name:     "testdb",
	}

	expected := "postgres://testuser:testpassword@testhost:1234/testdb"
	got := cfg.ConnectionString()

	if got != expected {
		t.Errorf("Expected %s, got %s", expected, got)
	}
}

func TestGetConfig(t *testing.T) {
	fileName, err := createTestConfigFile()
	if err != nil {
		t.Fatalf("Could not create test config file: %v", err)
	}
	defer os.Remove(fileName)

	connStr, err := GetConfig(fileName)
	if err != nil {
		t.Fatalf("GetConfig() error: %v", err)
	}

	expected := "postgres://testuser:testpassword@testhost:1234/testdb"
	if connStr != expected {
		t.Errorf("Expected %s, got %s", expected, connStr)
	}
}
