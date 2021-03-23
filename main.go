package main

import "fmt"

const Circle int8 = -1
const Empty int8 = 0
const Cross int8 = 1

func max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

type Game struct {
	board [9]int8
	turn  int8
}

func (g *Game) init() {
	for idx := range g.board {
		g.board[idx] = Empty
	}
	g.turn = Circle
}

func (g *Game) get_input() {
	var tile int8
	println("Please enter a tile(1-9): ")
	_, err := fmt.Scan(&tile)
	if err != nil {
		fmt.Println(err)
	}
	g.set_tile(tile - 1)
}

func (g Game) print() {
	println("-------------")
	for idx, tile := range g.board {
		if tile == Circle {
			fmt.Print("| O ")
		} else {
			if tile == Cross {
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

func (g *Game) set_tile(cord int8) {
	if g.board[cord] == Empty {
		g.board[cord] = g.turn
		g.turn *= -1
	} else {
		fmt.Println("Tile is already occupied")
	}
}
func (g *Game) takeback(cord int8) {
	if g.board[cord] != Empty {
		g.board[cord] = 0
		g.turn *= -1
	} else {
		fmt.Println("Tile is already empty")
	}
}

func (g Game) get_legal_moves() []int8 {
	var legal_moves []int8
	for idx, tile := range g.board {
		if tile == Empty {
			legal_moves = append(legal_moves, int8(idx))
		}
	}
	return legal_moves
}

func (g Game) is_game_ended() bool {
	row_start := [3]int8{0, 3, 6}
	for _, tile := range row_start {
		if g.board[tile] == g.board[tile+1] && g.board[tile] == g.board[tile+2] && g.board[tile] != Empty {
			return true
		}
	}
	for tile := 0; tile < 3; tile++ {
		if g.board[tile] == g.board[tile+3] && g.board[tile] == g.board[tile+6] && g.board[tile] != Empty {
			return true
		}
	}
	if g.board[0] == g.board[4] && g.board[0] == g.board[8] && g.board[0] != Empty {
		return true
	}
	if g.board[6] == g.board[4] && g.board[6] == g.board[2] && g.board[2] != Empty {
		return true
	}
	legal_moves := g.get_legal_moves()
	if len(legal_moves) == 0 {
		return true
	}
	return false
}
func (g Game) get_result() int8 {
	row_start := [3]int8{0, 3, 6}
	for _, tile := range row_start {
		if g.board[tile] == g.board[tile+1] && g.board[tile] == g.board[tile+2] && g.board[tile] != Empty {
			if g.board[tile] == Circle {
				return Circle
			} else {
				return Cross
			}
		}
	}
	for tile := 0; tile < 3; tile++ {
		if g.board[tile] == g.board[tile+3] && g.board[tile] == g.board[tile+6] && g.board[tile] != Empty {
			if g.board[tile] == Circle {
				return Circle
			} else {
				return Cross
			}
		}
	}
	if g.board[0] == g.board[4] && g.board[0] == g.board[8] && g.board[0] != Empty {
		if g.board[0] == Circle {
			return Circle
		} else {
			return Cross
		}
	}
	if g.board[6] == g.board[4] && g.board[6] == g.board[2] && g.board[2] != Empty {
		if g.board[6] == Circle {
			return Circle
		} else {
			return Cross
		}
	}
	return Empty
}

func (g Game) evaluate() int32 {
	result := g.get_result()
	if result == Circle {
		return -5000
	}
	if result == Cross {
		return 5000
	}
	return 0
}

func (g Game) negamax() int32 {
	if g.is_game_ended() {
		return g.evaluate() * int32(g.turn)
	}
	var v int32 = -10000
	for _, move := range g.get_legal_moves() {
		g.set_tile(move)
		v = max(v, -g.negamax())
		g.takeback(move)
	}
	return v
}
func (g *Game) ai_play() {
	var v int32 = -10000
	var best_move int8
	for _, move := range g.get_legal_moves() {
		g.set_tile(move)
		score := max(v, -g.negamax())
		g.takeback(move)
		if score > v {
			v = score
			best_move = move
		}
	}
	g.set_tile(best_move)
}

func (g *Game) play_vs_ai() {
	g.print()
	for !g.is_game_ended() {
		g.get_input()
		if g.is_game_ended() {
			g.print()
			break
		}
		g.ai_play()
		g.print()
	}
	result := g.get_result()
	if result == Circle {
		fmt.Println("Circle wins!")
	} else {
		if result == Cross {
			fmt.Println("Cross wins!")
		} else {
			fmt.Println("Draw!")

		}
	}
}

func main() {
	var game Game
	game.init()
	game.play_vs_ai()
}
