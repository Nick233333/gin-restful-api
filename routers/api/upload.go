package api

import (
	"github.com/gin-gonic/gin"

	"gin-restful-api/pkg/app"
	"gin-restful-api/pkg/logging"
	"gin-restful-api/pkg/upload"
)

func UploadImage(c *gin.Context) {

	data := make(map[string]string)
	app := app.Gin{C: c}
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		app.Response(500, "err", nil)
		return
	}

	if image == nil {
		app.Response(400, "请求参数错误", nil)
		return
	}
	imageName := upload.GetImageName(image.Filename)
	fullPath := upload.GetImageFullPath()
	savePath := upload.GetImagePath()

	src := fullPath + imageName
	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		app.Response(400, "图片类型不允许上传", nil)
		return
	}

	if err := upload.CheckImage(fullPath); err != nil {
		app.Response(400, "校验图片错误，图片格式或大小有问题", nil)
		return
	}
	if err := c.SaveUploadedFile(image, src); err != nil {
		app.Response(500, "保存图片失败", nil)
		return
	}
	data["image_url"] = upload.GetImageFullUrl(imageName)
	data["image_save_url"] = savePath + imageName
	app.Response(200, "ok", data)
	return
}
