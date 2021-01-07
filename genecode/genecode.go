package genecode

import "fmt"

type Gene interface {
	Eval(genes []Gene, index int) (int, int) // Returns value and the current index
}

type NumberGene struct {
	Value int
}

func (n NumberGene) Eval(genes []Gene, index int) (int, int) {
	return n.Value, index
}

type ComparatorGene struct {
	Op string
}

func (c ComparatorGene) Eval(genes []Gene, index int) (int, int) {
	output := 0
	val1 := 0
	val2 := 0
	val1, index = genes[index+1].Eval(genes, index+1)
	val2, index = genes[index+1].Eval(genes, index+1)
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

func (c FunctionGene) Eval(genes []Gene, index int) (int, int) {
	output := 0

	switch c.Op {
	//If gene is 1 express gene[index+1]
	case "if":
		val1 := 0
		val1, index = genes[index+1].Eval(genes, index+1)
		if val1 != 0 {
			index++
			output, index = genes[index].Eval(genes, index)
		}

		break
	case "output":
		index++
		output, index = genes[index].Eval(genes, index)
		fmt.Println("Output: ", output)

		break
	}

	return output, index
}

type OperatorGene struct {
	Op string
}

func (c OperatorGene) Eval(genes []Gene, index int) (int, int) {
	output := 0

	switch c.Op {
	case "+":
		val1 := 0
		val2 := 0
		val1, index = genes[index+1].Eval(genes, index+1)
		val2, index = genes[index+1].Eval(genes, index+1)
		output = val1 + val2
		break
	case "-":
		val1 := 0
		val2 := 0
		val1, index = genes[index+1].Eval(genes, index+1)
		val2, index = genes[index+1].Eval(genes, index+1)
		output = val1 - val2
		break
	case "/":
		val1 := 0
		val2 := 0
		val1, index = genes[index+1].Eval(genes, index+1)
		val2, index = genes[index+1].Eval(genes, index+1)
		if val2 != 0 {
			output = val1 / val2
		}

		break
	case "*":
		val1 := 0
		val2 := 0
		val1, index = genes[index+1].Eval(genes, index+1)
		val2, index = genes[index+1].Eval(genes, index+1)
		output = val1 * val2
		break
	}

	return output, index
}
