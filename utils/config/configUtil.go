package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path"
	"runtime"
	"time"
)

var Server *server
var Database *database
var Redis *myRedis
var Jwt *jwt
var Qiniu *qiniu

type conf struct {
	Svc         server   `yaml:"server"`
	DB          database `yaml:"database"`
	RedisConfig myRedis  `yaml:"redis"`
	Jwt         jwt      `yaml:"jwt"`
	Qiniu       qiniu    `yaml:"qiniu"`
}

type server struct {
	Port     string `yaml:"port"`
	RunMode  string `yaml:"runMode"`
	LogLevel string `yaml:"logLevel"`
}

type database struct {
	Type            string `yaml:"type"`
	Host            string `yaml:"host"`
	Port            string `yaml:"port"`
	UserName        string `yaml:"username"`
	Password        string `yaml:"password"`
	DbName          string `yaml:"dbname"`
	MaxIdleConn     int    `yaml:"maxIdleConn"`
	MaxOpenConn     int    `yaml:"maxOpenConn"`
	ConnMaxLifetime int    `yaml:"connMaxLifetime"`
}

type myRedis struct {
	Host        string        `yaml:"host"`
	Port        string        `yaml:"port"`
	Password    string        `yaml:"password"`
	DB          int           `yaml:"db"`
	MaxIdle     int           `yaml:"maxIdle"`
	MaxActive   int           `yaml:"maxActive"`
	IdleTimeOut time.Duration `yaml:"idleTimeOut"`
}

type jwt struct {
	Seed         string        `yaml:"seed"`
	TokenExpired time.Duration `yaml:"tokenExpired"`
	Issuer       string        `yaml:"issuer"`
}

type qiniu struct {
	AccessKey string `yaml:"accessKey"`
	SecretKey string `yaml:"secretKey"`
	Bucket    string `yaml:"bucket"`
}

func (c conf) String() string {
	return fmt.Sprintf("%v\n%v\n%v\n%v\n%v", c.Svc, c.DB, c.RedisConfig, c.Jwt, c.Qiniu)
}

func (s server) String() string {
	return fmt.Sprintf("server : \n"+
		"\tport : %v \n"+
		"\tRunMode : %v", s.Port, s.RunMode)
}

func (m database) String() string {
	return fmt.Sprintf("database : \n"+
		"\ttype : %v \n"+
		"\thost : %v \n"+
		"\tport : %v \n"+
		"\tusername : %v \n"+
		"\tpassword : %v \n"+
		"\tdbname : %v \n"+
		"\tmaxIdleConn : %v \n"+
		"\tmaxOpenConn : %v \n"+
		"\tconnMaxLifetime : %v",
		m.Type, m.Host, m.Port, m.UserName, m.Password, m.DbName, m.MaxOpenConn, m.MaxIdleConn, m.ConnMaxLifetime)
}
func (r myRedis) String() string {
	return fmt.Sprintf("redis : \n"+
		"\thost : %v \n"+
		"\tport : %v \n"+
		"\tPassword : %v \n"+
		"\tMaxIdle : %v \n"+
		"\tMaxActive : %v \n"+
		"\tIdleTimeOut : %v \n"+
		"\tdb : %v",
		r.Host, r.Port, r.Password, r.MaxIdle, r.MaxActive, r.IdleTimeOut, r.DB)
}

func (j jwt) String() string {
	return fmt.Sprintf("jwt : \n"+
		"\tseed : %v \n"+
		"\ttokenExpired : %v \n"+
		"\tissuer : %v", j.Seed, j.TokenExpired, j.Issuer)
}

func (q qiniu) String() string {
	return fmt.Sprintf("qiniu : \n"+
		"\taccessKey : %v \n"+
		"\tsecretKey : %v \n"+
		"\tbucket : %v \n", q.AccessKey, q.SecretKey, q.Bucket)
}

func InitConf(dataFile string) {
	// 解决相对路经下获取不了配置文件问题
	_, filename, _, _ := runtime.Caller(0)
	filePath := path.Join(path.Dir(filename), dataFile)
	_, err := os.Stat(filePath)
	if err != nil {
		log.Printf("config file path %s not exist", filePath)
	}
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	c := new(conf)
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Printf("Unmarshal: %v", err)
	}
	fmt.Printf("load conf success\n %v", c)
	// 绑定到外部可以访问的变量中
	Server = &c.Svc
	Database = &c.DB
	Redis = &c.RedisConfig
	Jwt = &c.Jwt
	Qiniu = &c.Qiniu
}
