package controllers

import (
	"net/http"

	"github.com/GabrielRendonP/ondemand-go-bootcamp/helpers"
)

// Home displays home.html page
func Home(w http.ResponseWriter, r *http.Request) {
	helpers.RenderTemplate(w, "../templates/home.gohtml")
}
