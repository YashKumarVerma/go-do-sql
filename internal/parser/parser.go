package parser

import (
	"strconv"
	"strings"

	ui "github.com/YashKumarVerma/go-lib-ui"
)

// StructuredCommandData represents everything a single command represents
type StructuredCommandData struct {
	Name          string
	Primary       bool
	Datatype      string
	AutoIncrement bool
	Length        int
	Unique        bool
	DefaultValue  string
	Foreign       string
	Null          bool
}

// validCommands : list of all commands that are accepted by shell
var validCommands = []string{
	"--name",
	"--primary",
	"--auto_increment",
	"--type",
	"--length",
	"--unique",
	"--foreign",
	"--null",
	"--default"}

// auxillary function to check if item exists in particular array
func in(haystack []string, needle string) bool {
	for _, i := range haystack {
		if needle == i {
			return true
		}
	}
	return false
}

// function to check if given schema is valid as per grammar
func checkIfGrammarCorrect(command string) bool {
	primaryCommand := strings.Split(command, " ")
	for _, value := range primaryCommand {
		if !in(validCommands, strings.Split(value, ":")[0]) {
			return false
		}
	}
	return true
}

// parseAllCommandData from schema to struct
func parseAllCommandData(command string) StructuredCommandData {
	var data StructuredCommandData
	tokens := strings.Split(command, " ")

	for _, value := range tokens {
		var val string
		keyValue := strings.Split(value, ":")
		key := strings.ReplaceAll(keyValue[0], "--", "")
		if len(keyValue) > 1 {
			val = keyValue[1]
		} else {
			val = "true"
		}

		switch key {

		case "name":
			{
				data.Name = val
			}
		case "primary":
			{
				data.Primary = val == "true"
			}
		case "type":
			{
				switch val {
				case "int":
					{
						data.Datatype = "INT(__LENGTH__)"
					}
				case "bool":
					{
						data.Datatype = "INT(1)"
					}
				case "datetime":
					{
						data.Datatype = "DATETIME"
					}
				case "string", "default":
					{
						data.Datatype = "VARCHAR(__LENGTH__)"
					}
				}
				data.Datatype = val
			}
		case "auto_increment":
			{
				data.AutoIncrement = val == "true"
			}
		case "length":
			{
				numericValue, _ := strconv.Atoi(val)
				data.Length = numericValue
			}
		case "unique":
			{
				data.Unique = val == "true"
			}
		case "null":
			{
				data.Null = val == "true"
			}
		case "foreign":
			{
				data.Foreign = val
			}
		case "default":
			{
				data.DefaultValue = val
			}
		}
	}
	return data
}

// GetStructuredCommands returns result from shell to generate a detailed schema
func GetStructuredCommands(commands []string) []StructuredCommandData {
	ui.ContextPrint("magnifying_glass_tilted_left", "Total "+strconv.Itoa(len(commands))+" columns entered.")
	// expand each schema and generate a structure of all processed data
	entitySchemas := make([]StructuredCommandData, 0)
	for _, command := range commands {
		if !checkIfGrammarCorrect(command) {
			ui.ContextPrint("fire", "Invalid syntax of command in "+command)
		} else {
			// ui.Info("parsing all data from command")
			entitySchemas = append(entitySchemas, parseAllCommandData(command))
		}
	}

	// return array of structures
	return entitySchemas
}

// Initialize sql parser
func Initialize() {
	ui.ContextPrint("brain", "Parsing SQL")
}
