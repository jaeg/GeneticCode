package main

import (
	"github.com/jaeg/genecode/genecode"
)

func main() {

	computer := genecode.CreateComputer()

	/*
		if 1 < 2 { output 5 }
	*/
	computer.AddGene(genecode.FunctionGene{Op: "if"})
	computer.AddGene(genecode.ComparatorGene{Op: "<"})
	computer.AddGene(genecode.NumberGene{Value: 1})
	computer.AddGene(genecode.NumberGene{Value: 2})
	computer.AddGene(genecode.FunctionGene{Op: "output"})
	computer.AddGene(genecode.NumberGene{Value: 1})
	computer.AddGene(genecode.NumberGene{Value: 5})
	computer.AddGene(genecode.FunctionGene{Op: "endif"})

	computer.AddGene(genecode.FunctionGene{Op: "if"})
	computer.AddGene(genecode.ComparatorGene{Op: "<"})
	computer.AddGene(genecode.NumberGene{Value: 1})
	computer.AddGene(genecode.NumberGene{Value: 2})
	computer.AddGene(genecode.FunctionGene{Op: "output"})
	computer.AddGene(genecode.NumberGene{Value: 1})
	computer.AddGene(genecode.NumberGene{Value: 6})
	computer.AddGene(genecode.FunctionGene{Op: "endif"})

	computer.AddGene(genecode.FunctionGene{Op: "output"})
	computer.AddGene(genecode.NumberGene{Value: 1})
	computer.AddGene(genecode.NumberGene{Value: 99})
	computer.Run()
}
