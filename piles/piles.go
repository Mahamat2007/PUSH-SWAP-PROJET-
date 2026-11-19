package piles

import (
	"errors"
	"strconv"
	"strings"
)

type Piles struct {
	A []int
	B []int
}

func New(nombres []int) *Piles {
	return &Piles{A: nombres}
}
func Parse(args []string) ([]int, error) {
	var nombres []int
	vus := map[int]bool{}

	for _, arg := range args {
		for _, mot := range strings.Fields(arg) {
			n, err := strconv.Atoi(mot)
			if err != nil {
				return nil, errors.New("pas un entier")
			}
			if vus[n] {
				return nil, errors.New("doublon")
			}
			vus[n] = true
			nombres = append(nombres, n)
		}
	}
	return nombres, nil
}

func echanger(s []int) {
	if len(s) >= 2 {
		s[0], s[1] = s[1], s[0]
	}
}

func tourner(s []int) {
	if len(s) >= 2 {
		premier := s[0]
		copy(s, s[1:])
		s[len(s)-1] = premier
	}
}

func tournerEnArriere(s []int) {
	if len(s) >= 2 {
		dernier := s[len(s)-1]
		copy(s[1:], s[:len(s)-1])
		s[0] = dernier
	}
}

func (p *Piles) Pa() {
	if len(p.B) > 0 {
		p.A = append([]int{p.B[0]}, p.A...)
		p.B = p.B[1:]
	}
}

func (p *Piles) Pb() {
	if len(p.A) > 0 {
		p.B = append([]int{p.A[0]}, p.B...)
		p.A = p.A[1:]
	}
}

func (p *Piles) Execute(instruction string) error {
	switch instruction {
	case "pa":
		p.Pa()
	case "pb":
		p.Pb()
	case "sa":
		echanger(p.A)
	case "sb":
		echanger(p.B)
	case "ss":
		echanger(p.A)
		echanger(p.B)
	case "ra":
		tourner(p.A)
	case "rb":
		tourner(p.B)
	case "rr":
		tourner(p.A)
		tourner(p.B)
	case "rra":
		tournerEnArriere(p.A)
	case "rrb":
		tournerEnArriere(p.B)
	case "rrr":
		tournerEnArriere(p.A)
		tournerEnArriere(p.B)
	default:
		return errors.New("instruction inconnue")
	}
	return nil
}

func (p *Piles) ATriee() bool {
	for i := 0; i < len(p.A)-1; i++ {
		if p.A[i] > p.A[i+1] {
			return false
		}
	}
	return true
}

func (p *Piles) EstTriee() bool {
	return len(p.B) == 0 && p.ATriee()
}
