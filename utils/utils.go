package utils

import (
	"reflect"
	"regexp"
	"strings"
)

// Parses all salary information from the given string.
func ParseSalaries(text string) []string {
	// Define a regex pattern to match salary ranges and amounts
	re := regexp.MustCompile(`\$\d{1,3}(,\d{3})*(\.\d{2})?(\s*-\s*\$\d{1,3}(,\d{3})*(\.\d{2})?)?`)

	// Find all matches in the text
	matches := re.FindAllString(text, -1)
	if len(matches) == 0 {
		return []string{"salary information not found"}
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
		return []string{"city or state information not found"}
	}
	return matches
}

func ParseJobWorkLocation(text string) []string {
	// Define a regex pattern to match common city or state patterns
	re := regexp.MustCompile(`(?i)\b(remote|hybrid|on.*site)\b`)

	// Find all matches in the text
	matches := re.FindAllString(text, -1)
	if len(matches) == 0 {
		return []string{"location information not found"}
	}
	return DeduplicateSliceContents(matches).([]string)
}

// Parses programming languages from the given string.
func ParseCommonProgrammingLanguages(text string) []string {
	// Define a regex pattern to match common programming languages
	re := regexp.MustCompile(`(?i)\b(node.*js|java|python|javascript|c\+\+|c#|ruby|go|swift|kotlin|php|typescript|rust|scala|perl|haskell|r|objective-c|dart|lua|matlab|groovy|shell|powershell|visual.*basic|assembly|cobol|fortran|pascal|ada|lisp|scheme|prolog|erlang|elixir|f#|ocaml|clojure|julia|vhdl|verilog|solidity|sql|pl/sql|t-sql|sas|spss|stata|racket|smalltalk|abap|actionscript|apex|awk|bash|batch|bc|brainfuck|caml|chapel|clean|clipper|cmake|cobol|coffeescript|crystal|curl|dcl|dylan|eiffel|emacs.*lisp|euphoria|forth|gams|gap|gdl|gdscript|gml|gnuplot|idl|jscript|labview|ladder.*logic|livecode|logo|m4|max/msp|mercury|ml|modula-2|mumps|natural|nim|nxc|opencl|openedge.*abl|openscad|p4|pike|pl/i|postscript|pure.*data|racket|raku|rexx|ring|s-lang|sml|spark|spin|tcl|turing|vala|vbscript|vim.*script|wolfram|x10|xbase|xojo|zig)\b`)

	// Find all matches in the text
	matches := re.FindAllString(text, -1)
	if len(matches) == 0 {
		return []string{"programming languages not found"}
	}
	return DeduplicateSliceContents(matches).([]string)
}

// Parses database types from the given string in a case-insensitive manner.
func ParseCommonFrameworks(text string) []string {
	// Define a regex pattern to match common coding frameworks
	re := regexp.MustCompile(`(?i)\b(react|angular|vue.*js|django|flask|spring|express|ruby.*on.*rails|rails|laravel|asp.*net|\.net|dotnet|symfony|svelte|ember.*js|backbone.*js|meteor|next.*js|nuxt.*js|gatsby|bootstrap|foundation|tailwind.*css|jquery|redux|nestjs|koa|fastapi|phoenix|play|struts|blade|gin|beego|echo|fiber|rocket|actix|tornado|bottle|pyramid|cherrypy|hug|falcon|sanic|fastify|hapi|loopback|feathers|adonisjs|sails|aurelia|alpine.*js|stimulus|litelement|stencil|nest.*js)\b`)

	// Find all matches in the text
	matches := re.FindAllString(text, -1)
	if len(matches) == 0 {
		return []string{"coding frameworks not found"}
	}
	return DeduplicateSliceContents(matches).([]string)
}

// Parses database types from the given string in a case-insensitive manner.
func ParseDatabaseTypes(text string) []string {
	// Define a regex pattern to match common database types (case-insensitive)
	re := regexp.MustCompile(`(?i)\b(no.*sql|ms.*sql|mysql|postgre.*sql|sqlite|mongodb|oracle|sql.*server|maria.*db|redis|cassandra|elasticsearch|firebase|dynamo.*db|couch.*db|neo4j|hbase|memcached|couchbase|db2|teradata|snowflake|bigquery|amazon.*aurora|amazon.*redshift|google.*cloud.*spanner|microsoft.*access|informix|ingres|interbase|firebird|sybase|volt.*db|greenplum)\b`)

	// Find all matches in the text
	matches := re.FindAllString(text, -1)
	if len(matches) == 0 {
		return []string{"database types not found"}
	}
	return DeduplicateSliceContents(matches).([]string)
}

// Parses database types from the given string in a case-insensitive manner.
func ParseAWSServiceNames(text string) []string {
	// Define a regex pattern to match common database types (case-insensitive)
	re := regexp.MustCompile(`(?i)\b(ec2|s3|rds|lambda|dynamodb|ecs|eks|cloudfront|route.*53|vpc|iam|cloudwatch|sns|sqs|redshift|glue|athena|kinesis|elastic.*beanstalk|cloudformation|elastic.*load.*balancing|sagemaker|fargate|aurora|elasticache|kinesis|step.*functions|appsync|codepipeline|codebuild|codedeploy|codecommit|amplify|lightsail|batch|opsworks|cloudtrail|cloud9|cloudhsm|direct.*connect|elastic.*transcoder|guardduty|inspector|macie|organizations|secrets.*manager|security.*hub|shield|waf|workspaces)\b`)

	// Find all matches in the text
	matches := re.FindAllString(text, -1)
	if len(matches) == 0 {
		return []string{"database types not found"}
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
		return []string{"no links found"}
	}
	return dedupedLinks
}

// Deduplicates a slice of any type and converts all strings to lowercase.
func DeduplicateSliceContents(slice interface{}) interface{} {
	// Use reflection to get the value and type of the input slice
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return []string{"input is not a slice"}
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
