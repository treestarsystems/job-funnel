package utils

import (
	"flag"
	"fmt"
	"os"
)

var EnvFilePath string

func InitCommandLineFlags() {
	// Define the flags
	h := flag.String("h", "", "Show help")
	e := flag.String("e", "./.env", "Path to file containing Environment variables for this application")

	// Parse the flags
	flag.Parse()

	// Show help and exit if the help flag is provided
	if *h != "" {
		flag.Usage()
		os.Exit(0)
	}

	// Set the default value of the 'e' flag to "./.env"
	if *e == "" {
		fmt.Println("No path provided. Using deault value(s).")
		EnvFilePath = "./.env"
	}

	// Check if any flags are provided
	if len(os.Args) <= 1 {
		fmt.Println("No flags provided. Using deault value(s).")
		EnvFilePath = "./.env"
	}
}
