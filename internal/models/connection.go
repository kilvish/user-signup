package models

import (
	"errors"
	"log"

	config "github.com/kilvish/user-signup/internal/configmanager"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Mysqldb struct {
	db *gorm.DB
}

//MysqlConnString string
var MysqlConnString string

//MysqlDBConn conn
var MysqlDBConn *Mysqldb

//InitMysqlConnection initialie mysql connection
func InitMysqlConnection(mysqlConfig config.MysqlConnection) error {
	MysqlConnString = mysqlConfig.User + ":" + mysqlConfig.Pass + "@tcp(" + mysqlConfig.Host + ")/" + mysqlConfig.DBName + "?charset=utf8&parseTime=True&loc=Local"
	log.Println("Mysql connection string", MysqlConnString)
	DB, err := gorm.Open("mysql", MysqlConnString)
	mysql := new(Mysqldb)
	mysql.db = DB
	if err != nil {
		log.Fatal("Connection to DB failed", err)
		return errors.New("connection Failed")
	}
	MysqlDBConn = mysql
	return nil
}

// GetDBConnection returns db connection
func GetDBConnection() *Mysqldb {
	return MysqlDBConn
}
