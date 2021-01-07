package main

import "fmt"

type Gene interface {
	eval(genes []Gene, index int) (int, int) // Returns value and the current index
}

type NumberGene struct {
	Value int
}

func (n NumberGene) eval(genes []Gene, index int) (int, int) {
	return n.Value, index
}

type ComparatorGene struct {
	Op string
}

func (c ComparatorGene) eval(genes []Gene, index int) (int, int) {
	output := 0
	val1 := 0
	val2 := 0
	val1, index = genes[index+1].eval(genes, index+1)
	val2, index = genes[index+1].eval(genes, index+1)
	switch c.Op {
	case "<":
		if val1 < val2 {
			output = 1
		}
		break
	case ">":
		if val1 > val2 {
			output = 1
		}
		break
	case "=":
		if val1 == val2 {
			output = 1
		}
		break
	case "!":
		if val1 != val2 {
			output = 1
		}
		break
	}

	return output, index
}

type FunctionGene struct {
	Op string
}

func (c FunctionGene) eval(genes []Gene, index int) (int, int) {
	output := 0

	switch c.Op {
	//If gene is 1 express gene[index+1]
	case "if":
		val1 := 0
		val1, index = genes[index+1].eval(genes, index+1)
		if val1 != 0 {
			index++
			output, index = genes[index].eval(genes, index)
		}

		break
	case "output":
		index++
		output, index = genes[index].eval(genes, index)
		fmt.Println("Output: ", output)

		break
	}

	return output, index
}

type OperatorGene struct {
	Op string
}

func (c OperatorGene) eval(genes []Gene, index int) (int, int) {
	output := 0

	switch c.Op {
	case "+":
		val1 := 0
		val2 := 0
		val1, index = genes[index+1].eval(genes, index+1)
		val2, index = genes[index+1].eval(genes, index+1)
		output = val1 + val2
		break
	case "-":
		val1 := 0
		val2 := 0
		val1, index = genes[index+1].eval(genes, index+1)
		val2, index = genes[index+1].eval(genes, index+1)
		output = val1 - val2
		break
	case "/":
		val1 := 0
		val2 := 0
		val1, index = genes[index+1].eval(genes, index+1)
		val2, index = genes[index+1].eval(genes, index+1)
		if val2 != 0 {
			output = val1 / val2
		}

		break
	case "*":
		val1 := 0
		val2 := 0
		val1, index = genes[index+1].eval(genes, index+1)
		val2, index = genes[index+1].eval(genes, index+1)
		output = val1 * val2
		break
	}

	return output, index
}

func main() {
	genes := make([]Gene, 0)

	/*
		if 1 < 2 { output 5 }
	*/
	genes = append(genes, FunctionGene{Op: "if"})
	genes = append(genes, ComparatorGene{Op: "<"})
	genes = append(genes, NumberGene{Value: 1})
	genes = append(genes, NumberGene{Value: 2})
	genes = append(genes, FunctionGene{Op: "output"})
	genes = append(genes, NumberGene{Value: 5})

	fmt.Println(genes[0].eval(genes, 0))

	genes = make([]Gene, 0)

	/*
		if 1 + 2 == 3 {output 6}
	*/
	genes = append(genes, FunctionGene{Op: "if"})
	genes = append(genes, ComparatorGene{Op: "="})
	genes = append(genes, OperatorGene{Op: "+"})
	genes = append(genes, NumberGene{Value: 1})
	genes = append(genes, NumberGene{Value: 2})
	genes = append(genes, NumberGene{Value: 3})
	genes = append(genes, FunctionGene{Op: "output"})
	genes = append(genes, NumberGene{Value: 6})
	fmt.Println("2:")
	fmt.Println(genes[0].eval(genes, 0))
}

/*
Code example
1.  if < 1 2  - evals as true
2.  if < + 1 2 1 - evals as false

1. if 1 < 2
2. if 1 + 2 < 1

Loop through genes.
Call eval on gene passing all the genes and the current index.
	If gene needs other genes, increment the index and call eval, increment index (which can change based on the next eval call) call eval.
	Do this until we run out of genes.
*/
