package main

import "fmt"

type Board struct {
	cells [3][3]rune
}

type Player rune

type Game struct {
	board         Board
	currentPlayer Player
	turn          int
}

func NewBoard() *Board {
	var row = [3]rune{'ー', 'ー', 'ー'}

	return &Board{
		cells: [3][3]rune{
			row,
			row,
			row,
		},
	}
}

func initGame() *Game {
	return &Game{
		board:         *NewBoard(),
		currentPlayer: Player('〇'),
		turn:          0,
	}
}

func gameLoop(game Game) {
	for {
		game.turn += 1

		fmt.Printf("ターン%d: %cのターンです\n", game.turn, game.currentPlayer)

		for {
			x, y := getPlayerInput()

			if inValidMove(x, y, game.board) {
				placeMark(x, y, &game)
				break
			} else {
				fmt.Println("そこには置けません")
				continue
			}
		}

		drawBoard(game.board)

		if checkWin(game.board, game.currentPlayer) {
			fmt.Printf("%cが勝ちました。もう一度プレイしますか？ Y/N\n", game.currentPlayer)
			if replay() {
				game.turn = 0
				continue
			} else {
				panic("See you again")
			}
		}

		if game.turn == 9 {
			fmt.Println("引き分けです。もう一度プレイしますか？ Y/N")
			if replay() {
				game.turn = 0
				continue
			} else {
				panic("See you again")
			}
		}

		switchPlayer(&game)
	}
}

func replay() bool {
	var i string

	for {
		fmt.Scan(&i)

		fmt.Println(i)

		switch i {
		case "Y":
			return true
		case "N":
			return false
		default:
			continue
		}
	}

}

func drawBoard(board Board) {
	for y := range 3 {
		for x := range 3 {
			fmt.Printf("%c", board.cells[y][x])
		}
		fmt.Println()
	}
}

func placeMark(x int, y int, game *Game) {
	game.board.cells[x-1][y-1] = rune(game.currentPlayer)
}

func inValidMove(x int, y int, board Board) bool {
	if board.cells[x-1][y-1] == rune('ー') {
		return true
	} else {
		return false
	}
}

func checkWin(board Board, player Player) bool {
	f := false

	// 横
	for _, r := range board.cells {
		if r[0] == 'ー' {
			continue
		}
		if r[0] == r[1] && r[1] == r[2] {
			f = true
		} else {

		}
	}

	// 縦
	for x := range 3 {
		if board.cells[0][x] == rune(player) && board.cells[2][x] == rune(player) {
			f = true
		} else if board.cells[2][x] == rune(player) && board.cells[0][x] == rune(player) {
			f = true
		}
	}

	// 斜め
	if board.cells[1][1] == rune(player) {
		if board.cells[0][0] == rune(player) && board.cells[2][2] == rune(player) {
			f = true
		} else if board.cells[2][0] == rune(player) && board.cells[0][2] == rune(player) {
			f = true
		}
	}

	return f
}

func switchPlayer(game *Game) {
	if game.currentPlayer == '〇' {
		game.currentPlayer = '✖'
	} else {
		game.currentPlayer = '〇'
	}
}

func getPlayerInput() (int, int) {
	var x, y int

	fmt.Printf("どこに置きますか？")

	for {
		fmt.Scan(&x, &y)
		if x < 1 || x > 3 || y < 1 || y > 3 {
			fmt.Println("範囲外です")
			continue
		}
		break
	}

	return x, y
}

func main() {
	game := initGame()
	gameLoop(*game)
}
