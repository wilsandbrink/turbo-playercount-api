// Package gameutils is not meant to be imported outside of turbo-playercount-api due to data folder
package gameutils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Game struct {
	Name        string
	ID          float64
	Playercount float64
}

func getGameIDFromName(name string) float64 {
	data, err := os.Open("data/data.json")
	if err != nil {
		panic(err)
	}
	defer data.Close()
	byteValue, _ := ioutil.ReadAll(data)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	b, ok := result[name].(float64)

	if ok {
		return b
	}
	return 404
}

// GetGame takes an id and name and returns a game
func GetGame(name string) Game {
	var id = getGameIDFromName(name)
	if id == 404 {
		g := Game{"Not found", 404, 0}
		return g
	}
	url := fmt.Sprintf("https://api.steampowered.com/ISteamUserStats/GetNumberOfCurrentPlayers/v1/?appid=%v", id)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	byteValue, _ := ioutil.ReadAll(resp.Body)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	var ree = result["response"]
	reemaps := ree.(map[string]interface{})

	pc, ok := reemaps["player_count"].(float64)

	if ok {
		g := Game{name, id, pc}
		return g
	}

	g := Game{"Not found", 404, 0}
	return g
}
