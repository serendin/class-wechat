package models

import (
	"strings"
	"fmt"
	"github.com/revel/revel"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	DB *gorm.DB
)

func getConnectionString() string {
	host := getParamString("db.host", "localhost")
	port := getParamString("db.port", "3306")
	user := getParamString("db.user", "root")
	pass := getParamString("db.password", "123456")
	name := getParamString("db.name", "wechat_class")
	protocol := getParamString("db.protocol", "tcp")
	dbargs := getParamString("db.args", "charset=utf8&parseTime=True&loc=Local")
	if strings.Trim(dbargs, " ") != "" {
		dbargs = "?" + dbargs
	} else {
		dbargs = ""
	}
	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s",
		user, pass, protocol, host, port, name, dbargs)
}
func getParamString(param string, defaultValue string) string {
	p, found := revel.Config.String(param)
	if !found {
		if defaultValue == "" {
			revel.ERROR.Fatal("Could not find parameter: " + param)
		} else {
			return defaultValue
		}
	}
	return p
}


func InitDB(){
	log.Println("initDB==========")
	driver:=getParamString("db.driver","mysql")
	connectionString := getConnectionString()
	var err error
	log.Println(driver,connectionString)
	DB, err = gorm.Open(driver,connectionString)
	if err != nil {
		revel.WARN.Printf("DB错误: %v", err)
	}
}

