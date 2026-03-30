package bot

import (
	"fmt"
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

	fmt.Println("Bot running...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}

func deleteMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	fmt.Printf("Message received: %s\n", m.Content)

	switch {
	case m.MentionEveryone:
		if err := s.ChannelMessageDelete(m.ChannelID, m.ID); err != nil {
			fmt.Printf("[WARN] couldn't delete the message: %s\n", err.Error())
		}
	}
}
