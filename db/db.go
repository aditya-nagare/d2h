package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

const (
	DBUSER     = "root"
	DBPASSWORD = "password"
	DBNAME     = "d2h"
	TESTDBNAME = "d2h-test"
)

//NewDbConnection for client connection
func NewDbConnection() *gorm.DB {

	strConnection := DBUSER + ":" + DBPASSWORD + "@/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"

	//Open Connection for GORM instance
	client, err := gorm.Open("mysql", strConnection)

	if err != nil {
		panic("Error! Can't create connection to database")
	}
	return client

}

//NewTestDbConnection for client connection
func NewTestDbConnection() *gorm.DB {

	strConnection := DBUSER + ":" + DBPASSWORD + "@/" + TESTDBNAME + "?charset=utf8&parseTime=True&loc=Local"

	//Open Connection for GORM instance
	client, err := gorm.Open("mysql", strConnection)

	if err != nil {
		panic("Error! Can't create connection to database")
	}
	return client

}
