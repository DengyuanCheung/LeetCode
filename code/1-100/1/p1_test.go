package p1

import (
	"reflect"
	"testing"
)
func testTwoSum(t *testing.T){
	type test struct {
		input1 []int
		input2 int
		want []int
	}

	tests := map[string]test{
		{input1: []int{2, 7, 11, 15}, input2: 9, want []int{0, 1}},
	}

	for _, v := range tests {
		result = twoSum(v.input1, v.input2)
		if !reflect.DeepEqual(v.want, result) {
			t.Errorf("want: %v, result: %v", v.want, result)
		}
	}


}