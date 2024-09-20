package utils

import (
	"fmt"
	"math/rand"
	"reflect"
	"regexp"
	"strings"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Parses all salary information from the given string.
func ParseSalaries(text string) []string {
	// Define a regex pattern to match salary ranges and amounts
	re := regexp.MustCompile(`\$\d{1,3}(,\d{3})*(\.\d{2})?(\s*-\s*\$\d{1,3}(,\d{3})*(\.\d{2})?)?`)

	// Find all matches in the text
	matches := re.FindAllString(text, -1)
	if len(matches) == 0 {
		return []string{}
	}
	return matches
}

// Parses city or state information from the given string.
// This can be done better: https://devcodef1.com/news/1018489/regex-code-for-extracting-city-state-and-zip-from-address-string
func ParseCityOrState(text string) []string {
	// Define a regex pattern to match common city or state patterns
	re := regexp.MustCompile(`\b[A-Z][a-z]+(?:\s+[A-Z][a-z]+)*(?:,\s*[A-Z]{2})?\b`)

	// Find all matches in the text
	matches := re.FindAllString(text, -1)
	if len(matches) == 0 {
		return []string{}
	}
	return matches
}

func ParseJobWorkLocation(text string) []string {
	// Define a regex pattern to match common city or state patterns
	re := regexp.MustCompile(`(?i)\b(remote|hybrid|on.*site)\b`)

	// Find all matches in the text
	matches := re.FindAllString(text, -1)
	if len(matches) == 0 {
		return []string{}
	}
	return DeduplicateSliceContents(matches).([]string)
}

// Parses programming languages from the given string.
func ParseCommonProgrammingLanguages(text string) []string {
	// Define a regex pattern to match common programming languages
	re := regexp.MustCompile(`(?i)\b(node\s?js|java|python|javascript|c\+\+|c#|ruby|go|swift|kotlin|php|typescript|rust|scala|perl|haskell|r|objective-c|dart|lua|matlab|groovy|shell|powershell|visual\s?basic|assembly|cobol|fortran|pascal|ada|lisp|scheme|prolog|erlang|elixir|f#|ocaml|clojure|julia|vhdl|verilog|solidity|sql|pl\/sql|t-sql|sas|spss|stata|racket|smalltalk|abap|actionscript|apex|awk|bash|batch|bc|brainfuck|caml|chapel|clean|clipper|cmake|cobol|coffeescript|crystal|curl|dcl|dylan|eiffel|emacs\s?lisp|euphoria|forth|gams|gap|gdl|gdscript|gml|gnuplot|idl|jscript|labview|ladder\s?logic|livecode|logo|m4|max\/msp|mercury|ml|modula((?:\s|-)2)?|mumps|natural|nim|nxc|opencl|openedge\s?abl|openscad|p4|pike|pl/i|postscript|pure\s?data|racket|raku|rexx|ring|s-lang|sml|spark|spin|tcl|turing|vala|vbscript|vim\s?script|wolfram|x10|xbase|xojo|zig)\b`)

	// Find all matches in the text
	matches := re.FindAllString(text, -1)
	if len(matches) == 0 {
		return []string{}
	}
	return DeduplicateSliceContents(matches).([]string)
}

// Parses database types from the given string in a case-insensitive manner.
func ParseCommonFrameworks(text string) []string {
	// Define a regex pattern to match common coding frameworks
	re := regexp.MustCompile(`(?i)\b(react|angular|vue\s?js|django|flask|spring|express|ruby\son\srails|rails|laravel|asp\s?net|\.?net|dotnet|symfony|svelte|ember\s?js|backbone\s?js|meteor|next\s?js|nuxt\s?js|gatsby|bootstrap|foundation|tailwind\s?css|jquery|redux|nestjs|koa|fastapi|phoenix|play|struts|blade|gin|beego|echo|fiber|rocket|actix|tornado|bottle|pyramid|cherrypy|hug|falcon|sanic|fastify|hapi|loopback|feathers|adonisjs|sails|aurelia|alpine\s?js|stimulus|litelement|stencil|nest\s?js)\b`)

	// Find all matches in the text
	matches := re.FindAllString(text, -1)
	if len(matches) == 0 {
		return []string{}
	}
	return DeduplicateSliceContents(matches).([]string)
}

// // Parses database types from the given string in a case-insensitive manner.
func ParseDatabaseTypes(text string) []string {
	// Define a regex pattern to match common database types (case-insensitive)
	re := regexp.MustCompile(`(?i)\b(?:no\s?sql|ms\s?sql|mysql|postgre(?:sql)?|sqlite|mongodb|mongo\s?db|oracle|maria(?:db)?|redis|cassandra|elasticsearch|firebase|dynamo(?:db)?|couch(?:db)?|neo4j|hbase|memcached|couchbase|db2|teradata|snowflake|bigquery|big\s?query|aurora|redshift|cloud\s?spanner|microsoft\s?access|informix|ingres|interbase|firebird|sybase|volt\s?db|voltdb|greenplum)\b`)

	// Find all matches in the text
	matches := re.FindAllString(text, -1)
	if len(matches) == 0 {
		return []string{}
	}
	return DeduplicateSliceContents(matches).([]string)
}

// Parses database types from the given string in a case-insensitive manner.
func ParseAWSServiceNames(text string) []string {
	// Define a regex pattern to match common database types (case-insensitive)
	re := regexp.MustCompile(`(?i)\b(?:ec2|s3|rds|lambda|dynamo\s?db|ecs|eks|cloudfront|route\s?53|vpc|iam|cloudwatch|sns|sqs|redshift|glue|athena|kinesis|elastic(?:.*beanstalk|.*load.*balancing|ache)?|cloud\s?formation|sagemaker|fargate|aurora|step\s?functions(s?)|appsync|code(?:pipeline|build|deploy|commit)|amplify|lightsail|batch|opsworks|cloudtrail|cloud\s?9|cloudhsm|direct\s?connect|elastic\s?transcoder|guardduty|inspector|macie|organizations|secrets\s?manager|security\s?hub|shield|waf|workspaces)\b`)

	// Find all matches in the text
	matches := re.FindAllString(text, -1)
	if len(matches) == 0 {
		return []string{}
	}
	return DeduplicateSliceContents(matches).([]string)
}

// Removes all HTML tags from the given string.
func RemoveHTMLTags(text string) string {
	// Define a regex pattern to match HTML tags
	re := regexp.MustCompile(`<[^>]*>`)
	// Replace all HTML tags with an empty string
	cleanText := re.ReplaceAllString(text, "")
	return cleanText
}

// Parses web links that are not part of image tags from the given string and deduplicates them.
func ParseNonImageLinks(text string) []string {
	// Define a regex pattern to match URLs
	urlRe := regexp.MustCompile(`https?://[^\s]+`)
	// Define a regex pattern to match image tags
	imgTagRe := regexp.MustCompile(`<img[^>]+src="([^"]+)"[^>]*>`)

	// Find all URLs in the text
	urlMatches := urlRe.FindAllString(text, -1)
	// Find all image URLs in the text
	imgTagMatches := imgTagRe.FindAllStringSubmatch(text, -1)

	// Create a set of image URLs for easy lookup
	imageURLs := make(map[string]bool)
	for _, match := range imgTagMatches {
		if len(match) > 1 {
			imageURLs[match[1]] = true
		}
	}

	// Filter out image URLs
	var nonImageLinks []string
	for _, link := range urlMatches {
		if !imageURLs[link] {
			nonImageLinks = append(nonImageLinks, link)
		}
	}

	// Deduplicate the non-image links
	dedupedLinks := DeduplicateSliceContents(nonImageLinks).([]string)

	if len(dedupedLinks) == 0 {
		return []string{}
	}
	return dedupedLinks
}

// Deduplicates a slice of any type and converts all strings to lowercase.
func DeduplicateSliceContents(slice interface{}) interface{} {
	// Use reflection to get the value and type of the input slice
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return []string{}
	}

	// Create a map to track unique elements
	uniqueMap := make(map[interface{}]bool)
	uniqueSlice := reflect.MakeSlice(v.Type(), 0, v.Len())

	// Iterate over the input slice and add unique elements to the result slice
	for i := 0; i < v.Len(); i++ {
		elem := v.Index(i).Interface()

		// Convert strings to lowercase
		if str, ok := elem.(string); ok {
			elem = strings.ToLower(str)
		}

		if !uniqueMap[elem] {
			uniqueMap[elem] = true
			uniqueSlice = reflect.Append(uniqueSlice, reflect.ValueOf(elem))
		}
	}

	return uniqueSlice.Interface()
}

