package genecode

import (
	"testing"
)

//Function tests
func Test_Output(t *testing.T) {
	genes := make([]Gene, 0)
	genes = append(genes, FunctionGene{Op: "output"})
	genes = append(genes, NumberGene{Value: 5})

	output, index := genes[0].Eval(genes, 0)
	if index != 1 {
		t.Errorf("Did not end up at the right index")
	}

	if output != 5 {
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
