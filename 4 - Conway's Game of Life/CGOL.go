package main

import (
	"fmt"
)

var board [][]int;

func createBoard(width int, height int) {
	board = make([][]int, height)
	for i := range board {
		board[i] = make([]int, width)
	}
}

func testSet() {
	board[2][1] = 0b01;
	board[3][2] = 0b01;
	board[1][3] = 0b01;
	board[2][3] = 0b01;
	board[3][3] = 0b01;
}

func display() {
	for row := range board {
		for col := range board[row] {
			if board[row][col] & 0b1 == 1 {
				fmt.Print("* ");
			} else {
				fmt.Print(". ");
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func update() {
	for row := range board {
		for col := range board[row] {
			var colAdded = (col + 1) % (len(board[0]) - 1);
			var colRemoved = (col + len(board[0]) - 1) % len(board[0]);
			var nearbyCount int = 0

			var rowExtra = (row + 1) % len(board);
			nearbyCount += board[rowExtra][colAdded] & 0b1
			nearbyCount += board[rowExtra][col] & 0b1
			nearbyCount += board[rowExtra][colRemoved] & 0b1

			nearbyCount += board[row][colAdded] & 0b1
			nearbyCount += board[row][colRemoved] & 0b1

			rowExtra = (row + len(board) - 1) % len(board);
			nearbyCount += board[rowExtra][colAdded] & 0b1
			nearbyCount += board[rowExtra][col] & 0b1
			nearbyCount += board[rowExtra][colRemoved] & 0b1

			if (nearbyCount == 3) {
				board[row][col] |= 2;
			} else if (nearbyCount == 2 && (board[row][col] & 0b1 == 1)) {
				board[row][col] |= 2;
			}
		}
	}
	for row := range board {
		for col := range board[row] {
			board[row][col] >>= 1;
		}
	}
}

func main() {
	var i = 1;
	createBoard(10, 10);
	testSet();
	for i <= 10 {
		fmt.Println(i)
		display();
		update()
		i += 1;
	}
	display()
}