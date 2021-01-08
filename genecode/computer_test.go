package genecode

import (
	"fmt"
	"testing"
)

func Test_Computer_Run(t *testing.T) {
	computer := CreateComputer()
	computer.AddGene(FunctionGene{Op: "set"})
	computer.AddGene(NumberGene{Value: 1})
	computer.AddGene(NumberGene{Value: 5})

	computer.Run()
	fmt.Println(computer.register)
	if computer.index != 3 {
		t.Errorf("Did not end up at the right index")
	}

	if computer.register[1] != 5 {
		t.Errorf("Did not set output map correctly")
	}
}
