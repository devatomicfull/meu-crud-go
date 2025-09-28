package routes

import "github.com/gin-gonic/gin"

func InitRoutes(rg *gin.RouterGroup) {
	rg.GET("/getUserById/:userId")
	rg.GET("/getUserByEmail/:email")
	rg.POST("/createUser/:userId")
	rg.PUT("/updateUser/:userId")
	rg.DELETE("/deleteUser/:userId")
}
