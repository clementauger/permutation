package permutation_test

import (
	"reflect"
	"testing"

	"github.com/clementauger/permutation"
)

type permutationsTest struct {
	input     []interface{}
	expect    [][]interface{}
	expectLen int
}

func TestPermutation(t *testing.T) {

	table := []permutationsTest{
		permutationsTest{
			input: []interface{}{"a"},
			expect: [][]interface{}{
				[]interface{}{"a"},
			},
			expectLen: 1,
		},
		permutationsTest{
			input: []interface{}{"a", "b"},
			expect: [][]interface{}{
				[]interface{}{"a", "b"},
				[]interface{}{"b", "a"},
			},
			expectLen: 2,
		},
		permutationsTest{
			input: []interface{}{1, 2},
			expect: [][]interface{}{
				[]interface{}{1, 2},
				[]interface{}{2, 1},
			},
			expectLen: 2,
		},
		permutationsTest{
			input: []interface{}{"a", "b", "c"},
			expect: [][]interface{}{
				[]interface{}{"a", "b", "c"},
				[]interface{}{"b", "a", "c"},
				[]interface{}{"c", "a", "b"},
				[]interface{}{"a", "c", "b"},
				[]interface{}{"b", "c", "a"},
				[]interface{}{"c", "b", "a"},
			},
			expectLen: 6,
		},
	}

	for _, test := range table {
		res := permutation.Generate(test.input)
		if got := len(res); got != test.expectLen {
			t.Fatalf("invalid result len %v, wanted %v for %v, res was %v", got, test.expectLen, test.input, res)
		}
		if !reflect.DeepEqual(res, test.expect) {
			t.Fatalf("invalid result %v, wanted %v", res, test.expect)
		}
	}

}
