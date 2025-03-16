package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kriswu/go_deepseek/internal/app/handler"
	"github.com/kriswu/go_deepseek/internal/app/middleware"
)

// SetupRouter 配置路由
func SetupRouter(userHandler *handler.UserHandler) *gin.Engine {
	r := gin.Default()

	// 添加全局中间件
	r.Use(middleware.Cors())
	r.Use(middleware.Logger())

	// API版本分组
	v1 := r.Group("/api/v1")
	{
		// 用户相关路由
		userGroup := v1.Group("/users")
		{
			userGroup.POST("/register", userHandler.Register)
			userGroup.POST("/login", userHandler.Login)
			userGroup.GET("/:id", userHandler.GetUser)
			userGroup.PUT("/:id", userHandler.UpdateUser)
			userGroup.DELETE("/:id", userHandler.DeleteUser)
			userGroup.GET("", userHandler.ListUsers)
		}
	}

	return r
}