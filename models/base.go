package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB //database

func init() {

	dbUri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", "root", "test", "db", "3306", "test") //Build connection string
	fmt.Println(dbUri)

	conn, err := gorm.Open("mysql", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	//db.Debug().AutoMigrate(&Account{}, &Contact{}) //Database migration
}

//returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
