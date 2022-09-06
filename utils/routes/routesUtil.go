package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Option = func(engine *gin.Engine)

var Options = []Option{}

func include(opts ...Option) {
	Options = append(Options, opts...)
}

func InitRouters(serverPort string, runMode string) {
	include(UserRouters)
	r := gin.Default()
	r.Use(CrossDomain())
	for _, opt := range Options {
		opt(r)
	}
	gin.SetMode(runMode)
	err := r.Run(":" + serverPort)
	if err != nil {
		fmt.Println("InitRouters err=" + err.Error())
	}
}
