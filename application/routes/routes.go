package routes

import (
    "github.com/gin-gonic/gin"
    "sandbox/application/controllers"
)

func SetupRouter() *gin.Engine {
    r := gin.Default()

    user := r.Group("/users")
    {
        user.GET("/", controllers.GetUsers)
        user.POST("/", controllers.CreateUser)
        user.GET("/:id", controllers.GetUser)
        user.PUT("/:id", controllers.UpdateUser)
        user.DELETE("/:id", controllers.DeleteUser)
    }

    return r
}
