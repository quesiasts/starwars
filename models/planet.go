package models

import (
	"starwars/db"
)

type Planet struct {
	Id      string
	Name    string
	Climate string
	Terrain string
}

func searchAllPlanet() []Planet {
	db := db.ConnectDatabase()

	selectAllPlanets, err := db.Query("select * from planets")
	if err != nil {
		panic(err.Error())
	}

	p := Planet{}
	planets := []Planet{}

	for selectAllPlanets.Next() {
		var id string
		var name, climate, terrain string

		err = selectAllPlanets.Scan(&id, &name, &climate, &terrain)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Name = name
		p.Climate = climate
		p.Terrain = terrain

		planets = append(planets, p)
	}
	defer db.Close()
	return planets
}

func CreatingNewPlanet(name, climate, terrain string) {
	db := db.ConnectDatabase()

	insertPlanetOnDatabase, err := db.Prepare("insert string planets(name, climate, terrain) values ($1, $2, $3)")
	if err != nil {
		panic(err.Error())
	}
	insertPlanetOnDatabase.Exec(name, climate, terrain)
	defer db.Close()
}

func DeletePlanet(id string) {
	db := db.ConnectDatabase()

	deletePlanetOnDatabase, err := db.Prepare("delete from planets where id=$1")
	if err != nil {
		panic(err.Error())
	}
	deletePlanetOnDatabase.Exec(id)
	defer db.Close()
}

func EditPlanet(id string) Planet {
	db := db.ConnectDatabase()

	planetOnDatabase, err := db.Query("select * from planets where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	planetForUpdate := Planet{}

	for planetOnDatabase.Next() {
		var id string
		var name, climate, terrain string

		err := planetOnDatabase.Scan(&id, &name, &climate, &terrain)
		if err != nil {
			panic(err.Error())
		}

		planetForUpdate.Id = id
		planetForUpdate.Name = name
		planetForUpdate.Climate = climate
		planetForUpdate.Terrain = terrain
	}
	defer db.Close()
	return planetForUpdate
}

func updatePlanets(id string, name, climate, terrain string) {
	db := db.ConnectDatabase()
	updatePlanet, err := db.Prepare("update planets set name=$1, climate=$2, terrain=$3 where id=$4")
	if err != nil {
		panic(err.Error())
	}
	updatePlanet.Exec(name, climate, terrain, id)
	defer db.Close()
}
