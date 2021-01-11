package genecode

import (
	"fmt"
	"sort"
	"testing"
)

func Test_CreatureMate(t *testing.T) {
	c1 := &Creature{}
	for i := 0; i < 10; i++ {
		c1.dna = append(c1.dna, GenerateRandomGene())
	}

	fmt.Println("C1 ", c1.dna)

	c2 := &Creature{}
	for i := 0; i < 10; i++ {
		c2.dna = append(c2.dna, GenerateRandomGene())
	}

	fmt.Println("C2 ", c2.dna)

	child := c1.BreedWith(c2, 0.3, 1)
	fmt.Println("C3 ", child[0].dna)

	if len(child[0].dna) != 10 {
		t.Errorf("Child doesn't have enough genes")
	}
}

func Test_CreatureMateCol(t *testing.T) {
	c1 := &Creature{}
	for i := 0; i < 10; i++ {
		c1.dna = append(c1.dna, GenerateRandomGene())
	}

	fmt.Println("C1 COL", c1.dna)

	c2 := &Creature{}
	for i := 0; i < 10; i++ {
		c2.dna = append(c2.dna, GenerateRandomGene())
	}

	fmt.Println("C2 COL", c2.dna)

	child := c1.BreedWithCol(c2, 0.3, 1)
	fmt.Println("C3 COL", child[0].dna)

	if len(child[0].dna) != 10 {
		t.Errorf("Child doesn't have enough genes")
	}
}

func Test_CreatureFitness(t *testing.T) {
	c1 := &Creature{}
	c1.dna = append(c1.dna, "FunctionGene set")
	c1.dna = append(c1.dna, "NumberGene 2")
	c1.dna = append(c1.dna, "NumberGene 1")
	tests := []CreatureTest{
		{InputRegister: map[int]int{1: 1}, ExpectedRegister: map[int]int{2: 1}}, //Passes
		{InputRegister: map[int]int{1: 1}, ExpectedRegister: map[int]int{2: 0}}, //Fails
	}
	passes := c1.CalculateFitness(tests)

	if passes != 1 {
		t.Errorf("Not enough passes")
	}
}

func Test_CreatureSort(t *testing.T) {
	creatures := []*Creature{{Fitness: 1}, {Fitness: 5}, {Fitness: 2}}
	if creatures[0].Fitness == 5 {
		t.Errorf("Array was already sorted")
	}

	sort.Sort(ByFitness(creatures))

	if creatures[0].Fitness != 5 {
		t.Errorf("Creature array failed to sort.")
	}

}
