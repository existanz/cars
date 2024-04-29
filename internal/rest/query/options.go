package query

import (
	"github.com/gin-gonic/gin"
)

const (
	regNum  = "reg_num"
	mark    = "mark"
	model   = "model"
	year    = "year"
	ownerID = "owner_id"
)

type Paginator struct {
	Page  string
	Limit string
}

type Filter struct {
	Field string
	Value string
}

func GetPaginator(c *gin.Context) Paginator {
	page, _ := c.GetQuery("page")
	limit, _ := c.GetQuery("limit")
	if page == "" {
		page = "1"
	}
	if limit == "" {
		limit = "10"
	}
	return Paginator{
		Page:  page,
		Limit: limit,
	}
}

func GetFilters(c *gin.Context) []Filter {
	var filters []Filter
	regNumVal, _ := c.GetQuery(regNum)
	if regNumVal != "" {
		filters = append(filters, Filter{
			Field: regNum,
			Value: regNumVal,
		})
	}
	markVal, _ := c.GetQuery(mark)
	if markVal != "" {
		filters = append(filters, Filter{
			Field: mark,
			Value: markVal,
		})
	}
	modelVal, _ := c.GetQuery(model)
	if modelVal != "" {
		filters = append(filters, Filter{
			Field: model,
			Value: modelVal,
		})
	}
	yearVal, _ := c.GetQuery(year)
	if yearVal != "" {
		filters = append(filters, Filter{
			Field: year,
			Value: yearVal,
		})
	}
	ownerIDVal, _ := c.GetQuery(ownerID)
	if ownerIDVal != "" {
		filters = append(filters, Filter{
			Field: ownerID,
			Value: ownerIDVal,
		})
	}
	return filters
}
