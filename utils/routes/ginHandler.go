package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"webapp01/model"
	"webapp01/utils/jwt"
)

func AuthHeader() gin.HandlerFunc {
	return func(context *gin.Context) {
		userToken := context.Request.Header.Values("userToken")
		_, err := jwt.ParseToken(userToken[0])
		if err != nil {
			fmt.Println(err)
			context.Abort()
			responseView := model.ResponseView{RespCode: http.StatusUnauthorized, RespMsg: err.Error(), RespData: nil}
			context.JSON(http.StatusUnauthorized, responseView)
		} else {
			context.Next()
		}
	}
}

func CrossDomain() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		context.Header("Access-Control-Allow-Origin", "*") // 可将将 * 替换为指定的域名
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		context.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		context.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			context.AbortWithStatus(http.StatusNoContent)
		}
		context.Next()
	}
}
