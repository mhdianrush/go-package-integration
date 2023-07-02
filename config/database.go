package config

import (
	"fmt"

	"github.com/mhdianrush/go-package-integration/model"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// to concate string
	connectionString := fmt.Sprintf("%v:%v@/%v?parseTime=true", ENV.DB_USERNAME, ENV.DB_PASSWORD, ENV.DB_DATABASE)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&model.User{})

	DB = db

	logger := logrus.New()
	logger.Println("Database Connected")
}
