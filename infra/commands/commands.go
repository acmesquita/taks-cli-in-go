package commands

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

var commands = []string{
	"add",
	"update",
	"delete",
	"list",
	"get",
	"mark-done",
	"mark-in-progress",
	"help",
}

func ParseCommands(args []string) (string, []string) {
	validateCommands(args)

	return args[0], args[1:]
}

func HandleHelperMessage() {
	fmt.Println("Usage: task-cli <command> <args...>")
	fmt.Println("Commands:")
	for _, command := range commands {
		fmt.Println("-", command)
	}
}

func validateCommands(args []string) {
	if len(args) == 0 || !slices.Contains(commands, args[0]) {
		HandleHelperMessage()
		os.Exit(1)
	}
}

func ParseOptions(command string, args []string) map[string]string {
	options := parseOptions(args)
	return options
}

func parseOptions(args []string) map[string]string {
	options := make(map[string]string)

	for index, arg := range args {
		key := ""
		value := ""

		if after, ok := strings.CutPrefix(arg, "--"); ok {
			key = strings.TrimSpace(after)
		} else if after, ok := strings.CutPrefix(arg, "-"); ok {
			key = strings.TrimSpace(after)
		}

		if index+1 < len(args) {
			value = args[index+1]
		}

		options[key] = value

		if index+1 == len(args) {
			break
		}
	}

	return options
}
