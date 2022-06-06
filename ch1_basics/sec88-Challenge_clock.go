package main

import (
	"fmt"
	"time"

	"github.com/inancgumus/screen"
)

func main() {
	type digit [5][3]string
	placeholder := digit{
		{" ", " ", " "},
		{" ", "█", " "},
		{" ", " ", " "},
		{" ", "█", " "},
		{" ", " ", " "},
	}
	placeholderBlank := digit{
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
		{" ", " ", " "},
	}
	num0 := digit{
		{"█", "█", "█"},
		{"█", " ", "█"},
		{"█", " ", "█"},
		{"█", " ", "█"},
		{"█", "█", "█"},
	}
	num1 := digit{
		{" ", "█", " "},
		{"█", "█", " "},
		{" ", "█", " "},
		{" ", "█", " "},
		{" ", "█", " "},
	}
	num2 := digit{
		{"█", "█", "█"},
		{"█", " ", "█"},
		{" ", "█", " "},
		{"█", " ", " "},
		{"█", "█", "█"},
	}
	num3 := digit{
		{"█", "█", "█"},
		{" ", " ", "█"},
		{" ", "█", "█"},
		{" ", " ", "█"},
		{"█", "█", "█"},
	}
	num4 := digit{
		{"█", " ", "█"},
		{"█", " ", "█"},
		{"█", " ", "█"},
		{"█", "█", "█"},
		{" ", " ", "█"},
	}
	num5 := digit{
		{"█", "█", "█"},
		{"█", " ", " "},
		{"█", "█", "█"},
		{" ", " ", "█"},
		{"█", "█", "█"},
	}
	num6 := digit{
		{"█", "█", "█"},
		{"█", " ", " "},
		{"█", "█", "█"},
		{"█", " ", "█"},
		{"█", "█", "█"},
	}
	num7 := digit{
		{"█", "█", "█"},
		{" ", " ", "█"},
		{" ", "█", " "},
		{"█", " ", " "},
		{"█", " ", " "},
	}
	num8 := digit{
		{"█", "█", "█"},
		{"█", " ", "█"},
		{"█", "█", "█"},
		{"█", " ", "█"},
		{"█", "█", "█"},
	}
	num9 := digit{
		{"█", "█", "█"},
		{"█", " ", "█"},
		{"█", "█", "█"},
		{" ", " ", "█"},
		{"█", "█", "█"},
	}
	digits := []digit{
		num0,
		num1,
		num2,
		num3,
		num4,
		num5,
		num6,
		num7,
		num8,
		num9,
		placeholder,
	}

	screen.Clear()
	screen.MoveTopLeft()
	// for _, num := range digits {
	// 	for _, j := range num {
	// 		fmt.Printf("%s%s%s", j[0], j[1], j[2])
	// 		fmt.Println()
	// 	}
	// 	fmt.Println()
	// }

	var num digit
	// for i := 0; i < 5; i++ {
	// 	for j := 0; j < len(digits); j++ {
	// 		num := digits[j]
	// 		fmt.Printf("%s%s%s ", num[i][0], num[i][1], num[i][2])
	// 		if j >= len(digits)-1 {
	// 			fmt.Println()
	// 		}
	// 	}
	// }

	var placeholderSlot digit
	for {
		// Moves the cursor to the top-left position of the screen
		screen.MoveTopLeft()

		// Animate the time always in the same position
		fmt.Println(time.Now())
		hour := time.Now().Hour()
		minute := time.Now().Minute()
		second := time.Now().Second()

		if second%2 != 0 {
			placeholderSlot = placeholder
		} else {
			placeholderSlot = placeholderBlank
		}
		clock := []digit{
			digits[hour/10],
			digits[hour%10],
			placeholderSlot,
			digits[minute/10],
			digits[minute%10],
			placeholderSlot,
			digits[second/10],
			digits[second%10],
		}
		for i := 0; i < 5; i++ {
			for j := 0; j < len(clock); j++ {
				num = clock[j]
				fmt.Printf("%s%s%s ", num[i][0], num[i][1], num[i][2])
				if j >= len(clock)-1 {
					fmt.Println()
				}
			}
		}

		time.Sleep(time.Second)
	}
}
