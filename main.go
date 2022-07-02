package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
	"math/rand"
)

const Range = 50

func main() {
	mode := prompt("select mode (generate/draw)")

	list := make([]int, Range)
	for i := 0; i < Range; i++ {
		list[i] = i + 1
	}
	now := time.Now()
	rand.Seed(now.Unix())

	if mode == "draw" {
		prompt("press ENTER to draw next number")

		for {
			index := rand.Intn(len(list))
			fmt.Printf("%d", list[index])
			list = remove(list, index)
			if len(list) == 0 {
				break
			}
			prompt("")
		}

		fmt.Println("draw all numbers")
	} else if mode == "generate" {
		var n string
		for x := 0; x < 5; x++ {
			for y := 0; y < 5; y++ {
				if x == 2 && y == 2 {
					n = " B"
				} else {
					index := rand.Intn(len(list))
					n = fmt.Sprintf("%2d", list[index])
					list = remove(list, index)
				}
				fmt.Printf("%s ", n)
			}
			fmt.Printf("\n")
		}
	} else {
		fmt.Println("invalid mode, mode must be either of 'generate' or 'draw'")
	}
}

func remove(slice []int, s int) []int {
    return append(slice[:s], slice[s+1:]...)
}


func prompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stderr, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}
