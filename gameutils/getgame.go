// Package gameutils used to fetch playercount and id of a game using only the name of the game.
package gameutils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Game struct idk was forced by my linter to comment here xddddd
type Game struct {
	Name        string
	ID          float64
	Playercount float64
}

func getGameIDFromName(name string) float64 {
	data, err := http.Get("https://raw.githubusercontent.com/wilswe/turbo-playercount-api/master/data/data.json")
	if err != nil {
		panic(err)
	}
	defer data.Body.Close()
	byteValue, _ := ioutil.ReadAll(data.Body)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)

	b, ok := result[name].(float64)

	if ok {
		return b
	}
	return 404
}

// GetGame name and returns a game struct
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
