package genecode

import (
	"fmt"
	"testing"
)

//Function tests
func Test_set(t *testing.T) {
	genes := make([]Gene, 0)
	genes = append(genes, FunctionGene{Op: "set"})
	genes = append(genes, NumberGene{Value: 1})
	genes = append(genes, NumberGene{Value: 5})

	output, index := genes[0].Eval(genes, 0)
	fmt.Println(output, index)
	if index != 2 {
		t.Errorf("Did not end up at the right index")
	}

	if output != 5 {
		t.Errorf("Did not output correctly")
	}
}

func Test_Write_WithComputer(t *testing.T) {
	computer := CreateComputer()
	computer.AddGene(FunctionGene{Op: "set"})
	computer.AddGene(NumberGene{Value: 1})
	computer.AddGene(NumberGene{Value: 5})

	output, index := computer.genes[0].Eval(computer.genes, computer.index)
	if index != 2 {
		t.Errorf("Did not end up at the right index")
	}

	if output != 5 {
		t.Errorf("Did not output correctly")
	}

	if computer.register[1] != 5 {
		t.Errorf("Did not set output map correctly")
	}
}

func Test_Read(t *testing.T) {
	computer := CreateComputer()
	computer.register[1] = 1
	computer.AddGene(FunctionGene{Op: "read"})
	computer.AddGene(NumberGene{Value: 1})

	output, index := computer.genes[0].Eval(computer.genes, computer.index)
	if index != 1 {
		t.Errorf("Did not end up at the right index")
	}

	if output != computer.register[1] {
		t.Errorf("Did not output correctly")
	}

}

func Test_IfGeneReturnsCorrectIndexWhenEvaluatesTrue(t *testing.T) {
	genes := make([]Gene, 0)
	genes = append(genes, FunctionGene{Op: "if"})
	genes = append(genes, ComparatorGene{Op: "="})
	genes = append(genes, NumberGene{Value: 5})
	genes = append(genes, NumberGene{Value: 5})
	genes = append(genes, NumberGene{Value: 5})

	_, index := genes[0].Eval(genes, 0)
	if index != 3 {
		t.Errorf("Did not end up at the right index")
	}
}

func Test_NestedIfs(t *testing.T) {
	genes := make([]Gene, 0)
	genes = append(genes, FunctionGene{Op: "if"})
	genes = append(genes, ComparatorGene{Op: "="})
	genes = append(genes, NumberGene{Value: 4})
	genes = append(genes, NumberGene{Value: 5})
	genes = append(genes, FunctionGene{Op: "if"})
	genes = append(genes, ComparatorGene{Op: "="})
	genes = append(genes, NumberGene{Value: 5})
	genes = append(genes, NumberGene{Value: 5})
	genes = append(genes, NumberGene{Value: 1})
	genes = append(genes, FunctionGene{Op: "endif"})
	genes = append(genes, FunctionGene{Op: "endif"})
	_, index := genes[0].Eval(genes, 0)
	fmt.Println(index)
	if index != 10 {
		t.Errorf("Did not end up at the right index")
	}
}

func Test_NestedIf2(t *testing.T) {
	c := CreateComputer()
	c.AddGene(FunctionGene{Op: "if"})
	c.AddGene(ComparatorGene{Op: "="})
	c.AddGene(NumberGene{Value: 5})
	c.AddGene(NumberGene{Value: 5})
	c.AddGene(FunctionGene{Op: "set"})
	c.AddGene(NumberGene{Value: 1})
	c.AddGene(NumberGene{Value: 1})
	c.AddGene(FunctionGene{Op: "if"})
	c.AddGene(ComparatorGene{Op: "="})
	c.AddGene(NumberGene{Value: 5})
	c.AddGene(NumberGene{Value: 5})
	c.AddGene(FunctionGene{Op: "set"})
	c.AddGene(NumberGene{Value: 2})
	c.AddGene(NumberGene{Value: 1})
	c.AddGene(FunctionGene{Op: "endif"})
	c.AddGene(FunctionGene{Op: "endif"})
	c.AddGene(FunctionGene{Op: "set"})
	c.AddGene(NumberGene{Value: 2})
	c.AddGene(NumberGene{Value: 1})

	c.Run()
	fmt.Println(c.index)
	if c.register[1] != 1 {
		t.Errorf("Did not process outer if")
	}
	if c.register[2] != 1 {
		t.Errorf("Did not process nested if")
	}
	if c.register[3] == 1 {
		t.Errorf("Didn't run code after the if")
	}
}

