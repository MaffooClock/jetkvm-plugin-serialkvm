package main

import (
	"encoding/json"
	"fmt"
	"go.bug.st/serial"
	"log"
	"os"
)

type KvmInput struct {
	ControlMessage string `json:"command"`
}

type Config struct {
	Device     string      `json:"device"`
	SerialMode serial.Mode `json:"serial_mode"`
	Inputs     []KvmInput  `json:"inputs"`
}

const configPath = "/userdata/jetkvm/plugins/serialkvm/config.json"

var defaultConfig = &Config{
	Device: "/dev/ttyS3",
	SerialMode: serial.Mode{
		BaudRate: 115200,
		DataBits: 8,
		Parity:   serial.NoParity,
		StopBits: serial.OneStopBit,
	},
	Inputs: []KvmInput{},
}

var config *Config

func LoadConfig() {
	if config != nil {
		return
	}

	file, err := os.Open(configPath)
	if err != nil {
		log.Println("default config file doesn't exist, using default")
		config = defaultConfig
		return
	}
	defer file.Close()

	var loadedConfig Config
	if err := json.NewDecoder(file).Decode(&loadedConfig); err != nil {
		log.Printf("config file JSON parsing failed, %v\n", err)
		config = defaultConfig
		return
	}

	config = &loadedConfig
}

func SaveConfig() error {
	file, err := os.Create(configPath)
	if err != nil {
		return fmt.Errorf("failed to create config file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(config); err != nil {
		return fmt.Errorf("failed to encode config: %w", err)
	}

	return nil
}
