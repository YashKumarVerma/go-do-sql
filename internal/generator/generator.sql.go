package generator

import (
	"strconv"
	"strings"

	"github.com/YashKumarVerma/go-do-sql/internal/parser"
	"github.com/YashKumarVerma/go-do-sql/internal/shell"
)

func commandToSQL(column parser.StructuredCommandData, tableName string) (string, []string) {
	constraints := make([]string, 0)
	sql := "\t"
	sql += column.Name + " "
	sql += column.Datatype + " "

	// handle null token
	if column.Null == true {
		sql += "NULL "
	} else {
		sql += "NOT NULL "
	}

	// handle auto-increment
	if column.AutoIncrement == true {
		sql += "AUTO_INCREMENT "
	}

	if column.DefaultValue != "" {
		sql += "DEFAULT " + writeAsPerDataType(column, column.DefaultValue)
	}

	// handle length token
	sql = strings.ReplaceAll(sql, "__LENGTH__", strconv.Itoa(column.Length))

	// generating constraints
	if column.Primary {
		primaryConstraint := "\tCONSTRAINT " + strings.ToLower(tableName) + "_primary_key PRIMARY KEY (" + column.Name + ")"
		constraints = append(constraints, primaryConstraint)
	}

	if column.Unique {
		uniqueConstraint := "\tCONSTRAINT unique_" + strings.ToLower(column.Name) + " UNIQUE (" + column.Name + ")"
		constraints = append(constraints, uniqueConstraint)
	}

	if column.Foreign != "" {
		// --foreign:users(id)
		foreignConstraint := "\tCONSTRAINT foreign_" + strings.ToLower(column.Name) + " FOREIGN KEY (" + column.Name + ") "
		foreignConstraint += "REFERENCES " + column.Foreign + " "
		constraints = append(constraints, foreignConstraint)
	}

	return sql, constraints
}

// function to assign default values based on datatype based on datatype
func applyDefaults(column parser.StructuredCommandData, tableName string) (parser.StructuredCommandData, string) {
	if column.Length == 0 {
		switch column.Datatype {
		case "INT(__LENGTH__)":
			{
				column.Length = 8
			}
		case "VARCHAR(__LENGTH__)":
			{
				column.Length = 64
			}
		}
	}
	return column, tableName
}

// writeSql : function to generate table generation commands and return s string
func writeSQL() (string, string) {
	// load data from other modules
	tableName := shell.TableName
	schema := parser.ProcessedCommands

	// filenames to access data from
	outputFile := strings.ToLower(tableName) + ".table.txt"

	// append data to generate final string
	var sqlString string
	constraintsStorage := make([]string, 0)
	sqlString += "CREATE TABLE __TABLE__ (\n"
	for _, command := range schema {
		// generate direct sql command and get constraints
		sql, constraints := commandToSQL(applyDefaults(command, tableName))
		sqlString += sql

		// if counter != len(schema)-1 {
		sqlString += ", \n"
		// }
		// store constrains in local array
		for _, constraint := range constraints {
			constraintsStorage = append(constraintsStorage, constraint)
		}
	}

	// append constraints towards end of table declaration
	for i, constraint := range constraintsStorage {
		sqlString += constraint
		if i != len(constraintsStorage)-1 {
			sqlString += ", \n"
		}
	}
	sqlString += "\n);"

	// inject table name
	sqlString = strings.ReplaceAll(sqlString, "__TABLE__", tableName)

	// write template file to directory
	return outputFile, sqlString
}
