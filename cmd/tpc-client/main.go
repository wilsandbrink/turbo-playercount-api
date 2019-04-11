package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/wilswe/turbo-playercount-api/gameutils"
)

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter name of game: ")
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1]
		fmt.Println(text)
		var game = gameutils.GetGame(text)
		if text == "exit" {
			os.Exit(0)
		} else {
			fmt.Printf("Game: %s, ID: %v, Concurrent players: %v\n", game.Name, game.ID, game.Playercount)
		}
	}
}
