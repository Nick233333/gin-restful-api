package tag

import (
	"gin-restful-api/models"
	"gin-restful-api/pkg/app"
	"gin-restful-api/pkg/setting"
	"gin-restful-api/pkg/util"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

//获取多个文章标签
func GetTags(c *gin.Context) {
	name := c.Query("name")

	maps := make(map[string]interface{})
	data := make(map[string]interface{})
	app := app.Gin{C: c}
	valid := validation.Validation{}
	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
		if v := valid.Range(state, 0, 1, "state").Message("状态只允许0或1"); !v.Ok {
			app.Response(400, v.Error.Message, nil)
			return
		}
	}

	data["lists"] = models.GetTags(util.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"] = models.GetTagTotal(maps)

	app.Response(200, "ok", data)
	return
}

//新增文章标签
func AddTag(c *gin.Context) {
	name := c.PostForm("name")
	state := com.StrTo(c.DefaultPostForm("state", "0")).MustInt()
	createdBy := c.PostForm("created_by")

	app := app.Gin{C: c}
	valid := validation.Validation{}
	if v := valid.Required(name, "name").Message("名称不能为空"); !v.Ok {
		app.Response(400, v.Error.Message, nil)
		return
	}
	if v := valid.MaxSize(name, 100, "name").Message("名称最长为100字符"); !v.Ok {
		app.Response(400, v.Error.Message, nil)
		return
	}
	if v := valid.Required(createdBy, "created_by").Message("创建人不能为空"); !v.Ok {
		app.Response(400, v.Error.Message, nil)
		return
	}
	if v := valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符"); !v.Ok {
		app.Response(400, v.Error.Message, nil)
		return
	}
	if v := valid.Range(state, 0, 1, "state").Message("状态只允许0或1"); !v.Ok {
		app.Response(400, v.Error.Message, nil)
		return
	}

	if !models.ExistTagByName(name) {

		models.AddTag(name, state, createdBy)
		app.Response(200, "ok", nil)
		return
	} else {
		app.Response(400, "标签已存在", nil)
		return
	}
}

//修改文章标签
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.PostForm("name")
	modifiedBy := c.PostForm("modified_by")

	app := app.Gin{C: c}
	valid := validation.Validation{}

	var state int = -1
	if arg := c.PostForm("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		if v := valid.Range(state, 0, 1, "state").Message("状态只允许0或1"); !v.Ok {
			app.Response(400, v.Error.Message, nil)
			return
		}

	}
	if v := valid.Required(id, "id").Message("ID不能为空"); !v.Ok {
		app.Response(400, v.Error.Message, nil)
		return
	}
	if v := valid.Required(modifiedBy, "modified_by").Message("修改人不能为空"); !v.Ok {
		app.Response(400, v.Error.Message, nil)
		return
	}
	if v := valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符"); !v.Ok {
		app.Response(400, v.Error.Message, nil)
		return
	}
	if v := valid.MaxSize(name, 100, "name").Message("名称最长为100字符"); !v.Ok {
		app.Response(400, v.Error.Message, nil)
		return
	}

	if models.ExistTagByID(id) {
		data := make(map[string]interface{})
		data["modified_by"] = modifiedBy
		if name != "" {
			data["name"] = name
		}
		if state != -1 {
			data["state"] = state
		}

		models.EditTag(id, data)
		app.Response(200, "ok", nil)
		return
	} else {
		app.Response(404, "标签不存在", nil)
		return
	}
}

//删除文章标签
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	app := app.Gin{C: c}
	valid := validation.Validation{}
	if v := valid.Min(id, 1, "id").Message("ID必须大于0"); !v.Ok {
		app.Response(400, v.Error.Message, nil)
		return
	}

	if models.ExistTagByID(id) {
		models.DeleteTag(id)
		app.Response(200, "ok", nil)
		return
	} else {
		app.Response(404, "标签不存在", nil)
		return
	}

}
