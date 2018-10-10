// Copyright © 2017 Mark F. Pleso. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

// senet board game

// TODO web
// TODO automated move
// TODO add cmdline options
// TODO special moves
// TODO pl2

package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

const (
	P = 2 //pyramids
	S = 3 //spools
)

type Board [30]int // type

func main() {
	board := Board{S, P, S, P, S, P, S, P, S, P, S, P, S, P} // add opt 5

	pl1Token := P // add opt S

	for token := P; ; token ^= 0x1 {
		for {
			fmt.Println("TOKEN=", token)
			printBoard(board)
			n := throwSticks()
			var v []int
			if v = validMoves(token, n, board); len(v) == 0 {
				continue
			}
			m := getUserMove(token, pl1Token, v, board)
			if updateBoard(token, m, n, &board) {
				s := "You lost."
				if token == pl1Token {
					s = "You won!"
					s = "End"
				}
				fmt.Println(s)
				os.Exit(0)
			}
			if n == 2 || n == 3 {
				break
			}
		}
	}
}

func printBoard(board Board) {
	k := 0
	fmt.Println()
	fmt.Println("     0   1   2   3   4   5   6   7   8   9 ")
	fmt.Println("    19  18  17  16  15  14  13  12  11  10 ")
	fmt.Println("  +---+---+---+---+---+---+---+---+---+---+")
	for j := 0; j < 3; j++ {
		fmt.Print("  |")
		for i := 0; i < 10; i++ {
			k = j*10 + i
			if j == 1 {
				k = 19 - i
			}
			switch board[k] {
			case P:
				fmt.Printf(" A |")
				continue
			case S:
				fmt.Printf(" O |")
				continue
			default:
				switch k {
				case 14:
					fmt.Printf(" t |") // ankh - rebirth
				case 25:
					fmt.Printf(" k |") // bird - happiness
				case 26:
					fmt.Printf(" ~ |") // water - chaos
				case 27:
					fmt.Printf(" : |") // three truths
				case 28:
					fmt.Printf(" @ |") // eye - Re-Atum
				default:
					fmt.Printf("   |")
				}
			}
		}
		fmt.Println("\n  +---+---+---+---+---+---+---+---+---+---+")
	}
	fmt.Println("    20  21  22  23  24  25  26  27  28  29 ")
}

func throwSticks() (n int) {
	fmt.Print("Sticks: ")
	nn := [5]int{5, 1, 2, 3, 4}
	s := rand.NewSource(time.Now().UnixNano() + 2)
	for i := 0; i < 4; i++ {
		r := rand.New(s)
		x := r.Intn(2)
		fmt.Print(x, " ")
		n += x
	}
	fmt.Print("= ", nn[n], "\n")
	return nn[n]
}

func validMoves(tok int, n int, board Board) (v []int) { // add opt jumping rules, opt 3rd row clear to bear off
	for i := 0; i < 26; i++ {
		if (i + n) > 26 { // stop at bird
			continue
		}
		if board[i+n] == tok && board[i] == tok { // can't bump own piece
			continue
		}
		if board[i+n] == 0 && board[i] == tok { // empty
			v = append(v, i)
			continue
		}
		if (i+n+1) < 26 && board[i] == tok { // blot
			if board[i+n] == (tok^1) && board[i+n+1] != (tok^1) && board[i+n-1] != (tok^1) {
				v = append(v, i)
				continue
			}
		}
	}
	for i := 26; i < 27; i++ {
		if i == 26 && (i+n) < 31 && board[i] == tok {
			v = append(v, i)
		}
	}
	for i := 27; i < 30; i++ {
		if (i+n) == 31 && board[i] == tok {
			v = append(v, i)
		}
	}
	return v
}

func getUserMove(tok int, pTok int, v []int, board Board) (m int) {
	input := ""
	if tok == pTok {
		for {
			fmt.Print(" Moves:")
			for i := range v {
				fmt.Print(" ", v[i])
			}
			fmt.Print("\n Enter: ")
			fmt.Scanln(&input)
			m, _ := strconv.Atoi(input)
			for i := range v {
				if v[i] == m {
					return m
				}
			}
		}
	}
	return m
}

func getComputerMove(tok int, pTok int, v []int, board Board) (m int) {
	return 0
}

func updateBoard(tok int, m int, n int, board *Board) (over bool) {
	if board[m+n] != 0 {
		board[m] = board[m+n]
	} else {
		board[m] = 0
	}
	board[m+n] = tok

	//ADD waters
	//ADD bearoff

	for _, j := range board {
		if j == tok {
			return false
		}
	}
	return true
}
