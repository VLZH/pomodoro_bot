package commands_service

import (
	"pomodoro_mod/commands_controller"
)

type CommandsService interface {
	Listen()
}

type TelegramCommandsService struct {
	Token        string
	commandsChan chan commands_controller.Command
}

func NewTelegramCommandsService(token string, commandsChan chan commands_controller.Command) *TelegramCommandsService {
	return &TelegramCommandsService{
		Token:        token,
		commandsChan: commandsChan,
	}
}

func (tg *TelegramCommandsService) Listen() {

}
