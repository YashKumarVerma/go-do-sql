package main

import (
	"github.com/YashKumarVerma/go-do-sql/internal/config"
	"github.com/YashKumarVerma/go-do-sql/internal/shell"
	ui "github.com/YashKumarVerma/go-lib-ui"
)

func main() {
	ui.ContextPrint("dolphin", "SQL Helper !")
	config.Initialize()
	shell.Initialize()
	shell.Display()
}
