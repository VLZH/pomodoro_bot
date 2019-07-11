package main

import (
	"pomodoro_mod/commands_controller"
	"pomodoro_mod/commands_service"
)

const (
	BOT_TOKEN = "831869694:AAGo5m5OsiCapzFyDfIRlov30SOVfyiUEks"
)

func main() {
	commandsChan := make(chan commands_controller.Command)
	controler := commands_controller.NewCommandsComnroller(commandsChan)
	commandsService := commands_service.NewTelegramCommandsService(BOT_TOKEN, commandsChan)
	// start listen messages from service
	controler.Listen()
	// start listen messages from telegram
	commandsService.Listen()
}
