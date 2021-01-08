package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"

	"github.com/jaeg/genecode/genecode"
)

func main() {
	rand.Seed(time.Now().Unix())
	/*
		computer := genecode.CreateComputer()

		/*
			if 1 < 2 { output 5 }
	*/
	/*
		computer.AddGene(genecode.FunctionGene{Op: "if"})
		computer.AddGene(genecode.ComparatorGene{Op: "<"})
		computer.AddGene(genecode.NumberGene{Value: 1})
		computer.AddGene(genecode.NumberGene{Value: 2})
		computer.AddGene(genecode.FunctionGene{Op: "set"})
		computer.AddGene(genecode.NumberGene{Value: 1})
		computer.AddGene(genecode.NumberGene{Value: 5})
		computer.AddGene(genecode.FunctionGene{Op: "endif"})

		computer.AddGene(genecode.FunctionGene{Op: "if"})
		computer.AddGene(genecode.ComparatorGene{Op: "<"})
		computer.AddGene(genecode.NumberGene{Value: 1})
		computer.AddGene(genecode.NumberGene{Value: 2})
		computer.AddGene(genecode.FunctionGene{Op: "set"})
		computer.AddGene(genecode.NumberGene{Value: 1})
		computer.AddGene(genecode.NumberGene{Value: 6})
		computer.AddGene(genecode.FunctionGene{Op: "endif"})

		computer.AddGene(genecode.FunctionGene{Op: "set"})
		computer.AddGene(genecode.NumberGene{Value: 1})
		computer.AddGene(genecode.NumberGene{Value: 99})
		computer.Run()

		fmt.Println("Register ", computer.ReadRegister(1))

		fmt.Println("- Random Computer -")
		computer = genecode.CreateComputer()
		for i := 0; i < 10000; i++ {
			geneString := genecode.GenerateRandomGene()

			gene, err := genecode.CreateGeneFromString(geneString)
			if err != nil {
				fmt.Println("Error with gene ", geneString)
				continue
			}
			computer.AddGene(gene)
		}
		computer.Run()*/

	//Simulate

	const MaxGenes = 100
	const PopulationSize = 100
	globalMaxFitness := 0
	tests := []genecode.CreatureTest{
		{InputRegister: map[int]int{1: 1, 2: 0}, ExpectedRegister: map[int]int{3: 0, 4: 1}},
		{InputRegister: map[int]int{1: 0, 2: 1}, ExpectedRegister: map[int]int{3: 1, 4: 0}},
		{InputRegister: map[int]int{1: 1, 2: 1}, ExpectedRegister: map[int]int{3: 0, 4: 0}},
		{InputRegister: map[int]int{1: 0, 2: 0}, ExpectedRegister: map[int]int{3: 1, 4: 1}},
	}
	population := []*genecode.Creature{}

	for i := 0; i < PopulationSize; i++ {
		c := &genecode.Creature{}
		for j := 0; j < MaxGenes; j++ {
			geneString := genecode.GenerateRandomGene()
			c.AddDNA(geneString)
		}
		population = append(population, c)
	}

	found := false
	for currentGeneration := 0; found == false; currentGeneration++ {
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
				found = true
				break
			}
		}

		sort.Sort(genecode.ByFitness(population))

		nextPopulation := []*genecode.Creature{}
		// Top 10% go to the next generation automatically
		s := (10 * PopulationSize) / 100
		for i := 0; i < s; i++ {
			nextPopulation = append(nextPopulation, population[i])
		}

		// Top 50% go to the next generation automatically
		s = (45 * PopulationSize) / 100

		for i := 0; i < s; i++ {
			r := rand.Intn(PopulationSize / 2)
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
}
