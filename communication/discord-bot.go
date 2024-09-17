package communication

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

func InitDiscordBot() {
	discordBotToken := os.Getenv("COMMUNICATION_DISCORD_BOT_TOKEN")
	discordBotSession, err := discordgo.New(discordBotToken)
	if err != nil {
		fmt.Printf("error - Starting Communication:Discord Bot - %v", err)
	}

	discordBotSession.AddHandler(func(session *discordgo.Session, message *discordgo.MessageCreate) {
		if message.Author.ID == session.State.User.ID {
			return
		}

		if message.Content == "hello job bot" {
			session.ChannelMessageSend(message.ChannelID, "Let's get you hired!")
		}
	})

	discordBotSession.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	err = discordBotSession.Open()
	if err != nil {
		fmt.Printf("error - Creating Session Communication:Discord Bot - %v", err)
	}

	// defer discordBotSession.Close()
	fmt.Println("Communication:Discord Bot - Online")

	// Listen for ctrl+c
	// sc := make(chan os.Signal, 1)
	// signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	// <-sc
}
