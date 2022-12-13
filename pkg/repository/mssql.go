package repository

import (
	//"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	DBtype   string
}


// Connect to database
func NewSqlDB(cfg Config) (*gorm.DB, error) {
	 dsn := fmt.Sprintf("%s://%s:%s@%s:%s?database=%s&connection+timeout=30", cfg.DBtype, cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)
	 db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	//  pingErr,_ := db.DB()
	//  if pingErr !=nil{
	// 	return nil, err
	//  }

	log.Printf("Connected!\n")
	return db, nil

}