func Test_NestedIfDontRun(t *testing.T) {
	c := CreateComputer()
	c.AddGene(FunctionGene{Op: "if"})
	c.AddGene(ComparatorGene{Op: "="})
	c.AddGene(NumberGene{Value: 3})
	c.AddGene(NumberGene{Value: 5})
	c.AddGene(FunctionGene{Op: "set"})
	c.AddGene(NumberGene{Value: 1})
	c.AddGene(NumberGene{Value: 1})
	c.AddGene(FunctionGene{Op: "if"})
	c.AddGene(ComparatorGene{Op: "="})
	c.AddGene(NumberGene{Value: 5})
	c.AddGene(NumberGene{Value: 5})
	c.AddGene(FunctionGene{Op: "set"})
	c.AddGene(NumberGene{Value: 2})
	c.AddGene(NumberGene{Value: 1})
	c.AddGene(FunctionGene{Op: "endif"})
	c.AddGene(FunctionGene{Op: "endif"})
	c.AddGene(FunctionGene{Op: "set"})
	c.AddGene(NumberGene{Value: 3})
	c.AddGene(NumberGene{Value: 1})

	c.Run()

	if c.register[1] == 1 {
		t.Errorf("Should not of processed outer if")
	}
	if c.register[2] == 1 {
		t.Errorf("Should not of processed nested if")
	}
	if c.register[3] != 1 {
		t.Errorf("Didn't run code after the if")
	}
}

func Test_NestedIfNestedDoesntProcess(t *testing.T) {
	c := CreateComputer()
	c.AddGene(FunctionGene{Op: "if"})
	c.AddGene(ComparatorGene{Op: "="})
	c.AddGene(NumberGene{Value: 5})
	c.AddGene(NumberGene{Value: 5})
	c.AddGene(FunctionGene{Op: "set"})
	c.AddGene(NumberGene{Value: 1})
	c.AddGene(NumberGene{Value: 1})
	c.AddGene(FunctionGene{Op: "if"})
	c.AddGene(ComparatorGene{Op: "="})
	c.AddGene(NumberGene{Value: 5})
	c.AddGene(NumberGene{Value: 4})
	c.AddGene(FunctionGene{Op: "set"})
	c.AddGene(NumberGene{Value: 2})
	c.AddGene(NumberGene{Value: 1})
	c.AddGene(FunctionGene{Op: "endif"})
	c.AddGene(FunctionGene{Op: "endif"})
	c.AddGene(FunctionGene{Op: "set"})
	c.AddGene(NumberGene{Value: 3})
	c.AddGene(NumberGene{Value: 1})
	c.Run()
	fmt.Println(c.register)
	if c.register[1] != 1 {
		t.Errorf("Should of processed outer if")
	}
	if c.register[2] == 1 {
		t.Errorf("Should not of processed nested if")
	}
	if c.register[3] != 1 {
		t.Errorf("Didn't run code after the if")
	}
}

func Test_IfGeneReturnsCorrectIndexWhenEvaluatesFalse(t *testing.T) {
	genes := make([]Gene, 0)
	genes = append(genes, FunctionGene{Op: "if"})
	genes = append(genes, ComparatorGene{Op: "="})
	genes = append(genes, NumberGene{Value: 5})
	genes = append(genes, NumberGene{Value: 6})
	genes = append(genes, NumberGene{Value: 5})
	genes = append(genes, FunctionGene{Op: "endif"})

	_, index := genes[0].Eval(genes, 0)
	if index != 5 {
		t.Errorf("Did not end up at the right index")
	}
}

func Test_NumberGeneEvaluatesItsValue(t *testing.T) {
	genes := make([]Gene, 0)
	genes = append(genes, NumberGene{Value: 5})

	output, _ := genes[0].Eval(genes, 0)
	if output != 5 {
		t.Errorf("Did not get the correct value")
	}
}

//Comparator Tests

