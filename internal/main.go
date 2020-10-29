package main

import (
	"github.com/YashKumarVerma/go-do-sql/internal/config"
	"github.com/YashKumarVerma/go-do-sql/internal/shell"
)

func main() {
	config.Initialize()
	shell.Initialize()
	shell.Display()
}
