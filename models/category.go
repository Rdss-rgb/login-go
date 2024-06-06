package models

import "github.com/uadmin/uadmin"

type Category struct {
	uadmin.Model
	Name   string `uadmin:"required"`
	Icon   string `uadmin:"image"`
	Active bool
}
