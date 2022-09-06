package dao

import "github.com/gomodule/redigo/redis"

var (
	MyUserDao *userDao
)

type userDao struct {
	pool *redis.Pool
}

func NewUserDao(redisPool *redis.Pool) *userDao {
	MyUserDao = &userDao{
		pool: redisPool,
	}
	return MyUserDao
}

func (this *userDao) GetPool() *redis.Pool {
	return this.pool
}
