package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

const LimitX = 10
const LimitY = 15

func main() {
	board := make([][]bool, LimitX)
	for i := range board {
		board[i] = make([]bool, LimitY)
	}
	px, py := 0, 0
	vx, vy := 1, 1

	// cal & update the next ball position
	// draw the board into a []rune buffer
	// screen. MoveToTopLeft()
	// print the []rune buffer: string(buffer)
	var cell rune
	board[0][0] = true
	for {
		for col := range board[0] {
			for row := range board {
				cell = '🌑'
				if board[row][col] {
					cell = '🌕'
				}
				fmt.Print(string(cell))
				fmt.Print(" ")

			}
			fmt.Println("")
		}
		board[px][py] = false
		fmt.Printf("x:%d \ty:%d", px, py)
		px, vx = cal(px, vx, LimitX)
		py, vy = cal(py, vy, LimitY)
		screen.MoveTopLeft()
		time.Sleep(time.Second / 10)
		screen.Clear()
		board[px][py] = true
	}

	// for {
	// 	// fmt.Printf("[x:%d vx:%d], [y:%d vy:%d]\n", px, vx, py, vy)
	// 	for i := 0; i < py; i++ {
	// 		fmt.Println("")
	// 	}
	// 	line := make([]string, px)
	// 	line = append(line, "O")
	// 	fmt.Println("%s", line)
	// 	px, vx = cal(px, vx, LimitX)
	// 	py, vy = cal(py, vy, LimitY)
	// 	screen.Clear()
	// 	screen.MoveTopLeft()
	// 	time.Sleep(time.Second / 20)
	// }

}

func cal(px int, vx int, limit int) (int, int) {
	if vx == 1 {
		px = px + 1
		if px+1 >= limit {
			vx = -1
		}
	} else if vx == -1 {
		px = px - 1
		if px-1 < 0 {
			vx = 1
		}
	}
	return px, vx
}
