package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Game struct {
	name        string
	id          float64
	playercount float64
}

func getGameIDFromName(name string) (string, float64) {
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
		return name, b
	}
	return "notfound", 404
}

// GetGame takes an id and string and return a game
func GetGame(id float64, name string) Game {
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

func main() {
	name, id := (getGameIDFromName("Rust"))
	var gg = GetGame(id, name)
	fmt.Printf("Game: %s, ID: %v, Playercount: %v", gg.name, gg.id, gg.playercount)
}
