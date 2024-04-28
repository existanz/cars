package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getCarsHandler(c *gin.Context) {
	c.String(http.StatusOK, "Get cars")
}

func postCarsHandler(c *gin.Context) {
	c.String(http.StatusOK, "Post cars")
}

func putCarsHandler(c *gin.Context) {
	c.String(http.StatusOK, "Put cars")
}

func deleteCarsHandler(c *gin.Context) {
	c.String(http.StatusOK, "Delete cars")
}
