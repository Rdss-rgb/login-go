package views

import (
	"net/http"
	"strings"

	"github.com/uadmin/uadmin"
)

// MainHandler is the main handler for the login system.
func MainHandler(w http.ResponseWriter, r *http.Request) { //needed prefix initialized from main.go
	page := strings.TrimPrefix(r.URL.Path, "/")
	page = strings.TrimSuffix(page, "/")
	session := uadmin.IsAuthenticated(r)

	if session == nil {
		http.Redirect(w, r, "/login_system/", http.StatusSeeOther)
	}

	c := map[string]interface{}{}
	switch page {
	case "home": //Name of HTML
		uadmin.Trail(uadmin.DEBUG, "test1")
		c = HomeHandler(w, r)
	case "dashboard": //Name of HTML
		c = DashboardHandler(w, r)
	default:
		uadmin.Trail(uadmin.DEBUG, "test2")
		page = "home"
	}
	//uadmin.Trail(uadmin.DEBUG, page)
	c["Page"] = page
	uadmin.Trail(uadmin.DEBUG, "page: ", page)

	Rendering(w, r, page, c)
}

func Rendering(w http.ResponseWriter, r *http.Request, page string, context map[string]interface{}) {
	templateList := []string{}
	// templateList = append(templateList, "./templates/base.html")

	path := "./templates/" + page + ".html"
	templateList = append(templateList, path)

	uadmin.RenderMultiHTML(w, r, templateList, context)
}
