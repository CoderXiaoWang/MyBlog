package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/go-programming-tour-book/blog-service/internal/routers/api/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	article := v1.NewArticle()
	tag := v1.NewTag()

	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/tags", tag.Create)       //增加标签
		apiv1.DELETE("/tags/:id", tag.Delete) //删除标签
		apiv1.PUT("/tags/:id", tag.Update)    //更改标签
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List) //查询标签

		apiv1.POST("/articles", article.Create)       //添加文章
		apiv1.DELETE("/articles/:id", article.Delete) //删除文章
		apiv1.PUT("/articles/:id", article.Update)    //更改文章
		apiv1.PATCH("/articles/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get) //查询文章
		apiv1.GET("/articles", article.List)    //查询文章列表
	}

	return r
}
