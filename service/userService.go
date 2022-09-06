package service

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"webapp01/dao"
	"webapp01/model"
)

func GetUserByUserId(userId int) (user *model.User, err error) {
	redisConn := dao.MyUserDao.GetPool().Get()
	defer redisConn.Close()
	res, err := redis.String(redisConn.Do("HGet", "users", userId))
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println(err)
	}
	return
}
