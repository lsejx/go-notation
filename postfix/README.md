# go-notation/postfix
* Postfix-notation calculator library
* Postfix-notation calculator command (details is in ./cmd)
<br><br>

# Import
	import "github.com/lsejx/go-notation/postfix"
<br><br>

# Example
	// 1 2 + 3 -
	pfnc := postfix.NewCalculator[int](2)
	pfnc.AppendValue(1)
	pfnc.AppendValue(2)
	operationResult, err := pfnc.Operate(func(i1, i2 int) (int, error) { return i1 + i2, nil })
	// handle err
	// operationResult == 3

	pfnc.AppendValue(3)
	operationResult, err = pfnc.Operate(func(i1, i2 int) (int, error) { return i1 - i2, nil })
	// handle err
	// operationResult == 0

	result, err := pfnc.Result()
	// handle err
	// result == 0
