package article

import (
	"gin-restful-api/models"
	"gin-restful-api/pkg/app"
	"gin-restful-api/pkg/setting"
	"gin-restful-api/pkg/util"

	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
)

//获取单个文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	app := app.Gin{C: c}
	valid := validation.Validation{}
	if v := valid.Required(id, "id").Message("id不能为空"); !v.Ok {
		app.Response(400, v.Error.Message, nil)
		return
	}
	var data interface{}
	if models.ExistArticleByID(id) {
		data = models.GetArticle(id)
		app.Response(200, "ok", data)
		return
	}
	app.Response(404, "文章不存在", nil)
	return

}

//获取多个文章
func GetArticles(c *gin.Context) {
	data := make(map[string]interface{})
	maps := make(map[string]interface{})
	app := app.Gin{C: c}
	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
		if v := valid.Range(state, 0, 1, "state").Message("状态只允许0或1"); !v.Ok {
			app.Response(400, v.Error.Message, nil)
			return
		}
	}

	var tagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		maps["tag_id"] = tagId
		if v := valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0"); !v.Ok {
			app.Response(400, v.Error.Message, nil)
			return
		}
	}
	data["lists"] = models.GetArticles(util.GetPage(c), setting.AppSetting.PageSize, maps)
	data["total"] = models.GetArticleTotal(maps)

	app.Response(200, "ok", data)
	return
}

//新增文章
func AddArticle(c *gin.Context) {
	tagId := com.StrTo(c.PostForm("tag_id")).MustInt()
	title := c.PostForm("title")
	desc := c.PostForm("desc")
	content := c.PostForm("content")
	createdBy := c.PostForm("created_by")
	state := com.StrTo(c.DefaultPostForm("state", "0")).MustInt()
	imageUrl := c.PostForm("image_url")

	app := app.Gin{C: c}
	valid := validation.Validation{}
	if v := valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0"); !v.Ok {
		app.Response(400, v.Error.Message, nil)
		return
	}
	if v := valid.Required(title, "title").Message("标题不能为空"); !v.Ok {
		app.Response(400, v.Error.Message, nil)
		return
	}
	if v := valid.Required(desc, "desc").Message("描述不能为空"); !v.Ok {
		app.Response(400, v.Error.Message, nil)
		return
	}
	if v := valid.Required(content, "content").Message("内容不能为空"); !v.Ok {
		app.Response(400, v.Error.Message, nil)
		return
	}
	if v := valid.Required(createdBy, "created_by").Message("创建人不能为空"); !v.Ok {
		app.Response(400, v.Error.Message, nil)
		return
	}
	if v := valid.Range(state, 0, 1, "state").Message("状态只允许0或1"); !v.Ok {
		app.Response(400, v.Error.Message, nil)
		return
	}

	if models.ExistTagByID(tagId) {
		data := make(map[string]interface{})
		data["tag_id"] = tagId
		data["title"] = title
		data["desc"] = desc
		data["content"] = content
		data["created_by"] = createdBy
		data["state"] = state
		data["image_url"] = imageUrl

		models.AddArticle(data)
		app.Response(200, "ok", nil)
		return
	} else {
		app.Response(404, "标签不存在", nil)
		return
	}
}

//修改文章
func EditArticle(c *gin.Context) {
	app := app.Gin{C: c}
	valid := validation.Validation{}

	id := com.StrTo(c.Param("id")).MustInt()
	tagId := com.StrTo(c.PostForm("tag_id")).MustInt()
	title := c.PostForm("title")
	desc := c.PostForm("desc")
	content := c.PostForm("content")
	modifiedBy := c.PostForm("modified_by")
	imageUrl := c.PostForm("image_url")
	var state int = -1
	if arg := c.PostForm("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		if v := valid.Range(state, 0, 1, "state").Message("状态只允许0或1"); !v.Ok {
			app.Response(400, v.Error.Message, nil)
			return
		}

	}
	if v := valid.Min(id, 1, "id").Message("ID必须大于0"); !v.Ok {
		app.Response(400, v.Error.Message, nil)
		return
	}
	if v := valid.MaxSize(title, 100, "title").Message("标题最长为100字符"); !v.Ok {
		app.Response(400, v.Error.Message, nil)
		return
	}
	if v := valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符"); !v.Ok {
		app.Response(400, v.Error.Message, nil)
		return
	}
	if v := valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符"); !v.Ok {
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

	if models.ExistArticleByID(id) {
		if models.ExistTagByID(tagId) {
			data := make(map[string]interface{})
			if tagId > 0 {
				data["tag_id"] = tagId
			}
			if title != "" {
				data["title"] = title
			}
			if desc != "" {
				data["desc"] = desc
			}
			if content != "" {
				data["content"] = content
			}
			if imageUrl != "" {
				data["image_url"] = imageUrl
			}
			data["modified_by"] = modifiedBy

			models.EditArticle(id, data)
			app.Response(200, "ok", nil)
			return
		} else {
			app.Response(404, "标签不存在", nil)
			return
		}
	} else {
		app.Response(404, "文章不存在", nil)
		return
	}

}

//删除文章
func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	app := app.Gin{C: c}
	valid := validation.Validation{}
	if v := valid.Min(id, 1, "id").Message("ID必须大于0"); !v.Ok {
		app.Response(400, v.Error.Message, nil)
		return
	}

	if models.ExistArticleByID(id) {
		models.DeleteArticle(id)
		app.Response(200, "删除成功", nil)
		return
	} else {
		app.Response(404, "文章不存在", nil)
		return
	}
}
