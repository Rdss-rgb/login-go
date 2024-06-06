package main

import (
	"net/http"

	"github.com/Rdss-rgb/login-go/models"
	"github.com/Rdss-rgb/login-go/views"
	"github.com/uadmin/uadmin"
)

func main() {

	uadmin.Register(
		models.Category{}, // <-- place it here
		models.Food{},
	)
	uadmin.RootURL = "/admin/"
	uadmin.SiteName = "Login System"

	// Login System Main Handler
	http.HandleFunc("/", uadmin.Handler(views.MainHandler))
	http.HandleFunc("/login_system/", uadmin.Handler(views.LoginHandler))
	http.HandleFunc("/logout/", uadmin.Handler(views.LogoutHandler))

	// Run the server
	uadmin.StartServer()
}
