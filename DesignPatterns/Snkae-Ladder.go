package designpatterns

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	UML Mapping:
	- Game HAS Board, Dice, Players
	- Board HAS Cells
	- Cell HAS optional Jump
	- Jump (start,end) => snake or ladder
*/

// ---------------------------
// Jump (Snake or Ladder)
// ---------------------------

type Jump struct {
	start int
	end   int
}

// ---------------------------
// Cell
// ---------------------------

type Cell struct {
	jump *Jump
}

// ---------------------------
// Board
// ---------------------------

type Board struct {
	cells []Cell
	size  int
}

func NewBoard(size int, jumps []Jump) *Board {
	board := &Board{
		cells: make([]Cell, size+1), // index 0 unused
		size:  size,
	}

	// Insert jumps (snakes/ladders)
	for _, j := range jumps {
		board.cells[j.start].jump = &Jump{start: j.start, end: j.end}
	}

	return board
}

// ---------------------------
// Player
// ---------------------------

type Player struct {
	id  string
	pos int
}

// ---------------------------
// Dice
// ---------------------------

type Dice struct {
	diceCount int
}

func NewDice(count int) *Dice {
	return &Dice{diceCount: count}
}

func (d *Dice) Roll() int {
	rand.Seed(time.Now().UnixNano())
	total := 0

	for i := 0; i < d.diceCount; i++ {
		total += rand.Intn(6) + 1
	}

	return total
}

// ---------------------------
// Game
// ---------------------------

type Game struct {
	board   *Board
	dice    *Dice
	players []Player
}

func NewGame(board *Board, dice *Dice, players []Player) *Game {
	return &Game{
		board:   board,
		dice:    dice,
		players: players,
	}
}

func (g *Game) Start() {
	turn := 0

	for {
		player := &g.players[turn%len(g.players)]

		fmt.Printf("\nPlayer %s at %d rolling dice...\n", player.id, player.pos)

		roll := g.dice.Roll()
		fmt.Printf("Dice rolled: %d\n", roll)

		nextPos := player.pos + roll

		if nextPos > g.board.size {
			fmt.Println("Roll exceeds board size, waiting next turn...")
		} else {
			player.pos = nextPos
			fmt.Printf("Player moved to %d\n", player.pos)

			// Check for jump
			cell := g.board.cells[player.pos]
			if cell.jump != nil {
				if cell.jump.end > cell.jump.start {
					fmt.Printf("Yay! Ladder from %d to %d\n", cell.jump.start, cell.jump.end)
				} else {
					fmt.Printf("Oh No! Snake from %d to %d\n", cell.jump.start, cell.jump.end)
				}
				player.pos = cell.jump.end
				fmt.Printf("Player new position: %d\n", player.pos)
			}
		}

		// Check win
		if player.pos == g.board.size {
			fmt.Printf("\n🎉 Player %s wins the game!\n", player.id)
			return
		}

		turn++
	}
}

// ---------------------------
// Main
// ---------------------------

func PlaySnakeLadder() {
	// Define snakes & ladders
	jumps := []Jump{
		{start: 3, end: 22},  // Ladder
		{start: 5, end: 8},   // Ladder
		{start: 11, end: 26}, // Ladder
		{start: 20, end: 29}, // Ladder
		{start: 27, end: 1},  // Snake
		{start: 21, end: 9},  // Snake
		{start: 17, end: 4},  // Snake
	}

	board := NewBoard(30, jumps)
	dice := NewDice(1)

	players := []Player{
		{id: "A", pos: 0},
		{id: "B", pos: 0},
	}

	game := NewGame(board, dice, players)
	game.Start()
}