func Test_ComparatorReturnZeroWhenNotEnougArguements(t *testing.T) {
	genes := make([]Gene, 0)
	genes = append(genes, ComparatorGene{Op: "="})
	genes = append(genes, NumberGene{Value: 5})

	output, index := genes[0].Eval(genes, 0)
	if index != len(genes) {
		t.Errorf("Did not end up at the right index")
	}
	if output != 0 {
		t.Errorf("Did not compare correctly")
	}
}

func Test_OperatorReturnZeroWhenNotEnougArguements(t *testing.T) {
	genes := make([]Gene, 0)
	genes = append(genes, OperatorGene{Op: "+"})
	genes = append(genes, NumberGene{Value: 5})

	output, index := genes[0].Eval(genes, 0)
	if index != len(genes) {
		t.Errorf("Did not end up at the right index")
	}
	if output != 0 {
		t.Errorf("Did not compare correctly")
	}
}

func Test_IfReturnZeroWhenNotEnougArguements(t *testing.T) {
	genes := make([]Gene, 0)
	genes = append(genes, FunctionGene{Op: "if"})

	output, index := genes[0].Eval(genes, 0)
	if index != len(genes) {
		t.Errorf("Did not end up at the right index")
	}
	if output != 0 {
		t.Errorf("Did not compare correctly")
	}
}

func Test_WriteReturns0(t *testing.T) {
	genes := make([]Gene, 0)
	genes = append(genes, FunctionGene{Op: "set"})
	genes = append(genes, NumberGene{Value: 5})

	output, index := genes[0].Eval(genes, 0)
	if index != len(genes) {
		t.Errorf("Did not end up at the right index")
	}
	if output != 0 {
		t.Errorf("Did not compare correctly")
	}
}

func Test_ReadReturns0(t *testing.T) {
	genes := make([]Gene, 0)
	genes = append(genes, FunctionGene{Op: "read"})

	output, index := genes[0].Eval(genes, 0)
	if index != len(genes) {
		t.Errorf("Did not end up at the right index")
	}
	if output != 0 {
		t.Errorf("Did not compare correctly")
	}
}
func Test_ComparatorEquals(t *testing.T) {
	genes := make([]Gene, 0)
	genes = append(genes, ComparatorGene{Op: "="})
	genes = append(genes, NumberGene{Value: 5})
	genes = append(genes, NumberGene{Value: 5})

	output, index := genes[0].Eval(genes, 0)
	if index != 2 {
		t.Errorf("Did not end up at the right index")
	}
	if output == 0 {
		t.Errorf("Did not compare correctly")
	}
}
func Test_ComparatorLessThan(t *testing.T) {
	genes := make([]Gene, 0)
	genes = append(genes, ComparatorGene{Op: "<"})
	genes = append(genes, NumberGene{Value: 5})
	genes = append(genes, NumberGene{Value: 6})

	output, index := genes[0].Eval(genes, 0)
	if index != 2 {
		t.Errorf("Did not end up at the right index")
	}
	if output == 0 {
		t.Errorf("Did not compare correctly")
	}
}

func Test_ComparatorGreaterThan(t *testing.T) {
	genes := make([]Gene, 0)
	genes = append(genes, ComparatorGene{Op: ">"})
	genes = append(genes, NumberGene{Value: 6})
	genes = append(genes, NumberGene{Value: 5})

	output, index := genes[0].Eval(genes, 0)
	if index != 2 {
		t.Errorf("Did not end up at the right index")
	}
	if output == 0 {
		t.Errorf("Did not compare correctly")
	}
}

func Test_ComparatorNot(t *testing.T) {
	genes := make([]Gene, 0)
	genes = append(genes, ComparatorGene{Op: "!"})
	genes = append(genes, NumberGene{Value: 5})
	genes = append(genes, NumberGene{Value: 6})

	output, index := genes[0].Eval(genes, 0)
	if index != 2 {
		t.Errorf("Did not end up at the right index")
	}
	if output == 0 {
		t.Errorf("Did not compare correctly")
	}
}

