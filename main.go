package main

import "fmt"

const circle int8 = -1
const empty int8 = 0
const cross int8 = 1

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

type game struct {
	board [9]int8
	turn  int8
}

func (g *game) init() {
	for idx := range g.board {
		g.board[idx] = empty
	}
	g.turn = circle
}

func (g *game) input() {
	var tile int8
	println("Please enter a tile(1-9): ")
	_, err := fmt.Scan(&tile)
	if err != nil {
		fmt.Println(err)
	}
	g.setTitle(tile - 1)
}

func (g game) print() {
	println("-------------")
	for idx, tile := range g.board {
		if tile == circle {
			fmt.Print("| O ")
		} else {
			if tile == cross {
				fmt.Print("| X ")
			} else {
				fmt.Print("|   ")
			}
		}
		if (idx+1)%3 == 0 {
			fmt.Println("|\n-------------")
		}

	}
}

func (g *game) setTitle(cord int8) {
	if g.board[cord] == empty {
		g.board[cord] = g.turn
		g.turn *= -1
	} else {
		fmt.Println("Tile is already occupied")
	}
}
func (g *game) takeback(cord int8) {
	if g.board[cord] != empty {
		g.board[cord] = 0
		g.turn *= -1
	} else {
		fmt.Println("Tile is already empty")
	}
}

func (g game) legalMoves() []int8 {
	var legalMoves []int8
	for idx, tile := range g.board {
		if tile == empty {
			legalMoves = append(legalMoves, int8(idx))
		}
	}
	return legalMoves
}

func (g game) isGameEnded() bool {
	rowStart := [3]int8{0, 3, 6}
	for _, tile := range rowStart {
		if g.board[tile] == g.board[tile+1] && g.board[tile] == g.board[tile+2] && g.board[tile] != empty {
			return true
		}
	}
	for tile := 0; tile < 3; tile++ {
		if g.board[tile] == g.board[tile+3] && g.board[tile] == g.board[tile+6] && g.board[tile] != empty {
			return true
		}
	}
	if g.board[0] == g.board[4] && g.board[0] == g.board[8] && g.board[0] != empty {
		return true
	}
	if g.board[6] == g.board[4] && g.board[6] == g.board[2] && g.board[2] != empty {
		return true
	}
	legalMoves := g.legalMoves()
	if len(legalMoves) == 0 {
		return true
	}
	return false
}
func (g game) result() int8 {
	rowStart := [3]int8{0, 3, 6}
	for _, tile := range rowStart {
		if g.board[tile] == g.board[tile+1] && g.board[tile] == g.board[tile+2] && g.board[tile] != empty {
			if g.board[tile] == circle {
				return circle
			} else {
				return cross
			}
		}
	}
	for tile := 0; tile < 3; tile++ {
		if g.board[tile] == g.board[tile+3] && g.board[tile] == g.board[tile+6] && g.board[tile] != empty {
			if g.board[tile] == circle {
				return circle
			} else {
				return cross
			}
		}
	}
	if g.board[0] == g.board[4] && g.board[0] == g.board[8] && g.board[0] != empty {
		if g.board[0] == circle {
			return circle
		} else {
			return cross
		}
	}
	if g.board[6] == g.board[4] && g.board[6] == g.board[2] && g.board[2] != empty {
		if g.board[6] == circle {
			return circle
		} else {
			return cross
		}
	}
	return empty
}

func (g game) evaluate() int32 {
	result := g.result()
	if result == circle {
		return -5000
	}
	if result == cross {
		return 5000
	}
	return 0
}

func (g game) negamax() int32 {
	if g.isGameEnded() {
		return g.evaluate() * int32(g.turn)
	}
	var v int32 = -10000
	for _, move := range g.legalMoves() {
		g.setTitle(move)
		v = max(v, -g.negamax())
		g.takeback(move)
	}
	return v
}
func (g *game) aiPlay() {
	var v int32 = -10000
	var bestMove int8
	for _, move := range g.legalMoves() {
		g.setTitle(move)
		score := max(v, -g.negamax())
		g.takeback(move)
		if score > v {
			v = score
			bestMove = move
		}
	}
	g.setTitle(bestMove)
}

func (g *game) playVsAi() {
	g.print()
	for !g.isGameEnded() {
		g.input()
		if g.isGameEnded() {
			g.print()
			break
		}
		g.aiPlay()
		g.print()
	}
	result := g.result()
	if result == circle {
		fmt.Println("circle wins!")
	} else {
		if result == cross {
			fmt.Println("cross wins!")
		} else {
			fmt.Println("Draw!")

		}
	}
}

func main() {
	var game game
	game.init()
	game.playVsAi()
}
