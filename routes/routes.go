package routes

import (
	"net/http"
	"starwars/controller"
)

func LoadRoutes() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/new", controller.NewPlanet)
	http.HandleFunc("/insert", controller.Insert)
	http.HandleFunc("/delete", controller.Delete)
	http.HandleFunc("/edit", controller.Edit)
	http.HandleFunc("/update", controller.Update)
}
