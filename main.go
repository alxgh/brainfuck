package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("insert your text:")
		text, _ := reader.ReadString('\n')
		text = text[0 : len(text)-1]
		if text == "" {
			break
		}
		curr := int32(0)
		for _, c := range text {
			p := math.Abs(float64(curr - c))
			z := int32(math.Sqrt(math.Abs(p)))
			y := int32(p) - (z * z)
			q := '+'
			if c < curr {
				q = '-'
			}
			// fmt.Println("\n", curr, c, p, z, y, q)
			for i := int32(0); i < z; i++ {
				fmt.Print("+")
			}
			fmt.Print("[>")
			for i := int32(0); i < z; i++ {
				fmt.Printf("%c", q)
			}
			fmt.Print("<-]>")
			for i := int32(0); i < y; i++ {
				fmt.Printf("%c", q)
			}
			fmt.Print(".<")
			curr = c
		}
		fmt.Println()
	}
}
