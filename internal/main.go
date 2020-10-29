package main

import (
	"github.com/YashKumarVerma/go-do-sql/internal/config"
	ui "github.com/YashKumarVerma/go-lib-ui"
)

func main() {
	config.Init()
	ui.ContextPrint("fire", "Hello World")
}
