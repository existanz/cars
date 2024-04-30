package rest

import (
	"cars/internal/rest/query"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (router router) getCarsHandler(c *gin.Context) {
	data, err := router.db.GetCars(query.GetPaginator(c), query.GetFilters(c))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, data)
}

func (router router) getCarByIdHandler(c *gin.Context) {
	data, err := router.db.GetCarById(c.Param("id"))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, data)
}
func postCarsHandler(c *gin.Context) {
	c.String(http.StatusOK, "Post cars")
}

func putCarsHandler(c *gin.Context) {
	c.String(http.StatusOK, "Put cars")
}

func (router router) deleteCarHandler(c *gin.Context) {
	err := router.db.DeleteCarById(c.Param("id"))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusOK, "Car id:"+c.Param("id")+" deleted")
}
