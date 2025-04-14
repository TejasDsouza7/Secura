package storage

import (
    "encoding/json"
    "log"
    "os"
)

type Config struct {
    StoragePath string `json:"storage_path"`
    ServerPort  string `json:"server_port"`
}

func LoadConfig() Config {
    file, err := os.Open("config.json")
    if err != nil {
        log.Fatal("Failed to open config.json")
    }
    defer file.Close()

    decoder := json.NewDecoder(file)
    config := Config{}
    err = decoder.Decode(&config)
    if err != nil {
        log.Fatal("Failed to parse config.json")
    }

    return config
}

func EnsureStoragePath(path string) {
    if _, err := os.Stat(path); os.IsNotExist(err) {
        err := os.MkdirAll(path, 0755)
        if err != nil {
            log.Fatal("Failed to create storage directory")
        }
    }
}