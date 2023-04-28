# go-notation/postfix
* Postfix-notation calculator library
* Postfix-notation calculator command (details are in ./cmd)
<br><br>

# Import
	import "github.com/lsejx/go-notation/postfix"
<br><br>

# Example (omit error handling)
	// 1 2 + 3 -
	pofc := postfix.NewCalculator[int](2)

	pofc.AppendValue(1)
	pofc.AppendValue(2)

	operationResult, err := pofc.Operate(func(i1, i2 int) (int, error) { return i1 + i2, nil })
	// operationResult == 3

	pofc.AppendValue(3)

	operationResult, err = pofc.Operate(func(i1, i2 int) (int, error) { return i1 - i2, nil })
	// operationResult == 0

	result, err := pofc.Result()
	// result == 0
