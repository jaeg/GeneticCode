package main

import (
	"fmt"
	"math/rand"
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

	const MaxGenes = 200
	const PopulationStart = 1000
	globalMaxFitness := 0
	tests := []genecode.CreatureTest{
		{InputRegister: map[int]int{1: 0}, ExpectedRegister: map[int]int{2: 1}},
		{InputRegister: map[int]int{1: 1}, ExpectedRegister: map[int]int{2: 2}},
		{InputRegister: map[int]int{1: 2}, ExpectedRegister: map[int]int{2: 3}},
		{InputRegister: map[int]int{1: 3}, ExpectedRegister: map[int]int{2: 0, 3: 1}},
	}
	population := []*genecode.Creature{}

	for i := 0; i < PopulationStart; i++ {
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
		fmt.Println("Generation ", currentGeneration)
		fmt.Println("Population: ", len(population))
		breedable := []*genecode.Creature{}
		//Calculate fitness for each creature
		for i := 0; i < len(population); i++ {
			fitness := population[i].CalculateFitness(tests)
			if fitness > maxFitness {
				maxFitness = fitness
			}

			if fitness > globalMaxFitness {
				globalMaxFitness = fitness
			}

			if fitness >= 2 {
				breedable = append(breedable, population[i])
			}

			if fitness == len(tests) {
				fmt.Println("Found the program!")
				fmt.Println("Generation: ", population[i].Generatation)
				fmt.Println(population[i])
				found = true
				break
			}
		}

		//Breed
		nextPopulation := []*genecode.Creature{}
		for i := 0; i < len(breedable); i++ {
			/*if len(breedable) == 1 {
				nextPopulation = append(nextPopulation, breedable[i]) //You you are breedable you at least go to the next generation
			}*/

			for j := i; j < len(breedable); j++ {
				nextPopulation = append(nextPopulation, breedable[i].BreedWith(breedable[j], 10)...)
			}
		}

		//Fill out with some random population
		for i := 0; i < 1000; i++ {
			c := &genecode.Creature{}
			for j := 0; j < MaxGenes; j++ {
				geneString := genecode.GenerateRandomGene()
				c.AddDNA(geneString)
			}
			nextPopulation = append(nextPopulation, c)
		}

		gap := PopulationStart - len(nextPopulation)
		if gap < 0 {
			//Randomly kill until we get to the population size.
			for i := gap; i < 0; i++ {
				s := rand.Intn(len(nextPopulation))
				nextPopulation = append(nextPopulation[:s], nextPopulation[s+1:]...)
			}
		}

		//Swap the generation
		population = nextPopulation
		fmt.Println("New Population: ", len(population))
		fmt.Println("Breedable Population: ", len(breedable))
		fmt.Println("Max Fitness: ", maxFitness)
		fmt.Println("Global Max Fitness: ", globalMaxFitness)
		fmt.Println("------------------")
	}
}
