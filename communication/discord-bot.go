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

// ComesFromDM returns true if a message comes from a DM channel
func comesFromDM(session *discordgo.Session, message *discordgo.MessageCreate) (bool, error) {
	channel, err := session.State.Channel(message.ChannelID)
	if err != nil {
		if channel, err = session.Channel(message.ChannelID); err != nil {
			return false, err
		}
	}
	return channel.Type == discordgo.ChannelTypeDM, nil
}

// MessageHandler handles messages sent to the Discord bot
func discordBotMessageHandler(session *discordgo.Session, message *discordgo.MessageCreate) {
	if message.Author.ID == discordBotId {
		return
	}

	isDirectMessage, _ := comesFromDM(session, message)
	if isDirectMessage == true {
		// Send a reply to the direct message
		discordBotSlashCommands(session, message)
	}
}

// InitDiscordBot initializes the Discord bot
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

	err = discordBotSession.Open()
	if err != nil {
		fmt.Printf("error - Creating Session Communication:Discord Bot - %v", err)
		return
	}

	fmt.Println("Communication:Discord Bot - Online")
}
