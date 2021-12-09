package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/middleware"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	auth := r.Group("/api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块路由接口
		auth.PUT("/user/:id", v1.EditUser)
		auth.DELETE("/user/:id", v1.DeleteUser)
		// 分类模块路由接口
		auth.POST("/category/add", v1.AddCategory)
		auth.PUT("/category/:id", v1.EditCate)
		auth.DELETE("/category/:id", v1.DeleteCate)
		// 文章模块路由接口
		auth.POST("/article/add", v1.AddArticle)
		auth.PUT("/article/:id", v1.EditArt)
		auth.DELETE("/article/:id", v1.DeleteArt)
	}

	rV1 := r.Group("/api/v1")
	{
		// 用户模块路由接口
		rV1.GET("/users", v1.GetUsers)
		rV1.POST("/user/add", v1.AddUser)
		// 分类模块路由接口
		rV1.GET("/category", v1.GetCate)
		// 文章模块路由接口
		rV1.GET("/article", v1.GetArt)
		rV1.GET("/article/list/:id", v1.GetCateArt)
		rV1.GET("/article/info/:id", v1.GetArtInfo)

		rV1.POST("/login", v1.Login)
	}
	_ = r.Run(utils.HttpPort)
}
