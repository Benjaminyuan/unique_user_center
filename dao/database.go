package dao

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"user_center/conf"
)
var (
	db *gorm.DB
)
func InitDB()error{
	var err error
	db, err  = gorm.Open("mysql",conf.Data.MysqlConf.MysqlAddr)
	if err != nil{
		logrus.Fatalf("fail to connect db,err: %v",err)
		return err
	}
	return nil
}