//Math tests
func Test_OperatorAdd(t *testing.T) {
	genes := make([]Gene, 0)
	genes = append(genes, OperatorGene{Op: "+"})
	genes = append(genes, NumberGene{Value: 1})
	genes = append(genes, NumberGene{Value: 1})

	output, index := genes[0].Eval(genes, 0)

	if index != 2 {
		t.Errorf("Did not end up at the right index")
	}
	if output != 2 {
		t.Errorf("Did not add correctly")
	}
}

func Test_OperatorMinus(t *testing.T) {
	genes := make([]Gene, 0)
	genes = append(genes, OperatorGene{Op: "-"})
	genes = append(genes, NumberGene{Value: 1})
	genes = append(genes, NumberGene{Value: 1})

	output, index := genes[0].Eval(genes, 0)

	if index != 2 {
		t.Errorf("Did not end up at the right index")
	}
	if output != 0 {
		t.Errorf("Did not subtract correctly")
	}
}

func Test_OperatorMultiply(t *testing.T) {
	genes := make([]Gene, 0)
	genes = append(genes, OperatorGene{Op: "*"})
	genes = append(genes, NumberGene{Value: 2})
	genes = append(genes, NumberGene{Value: 2})

	output, index := genes[0].Eval(genes, 0)

	if index != 2 {
		t.Errorf("Did not end up at the right index")
	}
	if output != 4 {
		t.Errorf("Did not multiply correctly")
	}
}

func Test_OperatorDivide(t *testing.T) {
	genes := make([]Gene, 0)
	genes = append(genes, OperatorGene{Op: "/"})
	genes = append(genes, NumberGene{Value: 4})
	genes = append(genes, NumberGene{Value: 2})

	output, index := genes[0].Eval(genes, 0)

	if index != 2 {
		t.Errorf("Did not end up at the right index")
	}
	if output != 2 {
		t.Errorf("Did not multiply correctly")
	}
}

//Gene factory tests
func Test_CreateNumberGene(t *testing.T) {
	geneString := "NumberGene 1"
	g, err := CreateGeneFromString(geneString)
	convertedGene, ok := g.(NumberGene)
	if err != nil {
		t.Errorf("Errored creating gene")
	}
	if !ok {
		t.Errorf("Did not create a number gene")
	}

	if convertedGene.Value != 1 {
		t.Errorf("Did not convert the gene value correctly")
	}
}

func Test_CreateNumberGeneError(t *testing.T) {
	geneString := "NumberGene f"
	_, err := CreateGeneFromString(geneString)
	if err == nil {
		t.Errorf("Did error creating bad number gene")
	}

}

func Test_CreateFunctionGene(t *testing.T) {
	geneString := "FunctionGene if"
	g, err := CreateGeneFromString(geneString)
	if err != nil {
		t.Errorf("Errored creating gene")
	}
	convertedGene, ok := g.(FunctionGene)
	if !ok {
		t.Errorf("Did not create a function gene")
	}
	if convertedGene.Op != "if" {
		t.Errorf("Did not convert the gene op correctly")
	}
}

func Test_CreateOperatorGene(t *testing.T) {
	geneString := "OperatorGene +"
	g, err := CreateGeneFromString(geneString)
	if err != nil {
		t.Errorf("Errored creating gene")
	}
	convertedGene, ok := g.(OperatorGene)
	if !ok {
		t.Errorf("Did not create an operator gene")
	}
	if convertedGene.Op != "+" {
		t.Errorf("Did not convert the gene op correctly")
	}
}

func Test_CreateComparatorGene(t *testing.T) {
	geneString := "ComparatorGene ="
	g, err := CreateGeneFromString(geneString)
	if err != nil {
		t.Errorf("Errored creating gene")
	}
	convertedGene, ok := g.(ComparatorGene)
	if !ok {
		t.Errorf("Did not create a comparator gene")
	}
	if convertedGene.Op != "=" {
		t.Errorf("Did not convert the gene op correctly")
	}
}

func Test_CreateGeneError(t *testing.T) {
	geneString := "BadGene asdf"
	_, err := CreateGeneFromString(geneString)
	if err == nil {
		t.Errorf("Didn't error when attempting to create a bad gene")
	}
}

func Test_GenerateRandomGene(t *testing.T) {
	for i := 0; i < 1000; i++ {
		gene := GenerateRandomGene()
		if len(gene) == 0 {
			t.Errorf("Didn't return a gene")
		}
	}
}
