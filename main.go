package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
)

func main() {

	router := gin.Default()
	// 注册路由和处理函数
	router.Any("/", WebRoot)

	router.GET("/get", getMethod)
	router.POST("/post", postethod)
	router.PUT("/put", putMethod)
	router.DELETE("/delete", deleteMethod)
	router.PATCH("/patch", patchMethod)
	router.OPTIONS("/options", optinsMethod)

	v1 := router.Group("/v1")
	{	
		// /v1/test
		v1.GET("/test", testMethod)
	}

	v2 := router.Group("/v2")
	// /v2/demo
	v2.GET("/demo", demoMethod)

	router.GET("/middleware", middleware1, middleware2, handler)


	// url为 /welcome?name=Jane&age=18
	router.GET("/welcome", func(c *gin.Context) {
		// 获取参数内容
		// 获取的所有参数内容的类型都是 string
		// 如果不存在，使用第二个当做默认内容
		name := c.DefaultQuery("name", "Guest")
		// 获取参数内容，没有则返回空字符串
		age := c.Query("age") 

		c.String(http.StatusOK, "Hello %s %s", name, age)
	})

	router.POST("/post-form", func(c *gin.Context) {
		// 获取post过来的message内容
		// 获取的所有参数内容的类型都是 string
		message := c.DefaultPostForm("message", "default")

		c.JSON(200, gin.H{
			"message": message,
		})
	})

	router.PUT("/put-form", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"message": message,
			"nick": nick,
		})
	})

	router.PATCH("/patch-form", func(c *gin.Context) {
		number := c.DefaultPostForm("number", "100")

		c.JSON(200, gin.H{
			"number": number,
		})
	})


	router.DELETE("/delete-form", func(c *gin.Context) {
		id := c.DefaultPostForm("id", "0")

		c.JSON(200, gin.H{
			"id": id,
		})
	})
		
	// url 匹配 /user/nick ， 但是它不会匹配 /user
    router.GET("/user/:name", func(c *gin.Context) {
        name := c.Param("name")
        c.String(http.StatusOK, "Hello %s", name)
    })

    // 可以匹配 /user/nick 和 /user/nick/add
    // 如果没有其他的路由匹配 /user/nick ， 它将重定向到 /user/john/
    router.GET("/user/:name/*action", func(c *gin.Context) {
        name := c.Param("name")
        action := c.Param("action")
        message := name + " is " + action
        c.String(http.StatusOK, message)
    })


	// 默认绑定 :8080 
	// 必须双引号
	router.Run(":8081")
}

/**
* 根请求处理函数
* 所有本次请求相关的方法都在 context 中，完美
* 输出响应 hello, world
*/
func WebRoot(context *gin.Context) {
	context.String(http.StatusOK, "hello, world")
}

// 字符串必须双引号
func getMethod(context *gin.Context) {
	// name := context.Param("name")
	// action := context.Param("action")
	// message := name + " is " + action
	// context.String(http.StatusOK, message)
	context.String(http.StatusOK, `getMethod`)
}

func postethod(context *gin.Context) {
	context.String(http.StatusOK, `postethod`)
}

func putMethod(context *gin.Context) {
	context.String(http.StatusOK, `putMethod`)
}

func deleteMethod(context *gin.Context) {
	context.String(http.StatusOK, `deleteMethod`)
}

func patchMethod(context *gin.Context) {
	context.String(http.StatusOK, `patchMethod`)
}

func optinsMethod(context *gin.Context) {
	context.String(http.StatusOK, `optinsMethod`)
}

func testMethod(context *gin.Context) {
	context.String(http.StatusOK, `test`)
}

func demoMethod(context *gin.Context) {
	context.String(http.StatusOK, `demo`)
}

func middleware1(c *gin.Context) {
	log.Println("run middleware1")
  
	//逻辑代码
  
	// 执行该中间件之后的逻辑
	c.Next()
}

func middleware2(c *gin.Context) {
	log.Println("arrive at middleware2")
	// 执行该中间件之前，先跳到流程的下一个方法
	c.Next()
	// 流程中的其他逻辑已经执行完了
	log.Println("run middleware2")
  
	//逻辑代码
}

func handler(c *gin.Context) {
	log.Println("run handler")
	c.String(http.StatusOK, "middleware")
}




