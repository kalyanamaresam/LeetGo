package two_sum

import "testing"

func TestTwoSum(t *testing.T) {
	tests := [][]int{
		{3, 2, 4},
		{0, 8, 7, 3, 3, 4, 2},
		{0, 1},
	}
	targets := []int{
		6,
		11,
		1,
	}
	results := [][]int{
		{1, 2},
		{1, 3},
		{0, 1},
	}
	caseNum := 3
	for i := 0; i < caseNum; i++ {
		if ret := twoSum(tests[i], targets[i]); ret[0] != results[i][0] && ret[1] != results[i][1] {
			t.Fatalf("case %d fails: %v\n", i, ret)
		}

	}

}
