package middleware

import (
	"github.com/gin-gonic/gin"
	"qiandao/pkg/app"
	"qiandao/pkg/token"
)

// Auth 中间件 即登陆过后每个接口都需要验证 token 信息
func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		if _, err := token.ParseRequest(context); err != nil {
			app.SendResponse(context, err, nil)
			context.Abort()
			return
		}
		context.Next()
	}
}
