// Copyright © 2017 Mark F. Pleso. All rights reserved.
// Use of this source code is governed by the GPL-2 license described in the
// LICENSE file.

// senet board game

package main

import (
      "fmt"
      "math/rand"
      "os"
      "strconv"
      "time"
)

func main() {
    cell := [30]int{1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0, 1, 0} // 7
    for i := 14; i < 30; i++ {
        cell[i] = -1
    }
    pl1Token := 0                      // user is pyramids

    for token := 0;; token ^= 1  {     // pyramids(0) always go first
        for {
            printBoard(cell)
            n := throwSticks()
            var v []int
            if v = validMoves(token, n, cell); len(v) == 0 {
                continue
            }
            m := getMove(token, pl1Token, v, cell)
            if updateBoard(token, m, n, cell) {
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

func printBoard(cell [30]int) {
    k := 0
    fmt.Println()
    fmt.Println("     1   2   3   4   5   6   7   8   9  10 ")
    fmt.Println("    20  19  18  17  16  15  14  13  12  11 ")
    fmt.Println("  +---+---+---+---+---+---+---+---+---+---+")
    for j := 0; j < 3; j++ {
        fmt.Print("  |")
        for i := 0; i<10; i++ {
            k = j * 10 + i
            if j == 1 {
                k = 19 - i
            }
            switch cell[k] {
            case 0:
                fmt.Printf(" A |")
                continue
            case 1:
                fmt.Printf(" 8 |")
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
                    fmt.Printf(" q |") // eye - Re-Atum
                default:
                    fmt.Printf("   |")
                }
            }
        }
        fmt.Println("\n  +---+---+---+---+---+---+---+---+---+---+")
    }
    fmt.Println("    21  22  23  24  25  26  27  28  29  30 ")
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

func validMoves(tok int, n int, cell [30]int) (v []int) { // add opt jumping rules, opt 3rd row clear
    for i := 0; i < 26; i++ {
        if cell[i] == tok {
            if (i + n) > 26 { // must stop at bird
		continue
            }
            if cell[i + n] == tok { // can't bump own piece
                continue
            }
            if cell[i + n] == -1 { // empty
                v = append(v, i)
                continue
            }
            if (i + n + 1) < 26 { // must be a blot
                if cell [i + n] == (tok ^ 1) && cell[i + n + 1] != (tok ^ 1)  && cell[i + n -1] != (tok ^ 1) {
                    v = append(v, i)
                    continue
                }
            }
            //add 26 case
        }
    }
    for i := 26; i < 30; i++ { //FIXME
    }
    return v
}

func getMove(tok int, pTok int, v []int, cell [30]int) (m int) {
    input := ""
    if tok == pTok {
        for {
            fmt.Print(" Moves:")
            for i := range v {
                fmt.Print(" ", v[i] + 1)
            }
            fmt.Print("\n Enter: ")
            fmt.Scanln(&input)
            m, _ := strconv.Atoi(input)
            m--
            for i := range v {
                if v[i] == m {
                  return m
                }
            }
        }
    }

    //FIXME computer move
    return m
}

func updateBoard(tok int, m int, n int, cell [30]int) (over bool) { //include swap, include waters, include bear off
//FIXME

    c := 0
    for _, j := range cell {
        if j == tok {
            c++
        }
    }
    c = 0 //TEMP
    if c == 0 {
        return true
    }
    return false
}
