package main

import (
	"fmt"
	"time"
	"webapp01/dao"
	"webapp01/utils/config"
	"webapp01/utils/routes"
)

func init() {
	// 加载配置文件
	config.InitConf("../../config/aplication-test.yaml")
	//config.InitConf("../../config/aplication-release.yaml")
	fmt.Println("\n配置文件加载完成...")
	//加载redis连接池
	config.RedisPool = config.InitRedisPool(config.Redis.Host+":"+config.Redis.Port, config.Redis.MaxIdle,
		config.Redis.MaxActive, config.Redis.IdleTimeOut)
	fmt.Println("redis连接池加载完成...")
	//加载mysql连接池
	config.MysqlPool = config.InitMysqlPool(config.Database.Type, config.Database.Host, config.Database.Port, config.Database.UserName, config.Database.Password, config.Database.DbName, config.Database.MaxIdleConn, config.Database.MaxOpenConn, time.Duration(config.Database.ConnMaxLifetime))
	fmt.Println("mysql连接池加载完成...")
	dao.MyUserDao = dao.NewUserDao(config.RedisPool)
	fmt.Println("MyuserDao 初始化完成=", dao.MyUserDao)
	dao.MyUserMysqlDao = dao.NewUserMysqlDao(config.MysqlPool)
	fmt.Println("MyUserMysqlDao 初始化完成=", dao.MyUserDao)
	//初始化jwt
	config.JwtConf = config.InitJwtConfig(config.Jwt.Seed, config.Jwt.Issuer, config.Jwt.TokenExpired*time.Hour)
	fmt.Println("InitJwtConfig 初始化jwt完成")
	//初始化gin
	routes.InitRouters(config.Server.Port, config.Server.RunMode)
	fmt.Println("InitRouters 初始化路由完成...")

}
