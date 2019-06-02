package auth

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"gin-restful-api/models"
	"gin-restful-api/pkg/app"
	"gin-restful-api/pkg/logging"
	"gin-restful-api/pkg/util"
)

type auth struct {
	Username string
	Password string
}

func GetAuth(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	data := make(map[string]interface{})
	app := app.Gin{C: c}
	valid := validation.Validation{}
	a := auth{Username: username, Password: password}

	if v := valid.Required(a.Username, "username").Message("用户名不能为空"); !v.Ok {
		app.Response(400, v.Error.Message, nil)
		return
	}

	if v := valid.MaxSize(a.Username, 50, "username").Message("长度不能超过50"); !v.Ok {
		logging.Info(v.Error.Key, v.Error.Message)
		app.Response(400, v.Error.Message, nil)
		return
	}
	if v := valid.Required(a.Password, "password").Message("密码不能为空"); !v.Ok {
		app.Response(400, v.Error.Message, nil)
		return
	}
	if v := valid.MaxSize(a.Password, 50, "password").Message("长度不能超过50"); !v.Ok {
		app.Response(400, v.Error.Message, nil)
		return
	}
	isExist, err := models.CheckAuth(username, password)

	if err != nil {
		app.Response(400, "Token鉴权失败", nil)
		return
	}
	if !isExist {
		app.Response(400, "token验证失败", nil)
		return
	}
	token, err := util.GenerateToken(username, password)
	if err != nil {
		app.Response(500, "生成token失败", nil)
		return
	}
	data["token"] = token
	app.Response(200, "ok", data)
	return
}
