package shell

import (
	"bufio"
	"os"
	"strings"

	ui "github.com/YashKumarVerma/go-lib-ui"

	"github.com/c-bata/go-prompt"
)

// AutoComplete suggestions from schema
func autoComplete(document prompt.Document) []prompt.Suggest {
	suggestions := []prompt.Suggest{
		{Text: "--name", Description: "name of column in table"},
		{Text: "--primary", Description: "to apply primary key constraint"},
		{Text: "--auto_increment", Description: "auto increment value of column"},

		{Text: "--type:int", Description: "datatype of column as integer"},
		{Text: "--type:string", Description: "datatype of column as string"},
		{Text: "--type:boolean", Description: "datatype of column as boolean"},
		{Text: "--type:datetime", Description: "datatype of column as date and time"},

		{Text: "--length:10", Description: "if set, will pass length parameter inside datatype"},
		{Text: "--unique", Description: "add unique constraint"},
		{Text: "--foreign:users(id)", Description: "link reference to another column"},
		{Text: "--null", Description: "set to allow null values, else not null by default"},
		{Text: "--default:val", Description: "set default value"},

		{Text: "--fill:name", Description: "insert random name in this column"},
		{Text: "--fill:email", Description: "insert random email in this column"},
		{Text: "--fill:address", Description: "insert random address in this column"},
		{Text: "--fill:city", Description: "insert random city in this column"},
		{Text: "--fill:country", Description: "insert random country in this column"},
		{Text: "--fill:zip", Description: "insert random zip in this column"},
		{Text: "--fill:car-maker", Description: "insert random car-maker in this column"},
		{Text: "--fill:color", Description: "insert random color in this column"},
		{Text: "--fill:url", Description: "insert random url in this column"},
		{Text: "--fill:animal", Description: "insert random animal in this column"},
		{Text: "--fill:int", Description: "insert random int in this column"},
	}
	return prompt.FilterHasPrefix(suggestions, document.GetWordBeforeCursor(), true)
}

// ColumnStorage to store entered columns
var ColumnStorage []string

// TableName to which said columns belong
var TableName string

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
	// commands to test without entering manually
	// ColumnStorage = append(ColumnStorage, "--name:id --type:int --auto_increment --primary")
	// ColumnStorage = append(ColumnStorage, "--name:name --type:string --length:64")
	// ColumnStorage = append(ColumnStorage, "--name:age --type:int")
	// ColumnStorage = append(ColumnStorage, "--name:dob --type:datetime")
	// ColumnStorage = append(ColumnStorage, "--name:mobile --type:string --null")
	// ColumnStorage = append(ColumnStorage, "--name:country --type:string --length:2 --default:IN")
	// ColumnStorage = append(ColumnStorage, "--name:aadhar --type:string --length:16 --default:0000000000000000 --unique")
	// ColumnStorage = append(ColumnStorage, "--name:email --type:string --length:32 --default:unknown@gmail.com --foreign:students(email)")

}

// Initialize command history storage
func Initialize() {
	// input relation name from user
	ui.ContextPrint("package", "Enter table name")
	ui.ContextPrint("package", "")
	reader := bufio.NewReader(os.Stdin)
	relationName, err := reader.ReadString('\n')
	ui.CheckError(err, "Error reading table/relation name", true)
	relationName = strings.Trim(relationName, " ")
	relationName = strings.Trim(relationName, "\n")

	// launch shell for user
	ui.ContextPrint("spiral_shell", "Launching shell")
	ColumnStorage = make([]string, 0)
	TableName = relationName
}
