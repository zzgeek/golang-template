package control

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"webapp01/model"
	"webapp01/utils/jwt"
	"webapp01/utils/qiniu"
)

func AddUserHandler(context *gin.Context) {
	context.JSON(200, gin.H{"message": "GET"})
}

func AuthHandler(context *gin.Context) {
	x, _ := strconv.ParseFloat(context.Param("x"), 64)
	y, _ := strconv.ParseFloat(context.Param("y"), 64)
	context.String(200, fmt.Sprintf("%f", x+y))
}

func GetQiniuTokenHandler(context *gin.Context) {
	token := qiniu.GetQiniuToken()
	responseView := model.ResponseView{RespCode: 200, RespMsg: "请求成功", RespData: token}
	context.JSON(200, responseView)

}

func GetUserHandler(context *gin.Context) {
	user := model.User{UserName: "admin", UserPwd: "123", UserId: 1, UserStatus: 0, UserRole: "admin"}
	userToken, _ := jwt.GenToken(user.UserName)
	fmt.Println(userToken)
	user.UserToken = userToken
	responseView := model.ResponseView{RespCode: 200, RespMsg: "请求成功", RespData: user}
	context.JSON(200, responseView)
}
