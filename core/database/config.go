package database

import (
	"fmt"
	"github.com/magiskboy/katalog/core"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func getDatabaseDSN() string {
	host := core.Getenv("MYSQL_HOST", "localhost")
	port := core.Getenv("MYSQL_PORT", "3306")
	user := core.Getenv("MYSQL_USER", "katalog")
	pass := core.Getenv("MYSQL_PASS", "password")
	db := core.Getenv("MYSQL_DB", "katalog")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, db)

	return dsn
}

// Get get a new connection for database
func GetDatabase() *gorm.DB {
	env := core.Getenv("ENV", "testing")

	config := gorm.Config{}

	if env == "testing" {
		db, err := gorm.Open(sqlite.Open(":memory:"), &config)

		if err != nil {
			panic(err)
		}

		return db
	}

	dsn := getDatabaseDSN()
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &config)

	if err != nil {
		panic(err)
	}

	return db
}
