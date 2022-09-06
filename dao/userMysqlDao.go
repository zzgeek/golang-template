package dao

import (
	"database/sql"
	"fmt"
	"webapp01/utils/config"
)

var MyUserMysqlDao *userMysqlDao

type userMysqlDao struct {
	mysqlPool *sql.DB
}

func NewUserMysqlDao(mysqlPool *sql.DB) *userMysqlDao {
	MyUserMysqlDao = &userMysqlDao{
		mysqlPool: mysqlPool,
	}
	return MyUserMysqlDao
}

func (this *userMysqlDao) GetMysqlPool() *sql.DB {
	return this.mysqlPool
}

func (this *userMysqlDao) AddUser1(userName, userPwd string, userStatus int) error {
	sqlStr := "insert into users(userName,userPwd,userStatus) values(?,?,?)"
	inStmt, err := this.mysqlPool.Prepare(sqlStr)
	if err != nil {
		fmt.Println("预编译:" + err.Error())
		return err
	}
	_, err = inStmt.Exec(userName, userPwd, userStatus)
	if err != nil {
		fmt.Println("执行:" + err.Error())
		return err
	}
	return nil
}

func (this *userMysqlDao) AddUser2(userName, userPwd string, userStatus int) error {
	sqlStr := "insert into users(userName,userPwd,userStatus) values(?,?,?)"
	_, err := config.MysqlPool.Exec(sqlStr, userName, userPwd, userStatus)
	if err != nil {
		fmt.Println("AddUser2执行:" + err.Error())
		return err
	}
	return nil
}
