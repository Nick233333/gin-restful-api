package jwt

import (
	"gin-restful-api/pkg/app"
	"time"

	"github.com/gin-gonic/gin"

	"gin-restful-api/pkg/util"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {

		app := app.Gin{C: c}
		token := ""
		token = c.Request.Header.Get("Token")
		if token == "" {
			app.Response(400, "参数错误token为空", nil)
			c.Abort()
			return
		}
		claims, err := util.ParseToken(token)
		if err != nil {
			app.Response(401, "token验证失败", nil)
			c.Abort()
			return
		}
		if time.Now().Unix() > claims.ExpiresAt {
			app.Response(401, "token已过期", nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
