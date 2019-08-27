package routers

import (
	"gin-restful-api/controllers/article"
	"gin-restful-api/controllers/auth"
	"gin-restful-api/controllers/tag"
	"gin-restful-api/controllers/upload"
	"gin-restful-api/middleware/cors"
	"gin-restful-api/middleware/jwt"
	"gin-restful-api/pkg/images"
	"gin-restful-api/pkg/setting"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.CORS())

	// 未知路由
	r.NoRoute(func(c *gin.Context) {
		c.String(404, "未找到路由地址")
	})
	// 未知调用方式
	r.HandleMethodNotAllowed = true
	r.NoMethod(func(c *gin.Context) {
		c.String(404, "错误调用方法")
	})

	r.POST("/auth", auth.GetAuth)
	r.POST("/upload", upload.UploadImage)
	r.StaticFS("/upload/images/"+time.Now().Format(setting.AppSetting.TimeFormat)+"/", http.Dir(images.GetImageFullPath()))

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", tag.GetTags)
		//新建标签
		apiv1.POST("/tags", tag.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", tag.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", tag.DeleteTag)

		//获取文章列表
		apiv1.GET("/articles", article.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", article.GetArticle)
		//新建文章
		apiv1.POST("/articles", article.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", article.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", article.DeleteArticle)
	}

	return r
}
