package main

import (
	"database/sql/driver"
	"errors"
	"strconv"
	"strings"
)

type FixedCost struct {
	ID        int `gorm:primary_key`
	Name      string
	Amount    int
	From      *YearMonth
	To        *YearMonth
	DueMonth  months `gorm:"type:string"`
	FinanceID int
}

type months []int

var ALL_MONTHS = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

func LoadFixedCosts(financeId int) *[]FixedCost {

	var costs []FixedCost
	DB.Where("finance_id = ?", financeId).Find(&costs)

	return &costs
}

func SaveFixedObject(financeId int, cost *FixedCost) error {

	costs := LoadFixedCosts(financeId)

	if len(*costs) >= MAX_COSTS {
		return errors.New("max count of costs already filled")
	}

	if cost.ID == 0 {
		DB.Create(cost)
	} else {
		checkPermissionFixedCost(cost.ID, cost.FinanceID)
		DB.Save(cost)
	}

	return nil
}

func DeleteFixedCost(id int, financeId int) {
	checkPermissionFixedCost(id, financeId)
	DB.Delete(&FixedCost{}, id)
}

func (this *months) Scan(value interface{}) error {
	str := value.(string)

	if len(str) == 0 {
		return nil
	}

	for _, month := range strings.Split(str, " ") {
		value, err := strconv.Atoi(month)
		if err == nil {
			*this = append(*this, value)
		} else {
			return err
		}
	}

	return nil
}

func (this months) Value() (driver.Value, error) {
	if this == nil {
		return nil, nil
	}

	var result string
	for i, month := range this {
		result += strconv.Itoa(month)

		if i != len(this) {
			result += " "
		}
	}

	return strings.TrimSpace(result), nil

}

func checkPermissionFixedCost(id int, financeId int) {
	var cost FixedCost

	DB.Where("id = ?", id).First(&cost)

	if cost.FinanceID != financeId {
		panic("access to costs not in your finance not allowed")
	}
}
