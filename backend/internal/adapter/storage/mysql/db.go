package mysql

import (
	"fmt"
	"hr_management/internal/adapter/config"
	"hr_management/internal/core/domain"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type database struct {
	DB *gorm.DB
}

var tables = []interface{}{
	&domain.User{},
	&domain.Leave{},
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

	fmt.Printf("Connecting to database with DSN: %s:***@tcp(%s:%s)/%s\n",
		config.Username, config.Host, config.Port, config.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Printf("Database connection failed: %v\n", err)
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
	fmt.Println("Starting database migration...")
	for i, table := range tables {
		fmt.Printf("Migrating table %d: %T\n", i+1, table)
		if err := db.DB.AutoMigrate(table); err != nil {
			fmt.Printf("Error migrating table %T: %v\n", table, err)
			return err
		}
		fmt.Printf("Successfully migrated table: %T\n", table)
	}
	fmt.Println("Database migration completed!")
	return nil
}
