package setting

import (
	"github.com/go-ini/ini"
	"log"
)

var (
	cfg *ini.File
	DatabaseSetting = &Database{}
)

// Database ...
type Database struct {
	Type	string
	User	string
	Password 	string
	Host	string
	Name	string
	TablePrefix	string
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