package genecode

import (
	"errors"
	"fmt"
	"math/rand"
	"sort"
)

//Simulation represents the virtual word to evolve programs in.
type Simulation struct {
	MutationChance float64
	Verbose        bool
}

//Solve solves
func (s *Simulation) Solve(tests []CreatureTest, populationSize int, maxGenes int, maxGenerations int) (*Creature, error) {
	globalMaxFitness := 0
	population := []*Creature{}

	for i := 0; i < populationSize; i++ {
		c := &Creature{}
		for j := 0; j < maxGenes; j++ {
			geneString := GenerateRandomGene()
			c.AddDNA(geneString)
		}
		population = append(population, c)
	}

	currentGeneration := 0
	for {
		currentGeneration++
		if currentGeneration > maxGenerations && maxGenerations > 0 {
			break
		}
		maxFitness := 0
		averageFitness := 0
		if s.Verbose {
			fmt.Println("Generation ", currentGeneration)
			fmt.Println("Population: ", len(population))
		}

		//Calculate fitness for each creature
		for i := 0; i < len(population); i++ {
			fitness := population[i].CalculateFitness(tests)
			averageFitness += fitness
			if fitness > maxFitness {
				maxFitness = fitness
			}

			if fitness > globalMaxFitness {
				globalMaxFitness = fitness
			}

			if fitness == len(tests) {
				if s.Verbose {
					fmt.Println("Found the program!")
					fmt.Println("Generation: ", population[i].Generatation)
					fmt.Println(population[i])
				}
				return population[i], nil
			}
		}

		sort.Sort(ByFitness(population))

		nextPopulation := []*Creature{}
		// Top 10% go to the next generation automatically and are guarenteed to breed at least once.
		elite := (10 * populationSize) / 100
		for i := 0; i < elite; i++ {
			nextPopulation = append(nextPopulation, population[i])
			r := rand.Intn(elite)
			child := population[i].BreedWithCol(population[r], s.MutationChance, 1)
			nextPopulation = append(nextPopulation, child...)
		}

		// Top 50% go to the next generation automatically
		elite = (80 * populationSize) / 100

		for i := 0; i < elite; i++ {
			r := rand.Intn(populationSize / 2)
			child := population[i].BreedWithCol(population[r], s.MutationChance, 1)
			nextPopulation = append(nextPopulation, child...)
		}
		//Swap the generation
		population = nextPopulation
		if s.Verbose {
			fmt.Println("New Population: ", len(population))
			fmt.Println("Max Fitness: ", maxFitness)
			fmt.Println("Average Fitness: ", averageFitness/len(population))
			fmt.Println("Global Max Fitness: ", globalMaxFitness)
			fmt.Println("------------------")
		}
	}

	return nil, errors.New("No solution found within generation max")
}