// JobPostsToString converts a slice of JobPost to a string representation of all job posts. Best for emails.
// func JobPostsToString(jobPosts []JobPost) string {
// 	var sb strings.Builder
// 	for _, job := range jobPosts {
// 		responseString := fmt.Sprintf("Title: %s\nSource: %s\nLocation: %s\nLanguages: %s\nFrameworks: %s\nDatabase: %s\nLinks:\n%s\n\n", job.JobTitle, job.JobSource, strings.Join(job.WorkLocation, ", "), strings.Join(job.CodingLanguage, ", "), strings.Join(job.CodingFramework, ", "), strings.Join(job.Database, ", "), strings.Join(job.Links, "\n"))
// 		sb.WriteString(responseString)
// 	}
// 	return sb.String()
// }

// JobPostsToString converts a slice of JobPost to a string representation of a random job post.
func JobPostsToStringSingle(jobPosts []JobPost) string {
	if len(jobPosts) == 0 {
		return "No job posts available."
	}

	// Pick a random job post
	randomIndex := rand.Intn(len(jobPosts))
	job := jobPosts[randomIndex]

	// Create the string representation of the job post
	var sb strings.Builder
	responseString := fmt.Sprintf(
		"Title: %s (ID: %s)\nSource: %s\nLocation: %s\nLanguages: %s\nFrameworks: %s\nDatabase: %s\nLinks:\n%s\n\n",
		job.JobTitle,
		job.JobId,
		job.JobSource,
		strings.Join(job.WorkLocation, ", "),
		strings.Join(job.CodingLanguage, ", "),
		strings.Join(job.CodingFramework, ", "),
		strings.Join(job.Database, ", "),
		strings.Join(job.Links, "\n"),
	)
	sb.WriteString(responseString)

	return sb.String()
}

