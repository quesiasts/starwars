package controller

import (
	"log"
	"net/http"
	"starwars/models"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	allPlanets := models.searchAllPlanet()
	temp.ExecuteTemplate(w, "Index", allPlanets)
}

func NewPlanet(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		climate := r.FormValue("climate")
		terrain := r.FormValue("terrain")

		models.CreatingNewPlanet(name, climate, terrain)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idOfPlanet := r.URL.Query().Get("id")
	models.DeletePlanet(idOfPlanet)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idOfPlanet := r.URL.Query().Get("id")
	planet := models.EditPlanet(idOfPlanet)
	temp.ExecuteTemplate(w, "Edit", planet)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		climate := r.FormValue("climate")
		terrain := r.FormValue("terrain")

		idConvertido, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do id: ", err)
		}

		models.updatePlanets(idConvertido, name, climate, terrain)
	}
	http.Redirect(w, r, "/", 301)
}
