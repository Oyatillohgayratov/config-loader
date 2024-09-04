package configloader

import (
    "os"
    "testing"
)

type Config struct {
    Host string `json:"host" yaml:"host"`
    Port int    `json:"port" yaml:"port"`
}

func TestLoadJSONConfig(t *testing.T) {
    var config Config
    err := LoadJSONConfig("testdata/config.json", &config)
    if err != nil {
        t.Fatalf("Failed to load JSON config: %v", err)
    }

    if config.Host != "localhost" || config.Port != 8080 {
        t.Errorf("Unexpected config: %+v", config)
    }
}

func TestLoadYAMLConfig(t *testing.T) {
    var config Config
    err := LoadYAMLConfig("testdata/config.yaml", &config)
    if err != nil {
        t.Fatalf("Failed to load YAML config: %v", err)
    }

    if config.Host != "localhost" || config.Port != 8080 {
        t.Errorf("Unexpected config: %+v", config)
    }
}

func TestLoadENVConfig(t *testing.T) {
    err := LoadENVConfig("testdata/.env")
    if err != nil {
        t.Fatalf("Failed to load ENV config: %v", err)
    }

    host := os.Getenv("HOST")
    port := os.Getenv("PORT")

    if host != "localhost" || port != "8080" {
        t.Errorf("Unexpected ENV config: HOST=%s, PORT=%s", host, port)
    }
}
