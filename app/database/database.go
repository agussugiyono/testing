package database

import (
	"fmt"

	"github.com/agussugiyono/coba/app/config"

	classdata "github.com/agussugiyono/coba/features/class/data"
	logdata "github.com/agussugiyono/coba/features/log/data"
	manteedata "github.com/agussugiyono/coba/features/mentee/data"
	userdata "github.com/agussugiyono/coba/features/user/data"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDBMysql(cfg *config.AppConfig) *gorm.DB {

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cfg.DB_USERNAME, cfg.DB_PASSWORD, cfg.DB_HOSTNAME, cfg.DB_PORT, cfg.DB_NAME)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}

//	func InitialMigration(db *gorm.DB) {
//		db.AutoMigrate(&data.User{}, &data.Class{}, &data.Mentee{}, &data.Log{})
//	}
func InitialMigration(db *gorm.DB) error {
	err := db.AutoMigrate(
		&userdata.User{},
		&classdata.Class{},
		&manteedata.Mentee{},
		&logdata.Log{},
	)
	if err != nil {
		return err
	}

	return nil
}
