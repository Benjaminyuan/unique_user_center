package dao

import (
	"context"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func CreateUser(c context.Context,user *User)error{
	if	err :=  Db.Table("user").Save(user).Error ; err != nil{
		logrus.Errorf("fail to create user,err: %v",err)
		return err
	}
	return nil
}
func GetUserInfoByName(c context.Context,name string)(*User,error){
	user := &User{}
	if err := Db.Table("user").Where("name = ?", name).Find(user).Error; err != nil{
		if err == gorm.ErrRecordNotFound{
			logrus.Errorf("user name %v not found,err: %v",name,err)
			return nil,err
		}
		logrus.Errorf("user  by name failed,err: %v",err)
		return nil,err
	}
	return user,nil
}

func GetUserInfoById(c context.Context, uid int64)(*User,error){
	user := User{}
	if err := Db.Table("user").Where("id = ?", uid).Find(&user).Error; err != nil{
		if err == gorm.ErrRecordNotFound{
			logrus.Errorf("user id %v not found,err: %v",err)
			return nil,err
		}
		logrus.Errorf("user  by id failed,err: %v",err)
		return nil,err
	}
	return &user,nil
}