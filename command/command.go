package command

import (
	"strings"
	"sync"

	"github.com/igungor/tlbot"
)

// A Command is an implementation of a bot command.
type Command struct {
	// Name of the command without the leading slash.
	Name string

	// Short description of the command
	ShortLine string

	// Hidden enables commands to be hidden from the /help output, such as
	// Telegram's built-in commands and easter eggs.
	Hidden bool

	// Run runs the command.
	Run func(bot *tlbot.Bot, msg *tlbot.Message)
}

var (
	// mu guards commands-map access
	mu       sync.Mutex
	commands = make(map[string]*Command)
)

func register(cmd *Command) {
	mu.Lock()
	defer mu.Unlock()

	commands[cmd.Name] = cmd
}

// Lookup looks-up name from registered commands and returns
// corresponding Command if any.
func Lookup(cmdname string) *Command {
	mu.Lock()
	defer mu.Unlock()

	cmdname = strings.TrimSuffix(cmdname, "@ilberbot")
	cmd, ok := commands[cmdname]
	if !ok {
		return nil
	}
	return cmd
}
