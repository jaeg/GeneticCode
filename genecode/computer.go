package genecode

import "fmt"

//Computer represents a gene computer
type Computer struct {
	inputs  []int
	outputs []int
	flag    bool
	index   int

	genes []Gene
}

//AddGene adds gene to computer
func (c *Computer) AddGene(gene Gene) {
	switch v := gene.(type) {
	case FunctionGene:
		v.computerRef = c
		gene = v
		break
	}
	c.genes = append(c.genes, gene)
}

//Run run computer
func (c *Computer) Run() {
	c.index = 0

	for c.index < len(c.genes) {
		//output := 0
		_, c.index = c.genes[c.index].Eval(c.genes, c.index)
		//fmt.Println("Line Output:", output, c.index)
		c.index++
	}

	fmt.Println("Flag:", c.flag)
}
