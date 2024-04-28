package rest

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()
	router.GET("/cars", getCarsHandler)
	router.POST("/cars", postCarsHandler)
	router.PUT("/cars/:id", putCarsHandler)
	router.DELETE("/cars/:id", deleteCarsHandler)
	return router
}

func StartServer(router *gin.Engine, port string) {
	router.Run(":" + port)
}
