package generator

import (
	"strconv"
	"strings"

	"github.com/YashKumarVerma/go-do-sql/internal/parser"
	"github.com/YashKumarVerma/go-do-sql/internal/shell"
	ui "github.com/YashKumarVerma/go-lib-ui"
	"github.com/brianvoe/gofakeit/v5"
)

// populateTemplate to inject data into columns
func populateTemplate() (string, string) {

	// load data from other modules
	tableName := shell.TableName
	schema := parser.ProcessedCommands
	finalFileContent := ""
	i := 0
	for i < 10 {
		dataString := ""
		for counter, command := range schema {
			if command.AutoIncrement == false {
				dataString += writeAsPerDataType(command, getRandomItem(command, counter))
				if counter != len(schema)-1 {
					dataString += ","
				}
			}
		}
		filledTemplate := strings.ReplaceAll(templateString, "__DATA_HERE__", dataString)
		finalFileContent += filledTemplate + "\n"
		ui.ContextPrint("construction", filledTemplate)
		i++
	}

	// filenames to access data from
	outputFile := strings.ToLower(tableName) + ".insert.filled.txt"
	ui.ContextPrint("construction", "Generating data for template : "+shell.TableName)
	return outputFile, finalFileContent
}

func getRandomItem(command parser.StructuredCommandData, counter int) string {
	if command.Fill == "" {
		return "_"
	}
	gofakeit.Seed(0)

	return command.Fill + "_" + strconv.Itoa(counter)
}
