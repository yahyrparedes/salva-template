package database

import (
	"fmt"
	"github.com/yahyrparedes/salva-template/cmd/config"
	"github.com/yahyrparedes/salva-template/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"os"
)

var Connection *gorm.DB

func InitializeConnection() {
	Connection = ConnectionORM(CreateStringConnection())
	if config.IsLocal() {
		ActivateLogger(Connection)
	}
	ActivateAutoMigrate(Connection)
}

func ActivateAutoMigrate(c *gorm.DB) {
	fmt.Println("Running Migrations")
	c.AutoMigrate(&models.Brand{})
}

func ActivateLogger(DB *gorm.DB) {
	DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	DB.Logger = logger.Default.LogMode(logger.Info)
}

func CreateStringConnection() string {
	user := os.Getenv("database.user")
	password := os.Getenv("database.password")
	host := os.Getenv("database.host")
	dbname := os.Getenv("database.name")
	port := os.Getenv("database.port")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s ", host, user, password, dbname, port)
	return dsn
}

func ConnectionORM(stringConnection string) *gorm.DB {
	Connection, err := gorm.Open(postgres.Open(stringConnection), &gorm.Config{})
	if err != nil {
		fmt.Println("Database Connection Error", err)
		panic(err)
	}

	fmt.Println("🚀 Connected Successfully to the Database")
	return Connection
}

func InitializeTestConnection() *gorm.DB {
	var path = config.RootPathTest + "test.db"
	Connection = ConnectionTestORM(path)
	return Connection
}

func ConnectionTestORM(path string) *gorm.DB {
	Connection, err := gorm.Open(sqlite.Open(path))

	if err != nil {
		fmt.Println("Database Connection Error", err)
		panic(err)
	}

	fmt.Println("🚀 Connected Successfully to the Database")

	return Connection
}
