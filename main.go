package main

import (
	"github.com/jaeg/genecode/genecode"
)

func main() {

	computer := genecode.Computer{}

	/*
		if 1 < 2 { output 5 }
	*/
	computer.AddGene(genecode.FunctionGene{Op: "if"})
	computer.AddGene(genecode.ComparatorGene{Op: "<"})
	computer.AddGene(genecode.NumberGene{Value: 1})
	computer.AddGene(genecode.NumberGene{Value: 2})
	computer.AddGene(genecode.FunctionGene{Op: "output"})
	computer.AddGene(genecode.NumberGene{Value: 5})
	computer.AddGene(genecode.FunctionGene{Op: "endif"})

	computer.AddGene(genecode.FunctionGene{Op: "if"})
	computer.AddGene(genecode.ComparatorGene{Op: "<"})
	computer.AddGene(genecode.NumberGene{Value: 1})
	computer.AddGene(genecode.NumberGene{Value: 2})
	computer.AddGene(genecode.FunctionGene{Op: "output"})
	computer.AddGene(genecode.NumberGene{Value: 6})
	computer.AddGene(genecode.FunctionGene{Op: "endif"})

	computer.AddGene(genecode.FunctionGene{Op: "output"})
	computer.AddGene(genecode.NumberGene{Value: 99})
	computer.Run()

	//genes := make([]genecode.Gene, 0)

	/*
		if 1 + 2 == 3 {output 6}
	*/
	/*
		genes = append(genes, genecode.FunctionGene{Op: "if"})
		genes = append(genes, genecode.ComparatorGene{Op: "="})
		genes = append(genes, genecode.OperatorGene{Op: "+"})
		genes = append(genes, genecode.NumberGene{Value: 1})
		genes = append(genes, genecode.NumberGene{Value: 2})
		genes = append(genes, genecode.NumberGene{Value: 3})
		genes = append(genes, genecode.FunctionGene{Op: "output"})
		genes = append(genes, genecode.NumberGene{Value: 6})
		fmt.Println("2:")
		fmt.Println(genes[0].Eval(genes, 0))
	*/
}
