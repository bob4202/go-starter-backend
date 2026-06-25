package user

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup, handler *Handler, authMiddleware *gin.HandlerFunc) {
	users := router.Group("/users")

	users.POST("/register", handler.Register)
	users.POST("/login", handler.Login)

	users.GET("/email", handler.GetUserByEmail)
	users.GET("/:id", handler.GetUserById)

	protected := users.Group("")
	protected.GET("/me", handler.Me)
	protected.GET("/me/password", handler.ChangePassword)
}
