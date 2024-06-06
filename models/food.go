package models

import "github.com/uadmin/uadmin"

type Food struct {
	uadmin.Model
	Icon        string `uadmin:"image"`
	Name        string `uadmin:"required"`
	Description string `uadmin:"html"`
	Price       float64
	Discount    float64
	Category    Category
	CategoryID  uint
	Active      bool
}
