package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jaeg/genecode/genecode"
)

const MaxGenes = 100
const PopulationSize = 100

func main() {
	rand.Seed(time.Now().Unix())
	//Simulate
	tests := []genecode.CreatureTest{
		{InputRegister: map[int]int{1: 1, 2: 0}, ExpectedRegister: map[int]int{3: 0, 4: 1}},
		{InputRegister: map[int]int{1: 0, 2: 1}, ExpectedRegister: map[int]int{3: 1, 4: 0}},
		{InputRegister: map[int]int{1: 1, 2: 1}, ExpectedRegister: map[int]int{3: 1, 4: 0}},
		{InputRegister: map[int]int{1: 0, 2: 0}, ExpectedRegister: map[int]int{3: 1, 4: 1}},
	}

	fmt.Println(genecode.Solve(tests, 100, 100, -1))
}
