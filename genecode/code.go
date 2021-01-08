package genecode

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

//NumberMax sets the max number for random genes
var NumberMax = 10

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
	if index > len(genes)-3 {
		return 0, len(genes)
	}
	val1, index = genes[index+1].Eval(genes, index+1)
	if index > len(genes)-2 {
		return 0, len(genes)
	}
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
		if index > len(genes)-2 {
			return 0, len(genes)
		}
		val1, index = genes[index+1].Eval(genes, index+1)
		if val1 == 0 {
			depth := 0
		Loop:
			for index < len(genes) {
				switch v := genes[index].(type) {
				case FunctionGene:
					if v.Op == "if" {
						depth++
					}
					if v.Op == "endif" {
						if depth == 0 {
							break Loop
						} else {
							depth--
						}
					}
				}
				index++
			}
		}

		break
	case "output":
		if index > len(genes)-3 {
			return 0, len(genes)
		}
		index++
		key := 0
		key, index = genes[index].Eval(genes, index)
		index++
		if index > len(genes)-2 {
			return 0, len(genes)
		}
		output, index = genes[index].Eval(genes, index)
		fmt.Println("Output: ", key, ",", output)
		if c.computerRef != nil {
			c.computerRef.outputs[key] = output
		}
		break

	case "input":
		if index > len(genes)-2 {
			return 0, len(genes)
		}
		index++
		key := 0
		key, index = genes[index].Eval(genes, index)

		if c.computerRef != nil {
			output = c.computerRef.inputs[key]
		}
		fmt.Println("Input: ", key, ",", output)
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
	if index > len(genes)-3 {
		return 0, len(genes)
	}

	switch c.Op {
	case "+":
		val1 := 0
		val2 := 0
		val1, index = genes[index+1].Eval(genes, index+1)
		if index > len(genes)-2 {
			return 0, len(genes)
		}
		val2, index = genes[index+1].Eval(genes, index+1)
		output = val1 + val2
		break
	case "-":
		val1 := 0
		val2 := 0
		val1, index = genes[index+1].Eval(genes, index+1)
		if index > len(genes)-2 {
			return 0, len(genes)
		}
		val2, index = genes[index+1].Eval(genes, index+1)
		output = val1 - val2
		break
	case "/":
		val1 := 0
		val2 := 0
		val1, index = genes[index+1].Eval(genes, index+1)
		if index > len(genes)-2 {
			return 0, len(genes)
		}
		val2, index = genes[index+1].Eval(genes, index+1)
		if val2 != 0 {
			output = val1 / val2
		}
		break
	case "*":
		val1 := 0
		val2 := 0
		val1, index = genes[index+1].Eval(genes, index+1)
		if index > len(genes)-2 {
			return 0, len(genes)
		}
		val2, index = genes[index+1].Eval(genes, index+1)
		output = val1 * val2
		break
	}

	return output, index
}

//CreateGeneFromString creates a gene from a string
func CreateGeneFromString(geneString string) (gene Gene, returnErr error) {
	gene = nil
	geneChunks := strings.Split(geneString, " ")
	switch geneChunks[0] {
	case "NumberGene":
		value, err := strconv.Atoi(geneChunks[1])
		if err == nil {
			g := NumberGene{Value: value}
			gene = g
		} else {
			returnErr = errors.New("Invalid number value")
		}
	case "FunctionGene":
		g := FunctionGene{Op: geneChunks[1]}
		gene = g
	case "ComparatorGene":
		g := ComparatorGene{Op: geneChunks[1]}
		gene = g
	case "OperatorGene":
		g := OperatorGene{Op: geneChunks[1]}
		gene = g
	default:
		returnErr = errors.New("Invalid gene string")
	}

	return
}

//GenerateRandomGene Creates a randomized gene
func GenerateRandomGene() string {
	geneTypes := []string{"FunctionGene", "ComparatorGene", "OperatorGene", "NumberGene"}
	opTypes := []string{"+", "-", "*", "/"}
	comparatorTypes := []string{">", "<", "!", "="}
	functionTypes := []string{"if", "endif", "output", "input"}
	geneString := geneTypes[randomNumber(0, len(geneTypes))]

	switch geneString {
	case "NumberGene":
		geneString += " " + strconv.Itoa(randomNumber(0, NumberMax))
	case "FunctionGene":
		geneString += " " + functionTypes[randomNumber(0, len(functionTypes))]
	case "ComparatorGene":
		geneString += " " + comparatorTypes[randomNumber(0, len(comparatorTypes))]
	case "OperatorGene":
		geneString += " " + opTypes[randomNumber(0, len(opTypes))]
	}
	return geneString
}

func randomNumber(min int, max int) int {
	return rand.Intn(max-min) + min
}
