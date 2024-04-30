package rest

import (
	"cars/internal/models"
	"cars/internal/rest/query"
	"fmt"
	"net/http"
	"strconv"

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

func (router router) updateCarHandler(c *gin.Context) {
	var car models.Car
	carID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	if err := c.ShouldBindJSON(&car); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(car.Id)
	if car.Id != 0 && carID != car.Id {
		c.String(http.StatusBadRequest, "Wrong id")
		return
	}
	err = router.db.UpdateCarById(carID, car)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	if car.Owner.Id != 0 {
		err := router.db.UpdatePeopleById(car.Owner.Id, car.Owner)
		fmt.Println(car.Owner)
		if err != nil {
			fmt.Println(fmt.Errorf("error in update people: %s", err))
		}
	}
	c.String(http.StatusOK, "Car id:"+strconv.Itoa(carID)+" updated")
}

func (router router) deleteCarHandler(c *gin.Context) {
	err := router.db.DeleteCarById(c.Param("id"))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusOK, "Car id:"+c.Param("id")+" deleted")
}
