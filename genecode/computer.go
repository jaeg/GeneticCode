package genecode

//Computer represents a gene computer
type Computer struct {
	register map[int]int
	index    int
	genes    []Gene
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
		//fmt.Println("Line Write:", output, c.index)
		c.index++
	}
}

//SetRegister set an input register for the computer
func (c *Computer) SetRegister(key int, value int) {
	c.register[key] = value
}

//ReadRegister gets value from output register
func (c *Computer) ReadRegister(key int) int {
	return c.register[key]
}

//CreateComputer creates an instance of a computer
func CreateComputer() *Computer {
	c := &Computer{}
	c.register = make(map[int]int)
	return c
}
