package mysql

import (
	"fmt"
	"hr_management/internal/adapter/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type database struct {
	DB *gorm.DB
}

var tables = []interface{}{
	// &d.User{},
}

func NewDatabase(config *config.DB) (*database, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return nil, err
	}
	fmt.Println("Database connected!")

	return &database{DB: db}, nil
}

func (db *database) Close() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		fmt.Println("Database close error -:", err)
		return err
	}
	sqlDB.Close()
	fmt.Println("Database close!")
	return nil
}

func (db *database) Migrate() error {
	tx := db.DB.Begin()
	for _, table := range tables {
		if err := db.DB.AutoMigrate(table); err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Error
}
