package main // Обязательная строка для исполняемых программ

import (
	"fmt"
	"math/rand"
	"time"
)

type Game struct {
	board [3][3]string
	turn  string
}

func NewGame() *Game {
	return &Game{
		board: [3][3]string{
			{" ", " ", " "},
			{" ", " ", " "},
			{" ", " ", " "},
		},
		turn: "X",
	}
}

func (g *Game) DisplayBoard() {
	fmt.Println("  0 1 2")
	for i, row := range g.board {
		fmt.Printf("%d ", i)
		for _, cell := range row {
			fmt.Printf("%s ", cell)
		}
		fmt.Println()
	}
}

func (g *Game) MakeMove(row, col int) bool {
	if row < 0 || row >= 3 || col < 0 || col >= 3 || g.board[row][col] != " " {
		return false
	}
	g.board[row][col] = g.turn
	g.turn = g.SwitchPlayer()
	return true
}

func (g *Game) SwitchPlayer() string {
	if g.turn == "X" {
		return "O"
	}
	return "X"
}

func (g *Game) CheckWin() bool {
	// Проверка строк и столбцов
	for i := 0; i < 3; i++ {
		if g.board[i][0] != " " && g.board[i][0] == g.board[i][1] && g.board[i][1] == g.board[i][2] {
			return true
		}
		if g.board[0][i] != " " && g.board[0][i] == g.board[1][i] && g.board[1][i] == g.board[2][i] {
			return true
		}
	}
	// Проверка диагоналей
	if g.board[0][0] != " " && g.board[0][0] == g.board[1][1] && g.board[1][1] == g.board[2][2] {
		return true
	}
	if g.board[0][2] != " " && g.board[0][2] == g.board[1][1] && g.board[1][1] == g.board[2][0] {
		return true
	}
	return false
}

func (g *Game) CheckDraw() bool {
	for _, row := range g.board {
		for _, cell := range row {
			if cell == " " {
				return false
			}
		}
	}
	return true
}

func (g *Game) ComputerMove() {
	rand.Seed(time.Now().UnixNano())
	for {
		row := rand.Intn(3)
		col := rand.Intn(3)
		if g.MakeMove(row, col) {
			break
		}
	}
}

func main() {
	game := NewGame()
	var mode int
	fmt.Println("Выберите режим игры: 1 - два игрока, 2 - против компьютера")
	fmt.Scan(&mode)

	for {
		game.DisplayBoard()

		if mode == 1 || game.turn == "X" {
			var row, col int
			fmt.Printf("Игрок %s, введите строку и столбец (0-2): ", game.turn)
			fmt.Scan(&row, &col)

			if !game.MakeMove(row, col) {
				fmt.Println("Некорректный ход, попробуйте снова.")
				continue
			}
		} else {
			game.ComputerMove()
			fmt.Println("Компьютер сделал ход.")
		}

		if game.CheckWin() {
			game.DisplayBoard()
			fmt.Printf("Игрок %s выиграл!\n", game.SwitchPlayer())
			break
		}

		if game.CheckDraw() {
			game.DisplayBoard()
			fmt.Println("Ничья!")
			break
		}
	}
}
