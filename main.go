package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type ConfigStruct struct {
	AWSSecretKey string `json:"aws_secret_key"`
	AWSAccessKey string `json:"aws_access_key"`
	MongoURI     string `json:"mongo_uri"`
}

func getConfigValues(reader *bufio.Reader) (*ConfigStruct, error) {
	var config ConfigStruct

	fmt.Print("AWS Access Key: ")
	awsAccessKey, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	config.AWSAccessKey = strings.TrimSpace(awsAccessKey)

	fmt.Print("AWS Secret Key: ")
	awsSecretKey, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	config.AWSSecretKey = strings.TrimSpace(awsSecretKey)

	fmt.Print("MongoDB URI: ")
	mongoURI, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	config.MongoURI = strings.TrimSpace(mongoURI)

	if config.AWSAccessKey == "" || config.AWSSecretKey == "" || config.MongoURI == "" {
		return nil, fmt.Errorf("incomplete configuration values")
	}

	return &config, nil
}
func createOrCheckConfig() {
	homeDir, _ := os.UserHomeDir()
	configDir := filepath.Join(homeDir, ".goawsutil")
	configFilePath := filepath.Join(configDir, "config.js")

	file, err := os.ReadFile(configFilePath)
	if err != nil {
		fmt.Println("Config file for AWS and MongoDB URI not found. Enter 'y' to create or 'n' to exit:")
		reader := bufio.NewReader(os.Stdin)
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)

		switch userInput {
		case "y":
			configContent, err := getConfigValues(reader)
			if err != nil {
				fmt.Println(err)
				return
			}

			configBytes, err := json.Marshal(configContent)
			if err != nil {
				fmt.Println("Failed to marshal config struct to JSON:", err)
				return
			}
			if err := os.MkdirAll(configDir, 0700); err != nil {
				fmt.Println("Failed to create config directory:", err)
				return
			}
			if err := os.WriteFile(configFilePath, configBytes, 0644); err != nil {
				fmt.Println("Failed to write config file:", err)
				return
			}

			fmt.Println("Config file created successfully.")

		case "n":
			fmt.Println("You can't use this utility without a config file. Exiting...")
			return

		default:
			fmt.Println("Invalid input. Exiting...")
			return
		}
	} else {
		//verifying the config file have all contents or not
		var configToVerify ConfigStruct
		errrr := json.Unmarshal(file, &configToVerify)
		if errrr != nil {
			log.Fatal(errrr)
			errrr = os.Remove(configFilePath)
			log.Fatal(errrr)
		}
		fmt.Print(configToVerify)
	}
}
func main() {
	createOrCheckConfig()
}
