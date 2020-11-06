package generator

import (
	"fmt"
	"strings"

	"github.com/YashKumarVerma/go-do-sql/internal/parser"
	"github.com/YashKumarVerma/go-do-sql/internal/shell"
	ui "github.com/YashKumarVerma/go-lib-ui"
)

// generateTemplate to insert data
func generateTemplate() (string, string) {
	// load data from other modules
	tableName := shell.TableName
	schema := parser.ProcessedCommands

	// filenames to access data from
	outputFile := "template." + strings.ToLower(tableName) + ".txt"

	sql := fmt.Sprintf("INSERT INTO %s (", tableName)
	for counter, column := range schema {
		if column.AutoIncrement == false {
			sql += column.Name
			if counter != len(schema)-1 {
				sql += ","
			}
		}
	}
	sql += ") VALUES (__DATA_HERE__);"
	ui.ContextPrint("construction", "Building insertion template for : "+shell.TableName)
	templateString = sql
	return outputFile, sql
}
