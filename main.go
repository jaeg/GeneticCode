package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jaeg/genecode/genecode"
)

func main() {
	rand.Seed(time.Now().Unix())
	computer := genecode.CreateComputer()

	/*
		if 1 < 2 { output 5 }
	*/
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
	computer.Run()

}
