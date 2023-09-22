package routers

import (
	"web/middleware/jwt"
	"web/routers/api"
	v1 "web/routers/api/v1"

	docs "web/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	docs.SwaggerInfo.BasePath = ""
	r.GET("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{

		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)

		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}
	apiget := r.Group("api/get")
	{
		//获取标签列表
		apiget.GET("/tags", v1.GetTags)
		//获取文章列表
		apiget.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiget.GET("/articles/:id", v1.GetArticle)
		//获取指定年份文章
		apiget.GET("/articles_by_year", v1.GetArticlesByYear)
	}
	return r
}
