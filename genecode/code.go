package genecode

import "fmt"

// Gene base gene interface
type Gene interface {
	Eval(genes []Gene, index int) (int, int) // Returns value and the current index
}

// NumberGene Represents a static number
type NumberGene struct {
	Value int
}

//Eval evaluate the gene
func (n NumberGene) Eval(genes []Gene, index int) (int, int) {
	return n.Value, index
}

//ComparatorGene Represents a compare operation < > = !
type ComparatorGene struct {
	Op string
}

//Eval evaluate the gene
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

//FunctionGene Represents a function operation ( if )
type FunctionGene struct {
	Op          string
	computerRef *Computer
}

//Eval evaluate the gene
func (c FunctionGene) Eval(genes []Gene, index int) (int, int) {
	output := 0

	switch c.Op {
	//If gene is 1 express gene[index+1]
	case "if":
		val1 := 0
		val1, index = genes[index+1].Eval(genes, index+1)
		if val1 == 0 {
		Loop:
			for index < len(genes) {
				switch v := genes[index].(type) {
				case FunctionGene:
					if v.Op == "endif" {
						break Loop
					}
					break
				}

				index++
			}
		}

		break
	case "output":
		index++
		output, index = genes[index].Eval(genes, index)
		fmt.Println("Output: ", output)
		if c.computerRef != nil {
			c.computerRef.flag = true
		}

		break
	}

	return output, index
}

//OperatorGene Represents an operator operation.  +, -, /, *
type OperatorGene struct {
	Op string
}

//Eval evaluate the gene
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
