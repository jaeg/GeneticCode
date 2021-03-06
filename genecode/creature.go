package genecode

import (
	"math/rand"
)

//Creature represents a creature with dna that can run the computer
type Creature struct {
	dna          []string
	Fitness      int
	Generatation int
}

//CreatureTest a pair of input and expect outputs.
type CreatureTest struct {
	InputRegister    map[int]int
	ExpectedRegister map[int]int
}

//AddDNA adds the dna
func (c *Creature) AddDNA(dna string) {
	c.dna = append(c.dna, dna)
}

//BreedWith breed with other creature
func (c *Creature) BreedWith(mate *Creature, mutationChance float64, count int) []*Creature {
	creatures := make([]*Creature, 0)

	for i := 0; i < count; i++ {
		creature := &Creature{}
		creature.Generatation = c.Generatation + 1
		splitPoint := randomNumber(0, len(c.dna))
		//Parent one's genes
		dnaI := 0
		for ; dnaI < splitPoint; dnaI++ {
			if rand.Float64() < mutationChance {
				creature.AddDNA(GenerateRandomGene())
			} else {
				creature.AddDNA(c.dna[dnaI])
			}
		}

		for ; dnaI < len(mate.dna); dnaI++ {
			if rand.Float64() < mutationChance {
				creature.AddDNA(GenerateRandomGene())
			} else {
				creature.AddDNA(mate.dna[dnaI])
			}
		}

		creatures = append(creatures, creature)
	}

	return creatures
}

//BreedWithCol breed with other creature by combining the dna vs splitting it
func (c *Creature) BreedWithCol(mate *Creature, mutationChance float64, count int) []*Creature {
	creatures := make([]*Creature, 0)

	for i := 0; i < count; i++ {
		creature := &Creature{}
		creature.Generatation = c.Generatation + 1

		for dnaI := 0; dnaI < len(c.dna); dnaI++ {
			if rand.Float64() < mutationChance {
				creature.AddDNA(GenerateRandomGene())
			} else {
				if rand.Float64() < 0.5 {
					creature.AddDNA(c.dna[dnaI])
				} else {
					creature.AddDNA(mate.dna[dnaI])
				}
			}
		}

		creatures = append(creatures, creature)
	}

	return creatures
}

//CalculateFitness Calculates the fitness of the creature based on passed in tests.
func (c *Creature) CalculateFitness(tests []CreatureTest) int {
	comp := CreateComputer()
	for i := 0; i < len(c.dna); i++ {
		gene, err := CreateGeneFromString(c.dna[i])
		if err == nil {
			comp.AddGene(gene)
		}
	}

	passes := 0
	for i := 0; i < len(tests); i++ {
		passed := true
		comp.Clear()
		//Copy input register to this register
		for k, v := range tests[i].InputRegister {
			comp.register[k] = v
		}
		comp.Run()
		//As a rule, don't fail tests if it's using a register the expected isn't.
		for k, v := range tests[i].ExpectedRegister {
			if comp.register[k] != v {
				passed = false
			}
		}

		if passed == true {
			passes++
		}

	}
	c.Fitness = passes
	return passes
}

// ByFitness implements sort.Interface for []Creature based on
// the Fitness field.
type ByFitness []*Creature

func (a ByFitness) Len() int           { return len(a) }
func (a ByFitness) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFitness) Less(i, j int) bool { return a[i].Fitness > a[j].Fitness }
