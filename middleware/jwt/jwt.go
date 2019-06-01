package jwt

import (
	"gin-restful-api/pkg/app"
	"time"

	"github.com/gin-gonic/gin"

	"gin-restful-api/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int = 200
		var msg string = ""
		var data interface{}
		app := app.Gin{C: c}
		token := ""
		token = c.Request.Header.Get("Token")
		if token == "" {
			code = 400
			msg = "参数错误"
		}
		claims, err := util.ParseToken(token)
		if err != nil {
			code = 500
			msg = "验证失败"
		} else if time.Now().Unix() > claims.ExpiresAt {
			code = 401
			msg = "token已过期"
		}

		if code != 200 {
			app.Response(code, msg, data)
			c.Abort()
			return
		}

		c.Next()
	}
}
