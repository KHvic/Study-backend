package dao

import (
	"fmt"
	"log"
	
	"github.com/KHvic/quiz-backend/pkg/setting"
	"github.com/jinzhu/gorm"

	// SQL driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

// Setup initialize the database
func Setup() {
	var err error
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		log.Fatalf("dao.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

// CloseDB closes database connection
func CloseDB() {
	defer db.Close()
}
