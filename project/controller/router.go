package controller

import (
	"github.com/gin-gonic/gin"
	"sample_api/framework/logger"
	"sample_api/framework/setting"
	"sample_api/framework/auth"
	"sample_api/project/controller/api"
	"sample_api/project/controller/api/v1"
)


func InitRouter() *gin.Engine {
	r := gin.New()
	//r.Use(gin.Logger())
	//r.Use(gin.Recovery())
	r.Use(logger.GinLogger(logger.Logger), logger.GinRecovery(logger.Logger, true))
	gin.SetMode(setting.RunMode)


	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	apiv1.Use(auth.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)


		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}


	return r
}