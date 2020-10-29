package generator

import (
	"os"

	"github.com/YashKumarVerma/go-do-sql/internal/parser"
	"github.com/YashKumarVerma/go-do-sql/internal/shell"
	ui "github.com/YashKumarVerma/go-lib-ui"
)

// Initialize code generator
func Initialize() {
	ui.ContextPrint("wrench", "Building code")
	ui.ContextPrint("construction", "Building table : "+shell.TableName)
	writeDataToDisk(writeSQL())
}

func writeAsPerDataType(column parser.StructuredCommandData, data string) string {

	if column.Datatype == "int" || column.Datatype == "boolean" {
		return data
	}

	return "'" + data + "'"
}

// WriteDataToDisk to write data to generated directory
func writeDataToDisk(filename string, dataToWrite string) bool {
	// ensure that generated directory exists
	_, err := os.Stat("./output")
	if os.IsNotExist(err) {
		dirError := os.MkdirAll("./output", 0755)
		ui.CheckError(dirError, "cannot create output directory", true)
		ui.ContextPrint("open_file_folder", "Output directory created")
	}

	// write data to file
	codeFile, err := os.Create("./output/" + filename)
	ui.CheckError(err, "Error creating "+filename, true)

	_, err = codeFile.WriteString(dataToWrite)
	ui.CheckError(err, "Error writing to "+filename, true)

	fileCloseError := codeFile.Close()
	ui.CheckError(fileCloseError, "Error closing file !", true)

	return true
}