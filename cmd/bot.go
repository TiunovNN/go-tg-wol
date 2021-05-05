package main

import (
	"os"

	"github.com/TiunovNN/go-tg-wol/pkg/bot"
)

func main() {
	token := os.Getenv("TOKEN")
	bot.Start(token)
}
