package communication

import (
	"fmt"
	"job-funnel/retrieve"
	"job-funnel/utils"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func discordBotSlashCommandHelpMenu() string {
	helpMenu := []string{
		"Commands:",
		"!job:some - Get some job posts (see notes below)",
		"!job:random - Get a single random job post",
		"!job:search <search term as single word no spaces> - Get some job posts that match the search string (see notes below)",
		"!applied:<job ID> - Mark a job post as applied so it is not shown in future searches (not available yet)",
		"!applied:all - Get all job posts that have been marked as applied (not available yet)",
		"!help - Get a list of commands",
		"\nNote: \n- Due to [message length limitations of 2k characters](https://support.discord.com/hc/en-us/articles/360034632292-Sending-Messages#h_01FSWZRDKD7310TJHGCA616642:~:text=place%20multiple%20emojis%3A-,Character%20Limit,-The%20character%20cap), I can not return all stored job posts. \n- If multiple job posts will be returned, the order will be random so you can get different jobs. \n- This bot is still in development.",
	}
	return strings.Join(helpMenu, "\n")
}

// sendAllJobPosts retrieves all job posts and sends them as a response.
func sendAllJobPosts(session *discordgo.Session, channelID string) {
	resultJobPosts := retrieve.RetrieveDataFromSqliteAll()
	if len(resultJobPosts) == 0 {
		session.ChannelMessageSend(channelID, "Sorry, no job posts available.")
		return
	}
	jobPostsResponse := utils.JobPostsToString(resultJobPosts)
	session.ChannelMessageSend(channelID, jobPostsResponse)
}

// sendRandomJobPost retrieves all job posts, selects a random one, and sends it as a response.
func sendRandomJobPost(session *discordgo.Session, channelID string) {
	resultJobPosts := retrieve.RetrieveDataFromSqliteAll()
	if len(resultJobPosts) == 0 {
		session.ChannelMessageSend(channelID, "Sorry, no job posts available.")
		return
	}
	jobPostsResponse := utils.JobPostsToStringSingle(resultJobPosts)
	session.ChannelMessageSend(channelID, jobPostsResponse)
}

// searchAndSendJobPosts searches for job posts based on the search term and sends the response.
func searchAndSendJobPosts(session *discordgo.Session, channelID string, searchTerm string) {
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
		session.ChannelMessageSend(channelID, "Sorry, no job posts found matching the search term.")
		return
	}
	jobPostsResponse := utils.JobPostsToString(filteredJobPosts)
	session.ChannelMessageSend(channelID, jobPostsResponse)
}

func discordBotSlashCommands(session *discordgo.Session, message *discordgo.MessageCreate) {
	appliedRegex := regexp.MustCompile(`^!applied:[a-zA-Z0-9]{20}$`)
	jobSearchRegex := regexp.MustCompile(`^!job(s)?:(all|random|search )`)

	switch {
	case message.Content == "!help":
		session.ChannelMessageSend(message.ChannelID, discordBotSlashCommandHelpMenu())
	case jobSearchRegex.MatchString(message.Content):
		if strings.Contains(message.Content, ":all") {
			sendAllJobPosts(session, message.ChannelID)
		}
		if strings.Contains(message.Content, ":random") {
			sendRandomJobPost(session, message.ChannelID)
		}
		if strings.Contains(message.Content, ":search ") {
			searchTerm := jobSearchRegex.ReplaceAllString(message.Content, "")
			searchAndSendJobPosts(session, message.ChannelID, searchTerm)
		}
	case message.Content == "!applied:all":
		session.ChannelMessageSend(message.ChannelID, "Feature not available yet.")
	case appliedRegex.MatchString(message.Content):
		jobId := strings.Split(message.Content, ":")[1]
		session.ChannelMessageSend(message.ChannelID, fmt.Sprintf("Feature not available yet. Job ID: %s", jobId))
	}
}
