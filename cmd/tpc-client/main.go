package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/wilswe/turbo-playercount-api/gameutils"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "exit" {
			os.Exit(0)
		}
		var game = gameutils.GetGame(scanner.Text())
		fmt.Printf("Game: %s, Concurrent players: %v, ID: %v\n", game.Name, game.Playercount, game.ID)
	}

	if scanner.Err() != nil {
		// handle error.
	}
}
