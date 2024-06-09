package main

import (
	"errors"
	"fmt"
)

const (
	Empty = iota
	X
	O
	Tie
)
type Coord struct {
	X int
	Y int
}

var board [3][3]int;

func resetBoard() {
	for row := range board {
		for col := range board[row] {
			board[row][col] = Empty;
		}
	}
}

func display() {
	for row := range board {
		println("+===+===+===+")
		print("|")
		for col := range board[row] {
			switch board[row][col] {
				case Empty:
					fmt.Print("   |");
				break;
				case X:
					fmt.Print(" X |");
				break;
				case O:
					fmt.Print(" O |");
				break;
				case Tie:
					fmt.Print(" ? |");
				break;
			}
		}
		fmt.Println()
	}
	println("+===+===+===+")
}

func discardBuffer() {
	var discard string
    fmt.Scanln(&discard)
}
func receiveCords() (Coord, error) {
	var x, y int

	fmt.Println("Type coordinates: ")
	for ok := true; ok; {
		fmt.Print("Col: ")
		_, err2 := fmt.Scanln(&x)
		if err2 != nil {
			discardBuffer();
			fmt.Println("Invaild Col Number")
			continue;
		}
		if 1 > x || 3 < x {
			fmt.Println("Invaild Col Number")
			continue;
		}
		ok = false;
	}
	for ok := true; ok; {
		fmt.Print("Row: ")
		_, err1 := fmt.Scanln(&y)
		if err1 != nil {
			discardBuffer();
			fmt.Println("Invaild Row Number")
			continue;
		}
		if 1 > y || 3 < y {
			fmt.Println("Invaild Row Number")
			continue;
		}
		ok = false;
	}

	coord := Coord{x - 1, y - 1};

	if board[coord.Y][coord.X] != Empty {
		return coord, errors.New("Duplicate Placing");
	}

	return coord, nil;
}

func checkBoard() int {
	var empties int = 0;
	for r := 0 ; r < 3 ; r++ {
		var rowEmpty bool = false;
		if (board[r][0] == Empty) {
			rowEmpty = true;
			empties += 1
		}
		if (board[r][1] == Empty) {
			rowEmpty = true;
			empties += 1
		}
		if (board[r][2] == Empty) {
			rowEmpty = true;
			empties += 1
		}

		if (!rowEmpty && board[r][0] == board[r][1] && board[r][1] == board[r][2]) {
			return board[r][0];
		}
	}
	for c := 0 ; c < 3 ; c++ {
		if (board[0][c] != Empty && board[0][c] == board[1][c] && board[1][c] == board[2][c]) {
			return board[0][c];
		}
	}
	if (board[0][0] != Empty && board[0][0] == board[1][1] && board[1][1] == board[2][2]) {
		return board[0][0];
	}
	if (board[0][2] != Empty && board[0][2] == board[1][1] && board[1][1] == board[2][0]) {
		return board[0][2];
	}
	if (empties == 0) {
		return Tie;
	}
	return Empty;
}

func runGame() {
	resetBoard();
	var currentPiece int = X;
	for gamerun := true ; gamerun ; {
		display();

		cords, err := receiveCords();
		if err != nil {
			println("That spot is already taken. Please select a different spot.")
			continue;
		}
		board[cords.Y][cords.X] = currentPiece;
		currentPiece ^= 0b11;

		win := checkBoard();

		switch win {
			case Empty:
				continue;
			break;
			case X:
				display();
				println("X won!")
			break;
			case O:
				display();
				println("O won!")
			break;
			case Tie:
				display();
				println("It's a Tie!")
			break;
		}
		return;
	}
}

func main() {
	runGame();
}