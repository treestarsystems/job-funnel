package communication

import (
	"fmt"
	"job-funnel/retrieve"
	"job-funnel/utils"
	"os"
	"strings"

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
	if message.Content == "!help" {
		helpMenu := []string{
			"Commands:",
			"!job:all - Get all job posts (see notes below)",
			"!job:random - Get a random single job post",
			"!job:search <search term as single word no spaces> - Get all job posts that match the search string (see notes below)",
			"!applied:<job ID> - Mark a job post as applied so it is not shown in future searches (not available yet)",
			"!applied:all - Get all job posts that have been marked as applied (not available yet)",
			"!help - Get a list of commands",
			"\nNote: \n- Due to [message length limitations of 2k characters](https://support.discord.com/hc/en-us/articles/360034632292-Sending-Messages#h_01FSWZRDKD7310TJHGCA616642:~:text=place%20multiple%20emojis%3A-,Character%20Limit,-The%20character%20cap), I can not return all stored job posts. \n- If you multiple job posts will be returned, the order will be random so you can get different jobs. \n- This bot is still in development.",
		}
		session.ChannelMessageSend(message.ChannelID, strings.Join(helpMenu, "\n"))
	}
	if message.Content == "!job:all" {
		resultJobPosts := retrieve.RetrieveDataFromSqliteAll()
		if len(resultJobPosts) == 0 {
			session.ChannelMessageSend(message.ChannelID, "Sorry, no job posts available.")
		}
		jobPostsResponse := utils.JobPostsToString(resultJobPosts)
		session.ChannelMessageSend(message.ChannelID, jobPostsResponse)
		// session.ChannelMessageSend(message.ChannelID, "Sorry no job posts available.")
	}
	if message.Content == "!job:random" {
		resultJobPosts := retrieve.RetrieveDataFromSqliteAll()
		if len(resultJobPosts) == 0 {
			session.ChannelMessageSend(message.ChannelID, "Sorry, no job posts available.")
		}
		jobPostsResponse := utils.JobPostsToStringSingle(resultJobPosts)
		session.ChannelMessageSend(message.ChannelID, jobPostsResponse)
	}
	if strings.HasPrefix(message.Content, "!job:search ") {
		searchTerm := strings.TrimPrefix(message.Content, "!job:search ")
		// Need new method to retrieve data from sqlite using search term.
		resultJobPosts := retrieve.RetrieveDataFromSqliteAll()
		var filteredJobPosts []utils.JobPost
		for _, job := range resultJobPosts {
			if strings.Contains(strings.ToLower(job.JobTitle), strings.ToLower(searchTerm)) ||
				strings.Contains(strings.ToLower(job.Description), strings.ToLower(searchTerm)) ||
				strings.Contains(strings.ToLower(strings.Join(job.CodingLanguage, " ")), strings.ToLower(searchTerm)) ||
				strings.Contains(strings.ToLower(strings.Join(job.CodingFramework, " ")), strings.ToLower(searchTerm)) ||
				strings.Contains(strings.ToLower(strings.Join(job.Database, " ")), strings.ToLower(searchTerm)) ||
				strings.Contains(strings.ToLower(strings.Join(job.WorkLocation, " ")), strings.ToLower(searchTerm)) {
				filteredJobPosts = append(filteredJobPosts, job)
			}
		}
		if len(filteredJobPosts) == 0 {
			session.ChannelMessageSend(message.ChannelID, "Sorry, no job posts found matching the search term.")
			return
		}
		jobPostsResponse := utils.JobPostsToString(filteredJobPosts)
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
