package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Player struct {
	Name string
	Char string
}

type Game struct {
	Player1 Player
	Player2 Player
	Board   [3][3]string
}

func main() {

	game := Game{}
	game.setup()
}

func (game *Game) setup() {
	game.initializeBoard()
	game.getPlayersInfo()
	game.drawBoard()
}

func (game *Game) getPlayersInfo() {
	fmt.Print("Enter player's 1 name: ")
	reader := bufio.NewReader(os.Stdin)
	p1name, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
		return
	}
	game.Player1.Name = strings.TrimSuffix(p1name, "\n")
	game.Player1.Char = "X"

	for {
		fmt.Print("Enter player's 2 name: ")
		p2name, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again", err)
			return
		}

		if p2name != p1name {
			game.Player2.Name = strings.TrimSuffix(p2name, "\n")
			game.Player2.Char = "O"

			break
		} else {
			fmt.Println("Name already taken, please choose another name")
		}
	}

}

func (game *Game) initializeBoard() {
	for i, line := range game.Board {
		for j, _ := range line {
			game.Board[i][j] = "-"
		}
	}
}

func (game *Game) drawBoard() {
	for i, line := range game.Board {
		for j := range line {
			fmt.Print("|")
			fmt.Print(game.Board[i][j])

		}
		fmt.Print("|")
		fmt.Println("\n-------")
	}
}
