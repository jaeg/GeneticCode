package genecode

//Computer represents a gene computer
type Computer struct {
	inputs  map[int]int
	outputs map[int]int
	index   int
	genes   []Gene
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
}

//SetInput set an input register for the computer
func (c *Computer) SetInput(key int, value int) {
	c.inputs[key] = value
}

//GetOutput gets value from output register
func (c *Computer) GetOutput(key int) int {
	return c.outputs[key]
}

//CreateComputer creates an instance of a computer
func CreateComputer() *Computer {
	c := &Computer{}
	c.inputs = make(map[int]int)
	c.outputs = make(map[int]int)
	return c
}
