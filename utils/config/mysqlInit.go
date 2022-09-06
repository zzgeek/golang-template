package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var MysqlPool *sql.DB

/*
安装mysql
mysqld.exe --install
初始化mysql
mysqld.exe --initialize
启动mysql
net start mysql
进入mysql
mysql -u root -p
设置MySQL密码
set password=zzgeek;
alter user 'root'@'localhost' password expire never;
flush privileges;
重启MySQL服务
net stop mysql
net start mysql

ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'zzgeek';
*/

func InitMysqlPool(dbType, host, port, userName, password, dbname string, maxIdleConn, maxOpenConn int, connMaxLifetime time.Duration) *sql.DB {
	MysqlPool, err := sql.Open(dbType, userName+":"+password+"@tcp("+host+":"+port+")/"+dbname)
	if err != nil {
		fmt.Println("InitMysqlPool:" + err.Error())
	}
	MysqlPool.SetMaxIdleConns(maxIdleConn)
	MysqlPool.SetMaxOpenConns(maxOpenConn)
	MysqlPool.SetConnMaxLifetime(connMaxLifetime)
	return MysqlPool
}
