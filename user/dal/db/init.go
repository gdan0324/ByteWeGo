package db

import (
	"time"

	"github.com/gdan0324/ByteWeGo/api/pkg/consts"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/logging/logrus"
)

var DB *gorm.DB

// Init init DB
func Init() {
	var err error
	gormlogrus := logger.New(
		logrus.NewWriter(),
		logger.Config{
			SlowThreshold: time.Millisecond,
			Colorful:      false,
			LogLevel:      logger.Info,
		},
	)
	DB, err = gorm.Open(mysql.Open(consts.MySQLDefaultDSN),
		&gorm.Config{
			PrepareStmt: true,
			Logger:      gormlogrus,
		},
	)
	if err != nil {
		panic(err)
	}
}
