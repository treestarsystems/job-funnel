package communication

import (
	"fmt"
	"job-funnel/retrieve"
	"job-funnel/utils"
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
	if message.Content == "!jobs" {
		resultJobPosts := retrieve.RetrieveDataFromSqliteAll()
		if len(resultJobPosts) == 0 {
			session.ChannelMessageSend(message.ChannelID, "Sorry no job posts available.")
		}

		jobPostsResponse := utils.JobPostsToString(resultJobPosts)
		// session.ChannelMessageSend(message.ChannelID, "Sending jobs")
		session.ChannelMessageSend(message.ChannelID, jobPostsResponse)
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

	fmt.Println("Communication:Discord Bot - Online")
}
