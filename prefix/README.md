# go-notation/prefix
Prefix-notation calculator library
<br><br>

# Import
	import "github.com/lsejx/go-notation/prefix"
<br><br>

# Example (omit error handling)
	// - + 1 2 3
	prfc := prefix.NewCalculator[int]()

	err := prfc.AppendOperation(func(lhs, rhs int) (int, error) { return lhs - rhs, nil })
	err = prfc.AppendOperation(func(lhs, rhs int) (int, error) { return lhs + rhs, nil })
	err = prfc.AppendValue(1)
	err = prfc.AppendValue(2)
	err = prfc.AppendValue(3)

	res, err := prfc.Run()
	// res == 0
