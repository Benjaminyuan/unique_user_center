package dao

import (
	"context"
	"github.com/sirupsen/logrus"
)

func CreateUser(c context.Context,user *User)error{
	if	err :=  db.Table("user").Create(user).Error ; err != nil{
		logrus.Errorf("fail to create user,err: %v",err)
		return err
	}
	return nil
}
