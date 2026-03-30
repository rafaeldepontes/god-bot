package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rafaeldepontes/god-bot/internal/bot"
)

func init() {
	_ = godotenv.Load(".env")
	bot.BotToken = os.Getenv("DISCORD_KEY")
}

func main() {
	bot.Run()
}
