package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func getGameID(name string) float64 {
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

func main() {
	fmt.Print(getGameID("Rust"))
}
