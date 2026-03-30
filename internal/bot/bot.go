package bot

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
)

var BotToken string

func Run() {
	s, err := discordgo.New("Bot " + BotToken)
	if err != nil {
		panic(fmt.Errorf("[ERROR] couldn't create a discord session: %w", err))
	}

	s.AddHandler(deleteMessage)

	s.Identify.Intents = discordgo.IntentsGuildMessages | discordgo.IntentsMessageContent
	if err := s.Open(); err != nil {
		panic(fmt.Errorf("[ERROR] couldn't create the websocket connection: %w", err))
	}
	defer s.Close()

	log.Println("Bot running...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func deleteMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	switch {
	case m.MentionEveryone:
		if err := s.ChannelMessageDelete(m.ChannelID, m.ID); err != nil {
			log.Printf("Everyone received -> %s\n", m.Content)
			log.Printf("[WARN] couldn't delete the message: %s\n", err.Error())
		}
	}
}
