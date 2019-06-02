package upload

import (
	"github.com/gin-gonic/gin"

	"gin-restful-api/pkg/app"
	"gin-restful-api/pkg/images"
)

func UploadImage(c *gin.Context) {

	data := make(map[string]string)
	app := app.Gin{C: c}
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		app.Response(400, "err", nil)
		return
	}

	if image == nil {
		app.Response(400, "请求参数错误", nil)
		return
	}
	imageName := images.GetImageName(image.Filename)
	fullPath := images.GetImageFullPath()
	savePath := images.GetImagePath()

	src := fullPath + imageName
	if !images.CheckImageExt(imageName) || !images.CheckImageSize(file) {
		app.Response(400, "图片类型不允许上传", nil)
		return
	}

	if err := images.CheckImage(fullPath); err != nil {
		app.Response(400, "校验图片错误，图片格式或大小有问题", nil)
		return
	}
	if err := c.SaveUploadedFile(image, src); err != nil {
		app.Response(500, "保存图片失败", nil)
		return
	}
	data["image_url"] = images.GetImageFullUrl(imageName)
	data["image_save_url"] = savePath + imageName
	app.Response(200, "ok", data)
	return
}
