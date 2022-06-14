package controllers

import (
	"net/http"

	"github.com/GabrielRendonP/ondemand-go-bootcamp/helpers"
)

func Home(w http.ResponseWriter, r *http.Request) {
	helpers.RenderTemplate(w, "../templates/home.gohtml")
}
