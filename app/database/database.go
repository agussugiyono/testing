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

	// Add foreign key constraints manually
	if err := db.Exec("ALTER TABLE classes ADD CONSTRAINT fk_classes_users FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT ON UPDATE RESTRICT").Error; err != nil {
		return err
	}

	if err := db.Exec("ALTER TABLE mentees ADD CONSTRAINT fk_mentees_classes FOREIGN KEY (class_id) REFERENCES classes(id) ON DELETE RESTRICT ON UPDATE RESTRICT").Error; err != nil {
		return err
	}

	if err := db.Exec("ALTER TABLE logs ADD CONSTRAINT fk_logs_users FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT ON UPDATE RESTRICT").Error; err != nil {
		return err
	}
	if err := db.Exec("ALTER TABLE logs ADD CONSTRAINT fk_logs_classes FOREIGN KEY (class_id) REFERENCES classes(id) ON DELETE RESTRICT ON UPDATE RESTRICT").Error; err != nil {
		return err
	}
	if err := db.Exec("ALTER TABLE logs ADD CONSTRAINT fk_logs_mentees FOREIGN KEY (mentee_id) REFERENCES mentees(id) ON DELETE RESTRICT ON UPDATE RESTRICT").Error; err != nil {
		return err
	}

	return nil
}
