package rest

import (
	"cars/internal/models"
	externalapi "cars/internal/rest/external-api"
	"cars/internal/rest/query"

	"log/slog"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get cars with pagination and filters
// @Description Get cars with pagination and filters
// @Accept  json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Number of items per page" default(10)
// @Param regNum query string false "Registration number"
// @Param mark query string false "Car mark"
// @Param model query string false "Car model"
// @Param year query int false "Car year"
// @Param ownerID query int false "Owner id"
// @Success 200 {array} models.Car "OK"
// @Failure 500 {string} string "Internal server error"
// @Router /cars [get]
func (router router) getCarsHandler(c *gin.Context) {
	data, err := router.db.GetCars(query.GetPaginator(c), query.GetFilters(c))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, data)
}

// @Summary Get car by id
// @Description Get car by id
// @ID id
// @Produce json
// @Param id path int true "Car ID"
// @Success 200 {object} models.Car
// @Failure 500 {string} string "Internal server error"
// @Router /cars/{id} [get]
func (router router) getCarByIdHandler(c *gin.Context) {
	data, err := router.db.GetCarById(c.Param("id"))
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, data)
}

// @Summary Update car by id
// @Description Update car by id
// @ID id
// @Accept  json
// @Produce json
// @Param id path int true "Car ID"
// @Param car body models.Car true "Car data"
// @Success 200 {string} string "Car id:{id} updated"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /cars/{id} [put]
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
		if err != nil {
			slog.Debug("error in update people: %s", err)
		}
	}
	c.String(http.StatusOK, "Car id:"+strconv.Itoa(carID)+" updated")
}

// @Summary Deletes a car by its ID
// @Description Deletes a car by its ID
// @ID id
// @Accept  json
// @Produce json
// @Param id path int true "Car ID"
// @Success 200 {string} string "Car id:{id} deleted"
// @Failure 500 {string} string "Internal server error"
// @Router /cars/{id} [delete]
func (router router) deleteCarHandler(c *gin.Context) {
	id := c.Param("id")
	if err := router.db.DeleteCarById(id); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.String(http.StatusOK, "Car id:"+id+" deleted")
}

// @Summary Add new cars
// @Description Add new cars
// @ID add-new-cars
// @Accept  json
// @Produce json
// @Param   body      body   []string  true  "registration numbers"
// @Success 200 {string} string "Cars added"
// @Failure 400 {string} string "Bad request"
// @Failure 500 {string} string "Internal server error"
// @Router /cars [post]
func (router router) addNewCarsHandler(c *gin.Context) {
	var regNums struct {
		Nums []string `json:"regNums"`
	}
	if err := c.ShouldBindJSON(&regNums); err != nil {
		c.JSON(http.StatusBadRequest, "Bad request")
		return
	}
	for _, regNum := range regNums.Nums {
		car, err := externalapi.GetCarByRegNum(regNum)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
		if err := router.db.AddNewCar(car); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}
	c.String(http.StatusOK, "Cars added")
}
