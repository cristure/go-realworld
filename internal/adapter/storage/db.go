package storage

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/go-realworld/internal/core/domain"
)

var DBURL string

func init() {
	DBHost := os.Getenv("DB_HOST")
	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASSWORD")
	DBName := os.Getenv("DB_NAME")
	DBPort := os.Getenv("DB_PORT")

	DBURL = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DBUser, DBPassword, DBHost, DBPort, DBName)
}

func New() (*gorm.DB, error) {
	DB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       DBURL, // data source name
		DefaultStringSize:         256,   // default size for string fields
		DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,  // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // autoconfigure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed to connect database: %w", err)
	}

	err = DB.AutoMigrate(&domain.User{}, &domain.Profile{}, &domain.Tag{}, &domain.Article{})
	if err != nil {
		return nil, fmt.Errorf("failed to auto migrate database: %w", err)
	}

	return DB, nil
}
