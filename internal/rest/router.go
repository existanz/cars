package rest

import (
	"cars/internal/database"

	"github.com/gin-gonic/gin"
)

type router struct {
	engine *gin.Engine
	db     database.CardyB
}

func NewRouter(db database.CardyB) router {
	engine := gin.Default()
	return router{engine: engine, db: db}
}

func StartServer(router router, port string) {
	router.engine.GET("/cars", router.getCarsHandler)
	router.engine.GET("/cars/:id", router.getCarByIdHandler)
	router.engine.POST("/cars", postCarsHandler)
	router.engine.PUT("/cars/:id", putCarsHandler)
	router.engine.DELETE("/cars/:id", router.deleteCarHandler)

	router.engine.Run(":" + port)
}
