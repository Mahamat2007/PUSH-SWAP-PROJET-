package main

import (
	"bufio"
	"fmt"
	"os"

	"push-swap/piles"
)

func erreur() {
	fmt.Fprintln(os.Stderr, "Error")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		return
	}

	nombres, err := piles.Parse(os.Args[1:])
	if err != nil {
		erreur()
	}

	p := piles.New(nombres)

	lecteur := bufio.NewScanner(os.Stdin)
	for lecteur.Scan() {
		ligne := lecteur.Text()
		if ligne == "" {
			continue
		}
		if p.Execute(ligne) != nil {
			erreur()
		}
	}

	if p.EstTriee() {
		fmt.Println("OK")
	} else {
		fmt.Println("KO")
	}
}
