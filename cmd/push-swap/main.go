package main

import (
	"fmt"
	"os"

	"push-swap/piles"
)

func main() {
	if len(os.Args) < 2 {
		return
	}

	nombres, err := piles.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error")
		os.Exit(1)
	}

	p := piles.New(nombres)
	for _, instruction := range trier(p) {
		fmt.Println(instruction)
	}
}

func positionDuMin(a []int) int {
	posMin := 0
	for i := 1; i < len(a); i++ {
		if a[i] < a[posMin] {
			posMin = i
		}
	}
	return posMin
}

func trier(p *piles.Piles) []string {
	var instructions []string

	faire := func(nom string) {
		p.Execute(nom)
		instructions = append(instructions, nom)
	}

	for !p.ATriee() {
		pos := positionDuMin(p.A)

		if pos <= len(p.A)/2 {
			for i := 0; i < pos; i++ {
				faire("ra")
			}
		} else {
			for i := 0; i < len(p.A)-pos; i++ {
				faire("rra")
			}
		}
		faire("pb")
	}

	for len(p.B) > 0 {
		faire("pa")
	}
	return instructions
}
