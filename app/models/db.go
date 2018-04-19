package models

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	//
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var isConnected bool
var dbName = getEnvWithDefault("DBNAME", "nondappsample")
var dbHost = getEnvWithDefault("DBHOST", "localhost")
var dbUser = getEnvWithDefault("DBUSER", "root")
var dbPass = getEnvWithDefault("DBPASS", "")
var dbPort = getEnvWithDefault("DBPORT", "3306")
var dbSalt = getEnvWithDefault("DBSALT", "1234") + "ksakaldfakljdsfkjladsfkjkjklajsdasdfa"

func init() {
	log.SetFlags(log.Lshortfile)

	log.Print("models.init")

	createDBIfNeeded()

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbName,
	)
	log.Printf("dsn:%s", dsn)
	var e error
	db, e = gorm.Open("mysql", dsn)
	if e != nil {
		log.Print("")
		log.Fatal(e)
	}
	log.Print("before migrateTables")
	migrateTables()
	log.Print("after migrateTables")
	isConnected = true
	return
}

// createDBIfNeeded ...
func createDBIfNeeded() error {
	log.Print("createDBIfNeeded")
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&collation=utf8mb4_unicode_ci&parseTime=True&loc=Local",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		"mysql",
	)
	for {
		dbtmp, e := gorm.Open("mysql", dsn)
		if e == nil {
			sql := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;",
				dbName)
			log.Printf("sql:%s", sql)
			dbtmp.Exec(sql)
			sql2 := fmt.Sprintf("use `%s`;", dbName)
			log.Printf("sql2:%s", sql2)
			dbtmp.Exec(sql2)
			dbtmp.Close()
			return nil
		}
		log.Print(e)
		log.Printf("dsn:%s", dsn)
		time.Sleep(time.Second * 5)

	}
}

func migrateTables() {
	log.Printf("User")
	db.AutoMigrate(&Comment{})

}
