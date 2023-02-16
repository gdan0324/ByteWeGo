package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/logging/logrus"
	"time"
)

var DB *gorm.DB

func Init() {
	InitDB()
}

func InitDB() {
	var err error
	gormlogrus := logger.New(
		logrus.NewWriter(),
		logger.Config{
			SlowThreshold: time.Millisecond,
			Colorful:      false,
			LogLevel:      logger.Info,
		},
	)
	//dsn := "root:n7Zs3usIM15HlkvQ@tcp(120.46.190.10:3306)/bytewego?charset=utf8&parseTime=True&loc=Local"
	dsn := "root:1223@tcp(41065a2h49.zicp.fun:20128)/bytewego?charset=utf8&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt: true,
			Logger:      gormlogrus,
		},
	)
	if err != nil {
		panic(err)
	}
}
