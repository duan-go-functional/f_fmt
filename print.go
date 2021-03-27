package f_fmt

import (
	"fmt"

	"github.com/duan-go-functional/base_type"
)

// Println对接了另一个IO流(os.Stdout)，不是纯函数，需要多一个返回值出口
func Println(res []base_type.Func) base_type.Func {
	return func(i ...interface{}) interface{} {
		n, err := fmt.Println(i...)
		if res != nil {
			res = []base_type.Func{
				base_type.Var(n),
				base_type.Var(err),
			}
		}
		return i
	}
}
