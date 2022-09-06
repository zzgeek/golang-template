package service

import (
	"fmt"
	"webapp01/dao"
)

func AddUserService(userName, userPwd string, userStatus int) error {
	err := dao.MyUserMysqlDao.AddUser2(userName, userPwd, userStatus)
	if err != nil {
		fmt.Println("AddUserService:" + err.Error())
		return err
	}
	return nil
}
