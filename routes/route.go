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

		// 文章模块路由接口
	}
	r.Run(utils.HttpPort)
}
