package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/wilswe/turbo-playercount-api/gameutils"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter name of game: ")
		text, _ := reader.ReadString('')
		fmt.Println(text)
		if text == "exit" {
			os.Exit(0)
		} else {
			var game = gameutils.GetGame(text)
			if game.ID == 404 {
				fmt.Println("No game by that name.")
			} else {
				fmt.Printf("Game: %s, ID: %v, Concurrent players: %v\n", game.Name, game.ID, game.Playercount)
			}
		}
	}

}