// ShuffleJobPosts shuffles the slice of JobPost.
func ShuffleJobPosts(jobPosts []JobPost) {
	rand.Shuffle(len(jobPosts), func(i, j int) {
		jobPosts[i], jobPosts[j] = jobPosts[j], jobPosts[i]
	})
}

// JobPostsToString converts a slice of JobPost to a string representation, randomizes the order, and truncates the final string at 8000 characters.
func JobPostsToString(jobPosts []JobPost) string {
	ShuffleJobPosts(jobPosts)

	var sb strings.Builder
	for _, job := range jobPosts {
		responseString := fmt.Sprintf(
			"Title: %s (ID: %s)\nSource: %s\nLocation: %s\nLanguages: %s\nFrameworks: %s\nDatabase: %s\nLinks:\n%s\n\n",
			job.JobTitle,
			job.JobId,
			job.JobSource,
			strings.Join(job.WorkLocation, ", "),
			strings.Join(job.CodingLanguage, ", "),
			strings.Join(job.CodingFramework, ", "),
			strings.Join(job.Database, ", "),
			strings.Join(job.Links, "\n"),
		)
		if sb.Len()+len(responseString) > 2000 {
			break
		}
		sb.WriteString(responseString)
	}
	return sb.String()
}

// RandomString generates a random string of the specified length.
func RandomAplhaNumericString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
