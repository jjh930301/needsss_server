package database

import (
	"fmt"
	"os"

	"github.com/jjh930301/market/common/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDb() *gorm.DB {
	DB = connection()
	return DB
}

func connection() *gorm.DB {
	var user string
	if os.Getenv("MYSQL_USER") == "" {
		user = "root"
	} else {
		user = os.Getenv("MYSQL_USER")
	}
	uri := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		user,
		os.Getenv("MYSQL_ROOT_PASSWORD"),
		"mysql",
		3306,
		os.Getenv("MYSQL_DATABASE"),
	)
	db, err := gorm.Open(mysql.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		fmt.Println("error:::", err)
		panic(err)
	}
	if os.Getenv("ENV") != "production" {
		db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(
			&models.UserModel{},
			&models.KrTickerModel{},
			&models.InterestedTickerModel{},
			&models.KrTickerChartsModel{},
			&models.KrTickerCommentModel{},
		)
	}

	return db
}
