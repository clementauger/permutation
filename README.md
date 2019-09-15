# permute

Implements non recursive heap permutations algorithm.

## usage

```go
import "github.com/clementauger/permutation"

out := permutation.Generate([]interface{}{1,2,3})

for _, permuted := range out {
  fmt.Println(permuted) //permuted is []interface{}
}

```
