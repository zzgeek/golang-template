package routes

import (
	"github.com/gin-gonic/gin"
	"webapp01/control"
)

func UserRouters(engine *gin.Engine) {
	engine.GET("/getUser", control.GetUserHandler)
	engine.GET("/getQiniuToken", AuthHeader(), control.GetQiniuTokenHandler)
	engine.GET("/Adduser", control.AddUserHandler)
	engine.GET("/Auth/:x/:y", control.AuthHandler)

}
