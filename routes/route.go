package routes

import (
	v1 "ginblog/api/v1"
	"ginblog/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	rV1 := r.Group("/api/v1")
	{
		// 用户模块路由接口
		rV1.POST("/user/add", v1.AddUser)
		rV1.GET("/users", v1.GetUsers)
		rV1.PUT("/user/:id", v1.EditUser)
		rV1.DELETE("/user/:id", v1.DeleteUser)
		// 分类模块路由接口
		rV1.POST("/category/add", v1.AddCategory)
		rV1.GET("/category", v1.GetCate)
		rV1.PUT("/category/:id", v1.EditCate)
		rV1.DELETE("/category/:id", v1.DeleteCate)
		// 文章模块路由接口
		rV1.POST("/article/add", v1.AddArticle)
		rV1.GET("/article", v1.GetArt)
		rV1.GET("/article/list/:id", v1.GetCateArt)
		rV1.GET("/article/info/:id", v1.GetArtInfo)
		rV1.PUT("/article/:id", v1.EditArt)
		rV1.DELETE("/article/:id", v1.DeleteArt)
	}
	r.Run(utils.HttpPort)
}
