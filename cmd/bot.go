package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/TiunovNN/go-tg-wol/pkg/bot"
	"github.com/TiunovNN/go-tg-wol/pkg/users"
)

type Config struct {
	Token string
	Users []*users.User
}

func readConfig() (*Config, error) {
	configFile, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer configFile.Close()
	decoder := json.NewDecoder(configFile)
	config := Config{}
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}
	return &config, err
}

func main() {
	config, err := readConfig()
	if err != nil {
		log.Fatalf("Could not read config %s", err)
	}
	bot.Start(config.Token, config.Users)
}
