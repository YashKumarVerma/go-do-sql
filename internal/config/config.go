package config

import (
	ui "github.com/YashKumarVerma/go-lib-ui"

	"github.com/spf13/viper"
)

// load configurations
func loadConfigurations() {
	viper.SetConfigName("default")
	viper.AddConfigPath("./config/")
	viper.SetConfigName("default")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	ui.CheckError(err, "config file not found", true)

	// set default configurations
	viper.SetDefault("app.name", "go-sql")
	viper.SetDefault("app.emoji", true)

	// save configurations onto exported object
	Load.Name = viper.GetString("app.name")
	Load.Emoji = viper.GetBool("app.emoji")
}

// Configuration store
type Configuration struct {
	Name  string
	Emoji bool
}

// Load configurations to be used from other modules
var Load Configuration

// Initialize config parser
func Initialize() {
	ui.ContextPrint("key", "Loading configurations")
	loadConfigurations()
}
