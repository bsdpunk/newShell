package shell

import (
	"fmt"
	//"github.com/bsdpunk/dndshell/shell/character"
	//"./character"
	"github.com/bsdpunk/newShell/shell/commands"
	//"github.com/bsdpunk/dndshell/shell/dice"
	"github.com/bsdpunk/newShell/shell/general"
	//	"github.com/bsdpunk/dndshell/shell/quotes"
	//"github.com/gobs/readline"
	"os"
	"strings"
)

var found string = "no"
var list []string
var cr character.Character
var matches = make([]string, 0, len(list))
var Qu quotes.Quotes
var coms = commands.Commands{
	{
		Name:      "Quit",
		ShortName: "quit",
		Usage:     "Exit the shell",
		Action:    general.End,
		Category:  "general",
	},
	{
		Name:      "Clear",
		ShortName: "clear",
		Usage:     "Clear the screen",
		Action:    general.Clear,
		Category:  "general",
	},
}

func AttemptedCompletion(text string, start, end int) []string {
	if start == 0 {
		return readline.CompletionMatches(text, CompletionEntry)
	} else {
		return nil
	}
}

func CompletionEntry(prefix string, index int) string {
	if index == 0 {
		matches = matches[:0]

		for _, w := range list {
			if strings.HasPrefix(w, prefix) {
				matches = append(matches, w)
			}
		}
	}

	if index < len(matches) {
		return matches[index]
	} else {
		return ""
	}
}
func NoAction() {
	fmt.Println("No action supplied with command")

}
func Shell(args []string) {
	character.Load()
	Qu.Load()
	for _, c := range coms {
		list = append(list, c.Name)
		list = append(list, c.ShortName)
	}
	//	cs := character.LoadClasses()
	//ms := monsters.Load()
	if len(args) > 1 {
		words := args
		if coms.HasCommand(words[1]) {
			cmd := coms.NameIs(words[1])
			if len(words) == 2 {
				if len(cmd.SubCommands) == 0 {
					cmd.Action()
				}

			} else {
				if cmd.SubCommands.HasCommand(words[2]) {
					cmd.Action()
				} else if !(cmd.SubCommands.HasCommand(words[2])) {
					if cmd.StringAction != nil {
						cmd.StringAction(words[2])
					}
				}
			}
		}
		os.Exit(0)
	}
	Qu.RandQ()
	prompt := "> "
	matches = make([]string, 0, len(list))

L:
	for {

		found = "no"
		readline.SetCompletionEntryFunction(CompletionEntry)
		readline.SetAttemptedCompletionFunction(nil)

		//readline.SetHistoryPath("~/.dnd_history")
		result := readline.ReadLine(&prompt)
		hist := *result
		readline.AddHistory(hist)
		//result.HistoryEnable()
		if result == nil {
			break L
		}

		input := *result
		words := strings.Fields(input)
		if len(words) > 0 && coms.HasCommand(words[0]) {
			cmd := coms.NameIs(words[0])
			if len(words) == 1 {
				if len(cmd.SubCommands) == 0 {
					cmd.Action()
				}

			} else {
				if cmd.SubCommands.HasCommand(words[1]) {
					cmd.Action()
				} else if !(cmd.SubCommands.HasCommand(words[1])) {
					if cmd.StringAction != nil {
						cmd.StringAction(words[1])
					}
				}
			}
		}
	}

	return
}
