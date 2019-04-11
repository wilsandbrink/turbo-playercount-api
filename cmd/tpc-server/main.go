package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/wilswe/turbo-playercount-api/gameutils"
)

func gameHandler(w http.ResponseWriter, r *http.Request) {
	var name = r.URL.Path[1:]
	var gg = gameutils.GetGame(name)
	gameJSON, err := json.Marshal(gg)
	if err != nil {
		panic(err)
	}

	w.Header().Set("Content-Type", "application/json")
	if gg.ID == 404 {
		w.WriteHeader(http.StatusNotFound)
	}
	w.Write(gameJSON)
}

func main() {
	http.HandleFunc("/", gameHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
