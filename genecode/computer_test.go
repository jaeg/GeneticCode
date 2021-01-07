package genecode

import (
	"testing"
)

func Test_Computer_Run(t *testing.T) {
	computer := CreateComputer()
	computer.AddGene(FunctionGene{Op: "output"})
	computer.AddGene(NumberGene{Value: 1})
	computer.AddGene(NumberGene{Value: 5})

	computer.Run()
	if computer.index != 3 {
		t.Errorf("Did not end up at the right index")
	}

	if computer.outputs[1] != 5 {
		t.Errorf("Did not set output map correctly")
	}
}
