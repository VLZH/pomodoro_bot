package main

import (
	"log"
	"os"
	"pomodoro_mod/commands_controller"
	"pomodoro_mod/commands_service"

	"github.com/joho/godotenv"
)

func loadDotenv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cannot load .env file")
	}
}

func main() {
	loadDotenv()
	BotToken := os.Getenv("BOT_TOKEN")
	//
	commandsChan := make(chan commands_controller.Command)
	controler := commands_controller.NewCommandsComnroller(commandsChan)
	commandsService := commands_service.NewTelegramCommandsService(BotToken, commandsChan)
	// start listen commands from user(incoming from commandsSevice)
	controler.Listen()
	// start listen messages from telegram
	commandsService.Listen()
}
