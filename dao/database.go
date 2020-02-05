package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"

	"user_center/conf"
)
var (
	Db *gorm.DB
)
func InitDB()error{
	var err error
	db, err  := gorm.Open("mysql",conf.Data.MysqlConf.MysqlAddr)
	if err != nil{
		logrus.Fatalf("fail to connect db,err: %v",err)
		return err
	}
	Db = db
	return nil
}