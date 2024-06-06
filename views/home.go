package views

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Rdss-rgb/login-go/models"
	"github.com/uadmin/uadmin"
)

type CategorizedFood struct {
	Category models.Category
	Food     []models.Food
}

// Handler is used to handle/load pages
func HomeHandler(w http.ResponseWriter, r *http.Request) map[string]interface{} {

	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/home")
	uadmin.Trail(uadmin.DEBUG, "test")
	context := map[string]interface{}{}

	categorized := []CategorizedFood{}

	all_category := []models.Category{}

	uadmin.Filter(&all_category, "active = ?", true)
	fmt.Println("categories: ", all_category)
	for i := range all_category {
		food := []models.Food{}
		uadmin.Filter(&food, "category_id = ?", all_category[i].ID)
		categorized = append(categorized, CategorizedFood{
			Category: all_category[i],
			Food:     food,
		})
	}

	context["Food"] = categorized
	return context
}
