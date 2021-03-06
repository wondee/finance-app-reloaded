package main

import (
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type JsonSpecialCost struct {
	ID      int        `json:"id"`
	Name    string     `json:"name" binding:"required"`
	Amount  int        `json:"amount" binding:"required"`
	DueDate *YearMonth `json:"dueDate" binding:"required"`
}

func GetSpecialCosts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, createSpecialCosts(LoadCurrentFinanceId(c)))
}

func SaveSpecialCosts(c *gin.Context) {
	var cost JsonSpecialCost
	err := c.ShouldBindJSON(&cost)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	dbObject := SpecialCost{
		ID:        cost.ID,
		Name:      cost.Name,
		Amount:    cost.Amount,
		DueDate:   cost.DueDate,
		FinanceID: LoadCurrentFinanceId(c),
	}

	err = SaveSpecialCost(LoadCurrentFinanceId(c), &dbObject)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
}

func DeleteSpecialCosts(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)

	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	DeleteSpecialCost(id, LoadCurrentFinanceId(c))
}

func createSpecialCosts(financeId int) (result []JsonSpecialCost) {
	result = make([]JsonSpecialCost, 0)

	specialCosts := LoadSpecialCosts(financeId)

	for _, cost := range *specialCosts {
		result = append(result, JsonSpecialCost{
			ID:      cost.ID,
			Name:    cost.Name,
			Amount:  cost.Amount,
			DueDate: cost.DueDate,
		})
	}

	sort.Slice(result, func(a, b int) bool {
		return strings.ToLower(result[a].Name) < strings.ToLower(result[b].Name)
	})

	return
}
