package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/pscompsci/bot/internal/bot"
)

func main() {
	cfg := loadConfiguration("settings.json")

	bot, err := bot.New(cfg)
	if err != nil {
		log.Fatal(err)
	}
	bot.Run()
}

func loadConfiguration(file string) bot.Config {
	var config bot.Config
	configFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer configFile.Close()

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)

	return config
}
