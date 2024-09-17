package communication

import (
	"fmt"
	"os"

	"github.com/bwmarrin/discordgo"
)

var discordBotToken string
var discordBotPrefix string
var discordBotSession *discordgo.Session
var discordBotId string

// ReadConfig reads the config.json file and unmarshals it into the Config struct
func discordBotReadConfig() error {
	discordBotToken = os.Getenv("COMMUNICATION_DISCORD_BOT_TOKEN")
	discordBotPrefix = "!"
	return nil
}

func discordBotMessageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == discordBotId {
		return
	}
	if message.Content == "hello job bot" {
		session.ChannelMessageSend(message.ChannelID, "Let's get you hired!")
	}
}

func InitDiscordBot() {
	discordBotReadConfig()
	discordBotSession, err := discordgo.New("Bot " + discordBotToken)
	if err != nil {
		fmt.Printf("error - Starting Communication:Discord Bot - %v", err)
		return
	}

	discordBotUser, err := discordBotSession.User("@me")
	if err != nil {
		fmt.Printf("error - Getting Session User Communication:Discord Bot - %v", err)
		return
	}
	discordBotId = discordBotUser.ID
	discordBotSession.AddHandler(discordBotMessageHandler)

	// discordBotSession.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	err = discordBotSession.Open()
	if err != nil {
		fmt.Printf("error - Creating Session Communication:Discord Bot - %v", err)
		return
	}

	<-make(chan struct{})

	fmt.Println("Communication:Discord Bot - Online")
}
