package functions

import (
	"github.com/yakawa/simpleDB/common/result"
)

func funcAbs(args []interface{}) result.Value {
	if len(args) != 1 {
		return result.Value{}
	}
	switch args[0].(type) {
	case int:
		if args[0].(int) < 0 {
			return result.Value{Type: result.Integral, Integral: -1 * args[0].(int)}
		}
		return result.Value{Type: result.Integral, Integral: args[0].(int)}
	}
	return result.Value{}
}
