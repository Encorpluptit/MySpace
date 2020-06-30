package database

import (
	"MySpace-Api/models"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

type Database struct {
	DB       *gorm.DB
	Driver   string
	Name     string
	Host     string
	Port     string
	User     string
	Password string
}

func (db *Database) prepareEnv() (err error) {
	//if godotenv.Load() != nil {
	//	log.Fatalf("Cannot Load .env file.")
	//}
	db.Driver = os.Getenv("DB_DRIVER")
	db.Name = os.Getenv("DB_NAME")
	db.Host = os.Getenv("DB_HOST")
	db.Port = os.Getenv("DB_PORT")
	db.User = os.Getenv("DB_USER")
	db.Password = os.Getenv("DB_PASSWORD")
	return nil
}

func (db *Database) Init() (err error) {
	if err := db.prepareEnv(); err != nil {
		return err
	}
	if db.Driver == "mysql" {
		url := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", db.User, db.Password, db.Host, db.Port, db.Name)
		db.DB, err = gorm.Open(db.Driver, url)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", db.Driver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", db.Driver)
		}
	}
	if db.Driver == "postgres" {
		url := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", db.Host, db.Port, db.User, db.Name, db.Password)
		db.DB, err = gorm.Open(db.Driver, url)
		if err != nil {
			fmt.Printf("Cannot connect to %s database", db.Driver)
			log.Fatal("This is the error:", err)
		} else {
			fmt.Printf("We are connected to the %s database", db.Driver)
		}
	}
	db.DB.Debug().AutoMigrate(&models.User{}) //database migration
	return nil
}

// Destroy d
func (db *Database) Destroy() {
	if err := db.DB.Close(); err != nil {
		log.Printf("Error on Database Close %v", err)
	}
}
