package designpatterns

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Cell int

const (
	Empty Cell = iota
	X
	O
)

func (c Cell) String() string {
	switch c {
	case X:
		return "X"
	case O:
		return "O"
	default:
		return " "
	}
}

type Board struct {
	grid [3][3]Cell
}

func NewBoard() *Board { return &Board{} }

func (b *Board) Place(r, c int, p Cell) error {
	if r < 0 || r > 2 || c < 0 || c > 2 {
		return errors.New("invalid coordinates")
	}
	if b.grid[r][c] != Empty {
		return errors.New("cell already occupied")
	}
	b.grid[r][c] = p
	return nil
}

func (b *Board) IsFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.grid[i][j] == Empty {
				return false
			}
		}
	}
	return true
}

// CheckWin returns (winner, true) if found
func (b *Board) CheckWin() (Cell, bool) {
	// rows & cols
	for i := 0; i < 3; i++ {
		if b.grid[i][0] != Empty && b.grid[i][0] == b.grid[i][1] && b.grid[i][1] == b.grid[i][2] {
			return b.grid[i][0], true
		}
		if b.grid[0][i] != Empty && b.grid[0][i] == b.grid[1][i] && b.grid[1][i] == b.grid[2][i] {
			return b.grid[0][i], true
		}
	}
	// diagonals
	if b.grid[0][0] != Empty && b.grid[0][0] == b.grid[1][1] && b.grid[1][1] == b.grid[2][2] {
		return b.grid[0][0], true
	}
	if b.grid[0][2] != Empty && b.grid[0][2] == b.grid[1][1] && b.grid[1][1] == b.grid[2][0] {
		return b.grid[0][2], true
	}
	return Empty, false
}

func (b *Board) AvailableMoves() [][2]int {
	var moves [][2]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b.grid[i][j] == Empty {
				moves = append(moves, [2]int{i, j})
			}
		}
	}
	return moves
}

func (b *Board) Copy() *Board {
	nb := NewBoard()
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			nb.grid[i][j] = b.grid[i][j]
		}
	}
	return nb
}

func (b *Board) String() string {
	var sb strings.Builder
	for i := 0; i < 3; i++ {
		sb.WriteString(" " + b.grid[i][0].String() + " | " + b.grid[i][1].String() + " | " + b.grid[i][2].String() + " \n")
		if i < 2 {
			sb.WriteString("---+---+---\n")
		}
	}
	return sb.String()
}

// ----- Game -----
type Game struct {
	Board   *Board
	Current Cell
	Winner  Cell
	Over    bool
}

func NewGame(start Cell) *Game {
	return &Game{
		Board:   NewBoard(),
		Current: start,
		Winner:  Empty,
		Over:    false,
	}
}

func (g *Game) MakeMove(r, c int) error {
	if g.Over {
		return errors.New("game is over")
	}
	if err := g.Board.Place(r, c, g.Current); err != nil {
		return err
	}
	if w, ok := g.Board.CheckWin(); ok {
		g.Winner = w
		g.Over = true
		return nil
	}
	if g.Board.IsFull() {
		g.Over = true
		return nil
	}
	g.Switch()
	return nil
}

func (g *Game) Switch() {
	if g.Current == X {
		g.Current = O
	} else {
		g.Current = X
	}
}

// ----- Minimax AI (unbeatable) -----
func (g *Game) BestMove(ai Cell) (int, int) {
	bestScore := math.Inf(-1)
	bestR, bestC := -1, -1
	for _, mv := range g.Board.AvailableMoves() {
		r, c := mv[0], mv[1]
		nb := g.Board.Copy()
		nb.Place(r, c, ai)
		score := minimax(nb, ai, false, ai)
		if score > bestScore {
			bestScore = score
			bestR, bestC = r, c
		}
	}
	return bestR, bestC
}

func scoreForWinner(w Cell, ai Cell) float64 {
	if w == ai {
		return 1
	}
	if w == Empty {
		return 0
	}
	return -1
}

func minimax(b *Board, ai Cell, isMaximizing bool, aiCell Cell) float64 {
	if w, ok := b.CheckWin(); ok {
		return scoreForWinner(w, aiCell)
	}
	if b.IsFull() {
		return 0
	}

	if isMaximizing {
		best := math.Inf(-1)
		for _, mv := range b.AvailableMoves() {
			nb := b.Copy()
			nb.Place(mv[0], mv[1], ai)
			val := minimax(nb, ai, false, aiCell)
			if val > best {
				best = val
			}
		}
		return best
	} else {
		// minimizing => opponent move
		op := X
		if ai == X {
			op = O
		}
		best := math.Inf(1)
		for _, mv := range b.AvailableMoves() {
			nb := b.Copy()
			nb.Place(mv[0], mv[1], op)
			val := minimax(nb, ai, true, aiCell)
			if val < best {
				best = val
			}
		}
		return best
	}
}

// ----- Console play demonstration -----
func Play() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Tic-Tac-Toe")
	fmt.Println("Choose mode: 1) Human vs Human  2) Human (X) vs AI (O)")
	fmt.Print("Enter choice: ")
	text, _ := reader.ReadString('\n')
	choice := strings.TrimSpace(text)

	if choice == "1" {
		playHumanVsHuman()
	}
}

func playHumanVsHuman() {
	g := NewGame(X)
	reader := bufio.NewReader(os.Stdin)
	for !g.Over {
		fmt.Println(g.Board)
		fmt.Printf("Player %s - enter move (row col) 0-based: ", g.Current)
		line, _ := reader.ReadString('\n')
		fields := strings.Fields(line)
		if len(fields) < 2 {
			fmt.Println("invalid input")
			continue
		}
		r, _ := strconv.Atoi(fields[0])
		c, _ := strconv.Atoi(fields[1])
		if err := g.MakeMove(r, c); err != nil {
			fmt.Println("error:", err)
			continue
		}
	}
	fmt.Println(g.Board)
	if g.Winner == Empty {
		fmt.Println("Game ended in a draw")
	} else {
		fmt.Println("Winner:", g.Winner)
	}
}
