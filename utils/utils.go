package utils

import (
	"regexp"
)

// Extracts all salary information from the given string.
func ExtractSalaries(text string) []string {
	// Define a regex pattern to match salary ranges and amounts
	re := regexp.MustCompile(`\$\d{1,3}(,\d{3})*(\.\d{2})?(\s*-\s*\$\d{1,3}(,\d{3})*(\.\d{2})?)?`)

	// Find all matches in the text
	matches := re.FindAllString(text, -1)
	if len(matches) == 0 {
		return []string{"salary information not found"}
	}
	return matches
}

// Extracts city or state information from the given string.
// This can be done better: https://devcodef1.com/news/1018489/regex-code-for-extracting-city-state-and-zip-from-address-string
func ExtractCityOrState(text string) []string {
	// Define a regex pattern to match common city or state patterns
	re := regexp.MustCompile(`\b[A-Z][a-z]+(?:\s+[A-Z][a-z]+)*(?:,\s*[A-Z]{2})?\b`)

	// Find all matches in the text
	matches := re.FindAllString(text, -1)
	if len(matches) == 0 {
		return []string{"city or state information not found"}
	}
	return matches
}

// Extracts programming languages from the given string.
func ExtractProgrammingLanguages(text string) []string {
	// Define a regex pattern to match common programming languages
	re := regexp.MustCompile(`(?i)\b(java|python|javascript|c\+\+|c#|ruby|go|swift|kotlin|php|typescript|rust|scala|perl|haskell|r|objective-c|dart|lua|matlab|groovy|shell|powershell|visual basic|assembly|cobol|fortran|pascal|ada|lisp|scheme|prolog|erlang|elixir|f#|ocaml|clojure|julia|vhdl|verilog|solidity|sql|pl/sql|t-sql|sas|spss|stata|racket|smalltalk|abap|actionscript|apex|awk|bash|batch|bc|brainfuck|caml|chapel|clean|clipper|cmake|cobol|coffeescript|crystal|curl|d|dcl|dylan|eiffel|emacs lisp|euphoria|forth|gams|gap|gdl|gdscript|gml|gnuplot|idl|io|j|jscript|labview|ladder logic|livecode|logo|m4|max/msp|mercury|ml|modula-2|mumps|natural|nim|nxc|opencl|openedge abl|openscad|p4|pike|pl/i|postscript|pure data|q|racket|raku|rexx|ring|s-lang|sml|spark|spin|tcl|turing|vala|vbscript|vim script|wolfram|x10|xbase|xojo|zig)\b`)

	// Find all matches in the text
	matches := re.FindAllString(text, -1)
	if len(matches) == 0 {
		return []string{"programming languages not found"}
	}
	return matches
}

func ExtractCommonFrameworks(text string) []string {
	// Define a regex pattern to match common coding frameworks
	re := regexp.MustCompile(`(?i)\b(react|angular|vue\.js|django|flask|spring|express|ruby.*on.*rails|ruby|rails|laravel|asp\.net|\.net|dotnet|symfony|svelte|ember\.js|backbone\.js|meteor|next\.js|nuxt\.js|gatsby|bootstrap|foundation|tailwind css|jquery|redux|nestjs|koa|fastapi|phoenix|play|struts|blade|gin|beego|echo|fiber|rocket|actix|tornado|bottle|pyramid|cherrypy|hug|falcon|sanic|fastify|hapi|loopback|feathers|adonisjs|sails|aurelia|alpine\.js|stimulus|litelement|stencil|node.*js|nest.*js)\b`)

	// Find all matches in the text
	matches := re.FindAllString(text, -1)
	if len(matches) == 0 {
		return []string{"coding frameworks not found"}
	}
	return matches
}

// Extracts database types from the given string in a case-insensitive manner.
func ExtractDatabaseTypes(text string) []string {
	// Define a regex pattern to match common database types (case-insensitive)
	re := regexp.MustCompile(`(?i)\b(no.*sql|ms.*sql|mysql|postgre.*sql|sqlite|mongodb|oracle|sql.*server|maria.*db|redis|cassandra|elasticsearch|firebase|dynamo.*db|couch.*db|neo4j|hbase|memcached|couchbase|db2|teradata|snowflake|bigquery|amazon.*aurora|amazon.*redshift|google.*cloud.*spanner|microsoft.*access|informix|ingres|interbase|firebird|sybase|volt.*db|greenplum)\b`)

	// Find all matches in the text
	matches := re.FindAllString(text, -1)
	if len(matches) == 0 {
		return []string{"database types not found"}
	}
	return matches
}

// Extracts database types from the given string in a case-insensitive manner.
func ExtractAWSServiceNames(text string) []string {
	// Define a regex pattern to match common database types (case-insensitive)
	re := regexp.MustCompile(`(?i)\b(EC2|S3|RDS|Lambda|DynamoDB|ECS|EKS|CloudFront|Route 53|VPC|IAM|CloudWatch|SNS|SQS|Redshift|Glue|Athena|Kinesis|Elastic Beanstalk|CloudFormation|Elastic Load Balancing|SageMaker|Fargate|Aurora|Elasticache|Kinesis|Step Functions|AppSync|CodePipeline|CodeBuild|CodeDeploy|CodeCommit|Amplify|Lightsail|Batch|OpsWorks|CloudTrail|Cloud9|CloudHSM|Direct Connect|Elastic Transcoder|GuardDuty|Inspector|Macie|Organizations|Secrets Manager|Security Hub|Shield|WAF|WorkSpaces)\b`)

	// Find all matches in the text
	matches := re.FindAllString(text, -1)
	if len(matches) == 0 {
		return []string{"database types not found"}
	}
	return matches
}

func RemoveHTMLTags(text string) string {
	// Define a regex pattern to match HTML tags
	re := regexp.MustCompile(`<[^>]*>`)
	// Replace all HTML tags with an empty string
	cleanText := re.ReplaceAllString(text, "")
	return cleanText
}

// Extracts web links that are not part of image tags from the given string and deduplicates them.
func ExtractNonImageLinks(text string) []string {
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

	// Filter out image URLs and deduplicate
	linkMap := make(map[string]bool)
	var nonImageLinks []string
	for _, link := range urlMatches {
		if !imageURLs[link] {
			if !linkMap[link] {
				linkMap[link] = true
				nonImageLinks = append(nonImageLinks, link)
			}
		}
	}

	if len(nonImageLinks) == 0 {
		return []string{"no links found"}
	}
	return nonImageLinks
}
