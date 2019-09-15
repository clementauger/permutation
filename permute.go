package permutation

// Generate all permutations of data.
func Generate(data []interface{}) [][]interface{} {
	n := len(data)
	if n < 1 {
		return nil
	}
	if n == 1 {
		t := make([]interface{}, n)
		copy(t, data)
		return [][]interface{}{t}
	}

	backtrack := make([]int, n)

	f := factorial(n)
	out := make([][]interface{}, f)

	e := 0
	t := make([]interface{}, n)
	copy(t, data)
	out[e] = t
	e++

	i := 0
	for i < n {
		if backtrack[i] < i {
			if i%2 == 0 {
				data[0], data[i] = data[i], data[0]
			} else {
				data[backtrack[i]], data[i] = data[i], data[backtrack[i]]
			}

			t := make([]interface{}, n)
			copy(t, data)
			out[e] = t
			e++

			backtrack[i]++
			i = 0
		} else {
			backtrack[i] = 0
			i++
		}
	}
	return out
}

func factorial(n int) (factVal uint64) {
	if n > 0 {
		factVal = 1
		for i := 1; i <= n; i++ {
			factVal *= uint64(i)
		}
	}
	return
}
