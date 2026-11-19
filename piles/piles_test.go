package piles

import "testing"

func TestParseDoublon(t *testing.T) {
	if _, err := Parse([]string{"1 2 1"}); err == nil {
		t.Error("un doublon devrait donner une erreur")
	}
}

func TestParseNonEntier(t *testing.T) {
	if _, err := Parse([]string{"0 one 2"}); err == nil {
		t.Error("un non-entier devrait donner une erreur")
	}
}

func TestSujet(t *testing.T) {
	p := New([]int{2, 1, 3, 6, 8, 5})
	for _, instr := range []string{"sa", "pb", "pb", "pb", "rb", "rrr", "pa", "pa", "pa"} {
		p.Execute(instr)
	}
	if !p.EstTriee() {
		t.Errorf("la pile devrait être triée : %v", p.A)
	}
}
