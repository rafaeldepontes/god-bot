package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rafaeldepontes/god-bot/internal/bot"
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		panic(fmt.Errorf("[ERROR] couldn't initialize the project, problem with .env: %w", err))
	}

	bot.BotToken = os.Getenv("DISCORD_KEY")
}

func main() {
	bot.Run()
}
