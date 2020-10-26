package functions

import (
	"github.com/yakawa/simpleDB/common/result"
)

type callFunction func([]interface{}) result.Value

var funcs = map[string]callFunction{
	"ABS": funcAbs,
}

func LookupFunction(name string) callFunction {
	f, exists := funcs[name]
	if !exists {
		return nil
	}

	return f
}
