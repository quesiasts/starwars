package main

import (
	"net/http"
	"starwars/routes"
)

type Response struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Climate string `json:"climate"`
	Terrain string `json:"terrain"`
}

func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8080", nil)
}
