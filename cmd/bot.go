package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/TiunovNN/go-tg-wol/pkg/bot"
	"github.com/TiunovNN/go-tg-wol/pkg/users"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	Token   string        `json:"token"`
	Users   []*users.User `json:"users"`
	LogFile string        `json:"log_file"`
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
	log.SetOutput(&lumberjack.Logger{
		Filename:   config.LogFile,
		MaxSize:    50, // megabytes
		MaxBackups: 3,
		MaxAge:     7,    //days
		Compress:   true, // disabled by default
	})
	bot.Start(config.Token, config.Users)
}
