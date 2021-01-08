package genecode

import (
	"fmt"
	"math/rand"
	"sort"
)

//Solve solves
func Solve(tests []CreatureTest, populationSize int, maxGenes int, maxGenerations int) *Creature {
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
		fmt.Println("Generation ", currentGeneration)
		fmt.Println("Population: ", len(population))

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
				fmt.Println("Found the program!")
				fmt.Println("Generation: ", population[i].Generatation)
				fmt.Println(population[i])
				return population[i]
			}
		}

		sort.Sort(ByFitness(population))

		nextPopulation := []*Creature{}
		// Top 10% go to the next generation automatically
		s := (10 * populationSize) / 100
		for i := 0; i < s; i++ {
			nextPopulation = append(nextPopulation, population[i])
		}

		// Top 50% go to the next generation automatically
		s = (45 * populationSize) / 100

		for i := 0; i < s; i++ {
			r := rand.Intn(populationSize / 2)
			//fmt.Println(r)
			child := population[i].BreedWith(population[r], 2)
			nextPopulation = append(nextPopulation, child...)
		}
		//Swap the generation
		population = nextPopulation
		fmt.Println("New Population: ", len(population))
		fmt.Println("Max Fitness: ", maxFitness)
		fmt.Println("Average Fitness: ", averageFitness/len(population))
		fmt.Println("Global Max Fitness: ", globalMaxFitness)
		fmt.Println("------------------")
	}

	return nil
}
