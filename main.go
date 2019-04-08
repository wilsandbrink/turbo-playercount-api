package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/wilswe/turbo-playercount-api/gameutils"
)

func gameHandler(w http.ResponseWriter, r *http.Request) {
	var name = r.URL.Path[1:]
	var gg = gameutils.GetGame(name)
	fmt.Fprintf(w, "Game: %s, ID: %v, Playercount: %v", gg.Name, gg.ID, gg.Playercount)
}

func main() {
	http.HandleFunc("/", gameHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
