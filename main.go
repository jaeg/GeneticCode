package main

import (
	"fmt"

	"github.com/jaeg/genecode/genecode"
)

func main() {
	genes := make([]genecode.Gene, 0)

	/*
		if 1 < 2 { output 5 }
	*/
	genes = append(genes, genecode.FunctionGene{Op: "if"})
	genes = append(genes, genecode.ComparatorGene{Op: "<"})
	genes = append(genes, genecode.NumberGene{Value: 1})
	genes = append(genes, genecode.NumberGene{Value: 2})
	genes = append(genes, genecode.FunctionGene{Op: "output"})
	genes = append(genes, genecode.NumberGene{Value: 5})

	fmt.Println(genes[0].Eval(genes, 0))

	genes = make([]genecode.Gene, 0)

	/*
		if 1 + 2 == 3 {output 6}
	*/
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
}
