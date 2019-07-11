package commands_controller

import (
	"fmt"

	log "github.com/sirupsen/logrus"
)

type Command struct {
	UserId  string
	Command string
}

func NewCommand(userId, command string) Command {
	return Command{
		UserId:  userId,
		Command: command,
	}
}

type CommandsController struct {
	commandsChan <-chan Command
}

func NewCommandsComnroller(commandsChan chan Command) *CommandsController {
	return &CommandsController{}
}

func (cc *CommandsController) Listen() {
	for cmd := range cc.commandsChan {
		log.Debug(fmt.Sprintf("New command in controller: %v, by user: %v", cmd.Command, cmd.UserId))
	}
}
