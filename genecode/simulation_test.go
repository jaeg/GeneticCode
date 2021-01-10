package genecode

import (
	"testing"
)

func Test_CanSolveToGetAProgram(t *testing.T) {
	tests := []CreatureTest{
		{InputRegister: map[int]int{1: 1, 2: 0}, ExpectedRegister: map[int]int{3: 0, 4: 1}},
	}

	simulation := &Simulation{MutationChance: 0.3}
	results, err := simulation.Solve(tests, 100, 100, -1)
	if results == nil {
		t.Error("Failed to solve")
	}
	if err != nil {
		t.Error("Error trying to solve")
	}
}

func Test_SolveCanFailDuetoGenerations(t *testing.T) {
	tests := []CreatureTest{
		{InputRegister: map[int]int{1: 1, 2: 0}, ExpectedRegister: map[int]int{3: 0, 4: 1}},
		{InputRegister: map[int]int{1: 1, 2: 2}, ExpectedRegister: map[int]int{3: 0, 4: 4}},
		{InputRegister: map[int]int{1: 1, 2: 3}, ExpectedRegister: map[int]int{3: 2, 4: 1}},
		{InputRegister: map[int]int{1: 1, 2: 4}, ExpectedRegister: map[int]int{3: 5, 4: 1}},
		{InputRegister: map[int]int{1: 1, 2: 5}, ExpectedRegister: map[int]int{3: 1, 4: 1}},
	}

	simulation := &Simulation{MutationChance: 0.3}
	results, err := simulation.Solve(tests, 100, 100, 1)

	if results != nil {
		t.Error("Test found a solution anyway.")
	}
	if err.Error() != "No solution found within generation max" {
		t.Error("Didn't get error about exceeding generation max.")
	}
}
