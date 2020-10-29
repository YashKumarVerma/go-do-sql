package shell

import (
	"strings"

	"github.com/c-bata/go-prompt"
)

// AutoComplete suggestions from schema
func autoComplete(document prompt.Document) []prompt.Suggest {
	suggestions := []prompt.Suggest{
		{Text: "--name", Description: "set name of column in table"},
		{Text: "--primary", Description: "set to apply primary key constraint"},
		{Text: "--auto_increment", Description: "auto increment value of column"},

		{Text: "--type:int", Description: "datatype of column as integer"},
		{Text: "--type:string", Description: "datatype of column as string"},
		{Text: "--type:boolean", Description: "datatype of column as boolean"},
		{Text: "--type:datetime", Description: "datatype of column as date and time"},

		{Text: "--length:10", Description: "if set, will pass length parameter inside datatype"},
		{Text: "--unique", Description: "datatype of column as boolean"},
		{Text: "--foreign:id|users(id)", Description: "id references id of table users"},
		{Text: "--null", Description: "set to allow null values, else not null by default"},
		{Text: "--default:val", Description: "default value of field"},
	}
	return prompt.FilterHasPrefix(suggestions, document.GetWordBeforeCursor(), true)
}

// ColumnStorage to store entered columns
var ColumnStorage []string

// Display shell onto terminal
func Display() {
	// initialize shell with auto-complete feature
	commandHistory := make([]string, 0)
	commandHistory = append(commandHistory, "exit")
	for true {
		command := prompt.Input(" > ", autoComplete,
			prompt.OptionTitle("SQL Helper"),
			prompt.OptionHistory(commandHistory),
			prompt.OptionPrefixTextColor(prompt.Yellow),
			prompt.OptionPreviewSuggestionTextColor(prompt.Blue),
			prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
			prompt.OptionSuggestionBGColor(prompt.DarkGray))
		if command == "exit" {
			break
		} else {
			if strings.TrimSpace(command) != "" {
				commandHistory = append(commandHistory, strings.TrimSpace(command))
				ColumnStorage = append(ColumnStorage, strings.TrimSpace(command))
			}
		}
	}
}

// Initialize command history storage
func Initialize() {
	ColumnStorage = make([]string, 0)
}
