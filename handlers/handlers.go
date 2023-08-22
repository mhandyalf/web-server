package handlers

import (
	"encoding/json"
	"net/http"
	"web-server/database"
)

func HeroesHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetDBConnection()
	defer db.Close()

	heroes := database.GetHeroesFromDB(db)
	writeJSONResponse(w, heroes)
}

func VillainsHandler(w http.ResponseWriter, r *http.Request) {
	db := database.GetDBConnection()
	defer db.Close()

	villains := database.GetVillainsFromDB(db)
	writeJSONResponse(w, villains)
}

func writeJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
