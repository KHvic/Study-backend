package setting

import (
	"log"
	"os"
	"time"

	"github.com/go-ini/ini"
)

var (
	cfg *ini.File
	// AppSetting ...
	AppSetting = &App{}
	// ServerSetting ...
	ServerSetting = &Server{}
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

// Server ...
type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
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
	env := os.Getenv("quiz_env")
	if env == "production" {
		cfg, err = ini.Load("conf/production.ini")
	} else {
		cfg, err = ini.Load("conf/staging.ini")
	}

	if err != nil {
		log.Fatalf("setting.Setup failed, err: %v", err)
	}
	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("database", DatabaseSetting)
}

func mapTo(section string, v interface{}) {
	if err := cfg.Section(section).MapTo(v); err != nil {
		log.Fatalf("Cfg.Mapto failed, err: %v", err)
	}
}
