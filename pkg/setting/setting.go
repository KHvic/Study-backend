package setting

import (
	"log"

	"github.com/go-ini/ini"
)

var (
	cfg *ini.File
	// AppSetting ...
	AppSetting = &App{}
	// DatabaseSetting ...
	DatabaseSetting = &Database{}
)

// App ...
type App struct {
	PrefixURL string

	RuntimeRootPath string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

// Database ...
type Database struct {
	Type        string
	User        string
	Password    string
	Host        string
	Name        string
	TablePrefix string
}

// Setup ...
func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup failed, err: %v", err)
	}

	mapTo("database", DatabaseSetting)
}

func mapTo(section string, v interface{}) {
	if err := cfg.Section(section).MapTo(v); err != nil {
		log.Fatalf("Cfg.Mapto failed, err: %v", err)
	}
}
