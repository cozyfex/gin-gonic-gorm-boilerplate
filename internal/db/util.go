package db

import (
	"fmt"
	"gin-gonic-gorm-boilerplate/configs"
	"gin-gonic-gorm-boilerplate/internal/util/logger"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func Connection(config configs.DBConfig) (*gorm.DB, error) {
	var db *gorm.DB
	var err error

	switch config.Type {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.DBName, config.CharSet)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", config.Host, config.User, config.Password, config.DBName, config.Port, config.SSLMode, config.Timezone)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(config.DBName), &gorm.Config{})
	case "sqlserver":
		dsn := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", config.User, config.Password, config.Host, config.Port, config.DBName)
		db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	}

	return db, err
}

func Close(d *gorm.DB) error {
	db, err := d.DB()
	if err != nil {
		logger.Error("failed to get db")
	}
	err = db.Close()
	if err != nil {
		logger.Error("failed to close db")
	}

	return err
}
