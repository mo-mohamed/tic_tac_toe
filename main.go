package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Player struct {
	Name string
	Char string
}

type Game struct {
	Player1      Player
	Player2      Player
	Board        [3][3]string
	DefaultBlank string
}

func main() {

	game := Game{}
	game.setup()

	player1Turn := true
	gameEnded := false
	for !gameEnded {
		if player1Turn {
			fmt.Println(game.Player1.Name + "'s turn..")
			c := game.askForLocation(&game.Player1)
			game.Board[c[1]-1][c[0]-1] = game.Player1.Char

			player1Turn = false
		} else {
			fmt.Println(game.Player2.Name + "'s turn")
			c := game.askForLocation(&game.Player2)
			game.Board[c[1]-1][c[0]-1] = game.Player2.Char
			player1Turn = true
		}
		game.drawBoard()
		if game.checkIfWinner() {
			fmt.Println("Game ended")
			game.drawBoard()
			break
		}
	}
}

func (game *Game) checkIfWinner() bool {

	// check horizontal
	for _, row := range game.Board {
		check := make(map[string]int)
		for _, cell := range row {
			if cell == game.DefaultBlank {
				break
			} else {
				if check[cell] == 0 {
					check[cell] = 1
				} else {
					check[cell] += 1
				}

				if check[cell] == 3 {
					return true
				}
			}
		}
	}

	// check vertical
	for col := 0; col < 3; col++ {
		check := make(map[string]int)

		for row := 0; row < 3; row++ {
			if game.Board[row][col] == game.DefaultBlank {
				break
			} else {
				if check[game.Board[row][col]] == 0 {
					check[game.Board[row][col]] = 1
				} else {
					check[game.Board[row][col]] += 1
				}

				if check[game.Board[row][col]] == 3 {
					return true
				}
			}
		}
	}

	//  check diagonal
	check := make(map[string]int)
	for i := 0; i < 3; i++ {
		if game.Board[i][i] == game.DefaultBlank {
			break
		} else {
			if check[game.Board[i][i]] == 0 {
				check[game.Board[i][i]] = 1
			} else {
				check[game.Board[i][i]] += 1
			}

			if check[game.Board[i][i]] == 3 {
				return true
			}
		}
	}

	check = make(map[string]int)
	decremental := 2
	for i := 0; i < 3; i++ {
		if game.Board[i][decremental] == game.DefaultBlank {
			break
		} else {
			if check[game.Board[i][decremental]] == 0 {
				check[game.Board[i][decremental]] = 1
			} else {
				check[game.Board[i][decremental]] += 1
			}

			if check[game.Board[i][decremental]] == 3 {
				return true
			}
		}

		decremental -= 1
	}

	return false
}

func (game *Game) askForLocation(player *Player) []int {
	reader := bufio.NewReader(os.Stdin)
	var coordinateCollected bool = false
	var returnCoordinate []int
	for !coordinateCollected {
		returnCoordinate = make([]int, 0)
		game.drawBoard()
		fmt.Println("Enter location: ")

		// Read input from stdin .. expected two numbers
		value, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occured while reading input. Please try again..", err)
			continue
		}

		// check if two numbers were provided
		coordinates := strings.Split(strings.TrimSpace(value), " ")
		if len(coordinates) != 2 {
			fmt.Println("Please enter two numbers..")
			continue
		}

		// try and parse each value to int
		correctCoordinates := true
		for _, v := range coordinates {
			c, e := strconv.Atoi(strings.TrimSpace(v))
			if e != nil {
				fmt.Println("You must enter a number..")
				correctCoordinates = false
				break
			}

			if c <= 3 && c >= 1 {
				returnCoordinate = append(returnCoordinate, c)
			} else {
				fmt.Println("Please enter numbers between 1 and 3..")
				correctCoordinates = false
				break
			}
		}

		// check if the location was picked before
		if game.Board[returnCoordinate[1]-1][returnCoordinate[0]-1] != game.DefaultBlank {
			fmt.Println("That location was already picked, pick another one..")

			continue
		}

		if correctCoordinates {
			break
		}
	}
	return returnCoordinate
}

func (game *Game) setup() {
	game.DefaultBlank = "-"
	game.initializeBoard()
	game.getPlayersInfo()
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
		for j := range line {
			game.Board[i][j] = game.DefaultBlank
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
