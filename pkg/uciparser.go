package pkg

import (
	"fmt"
)

func Uget2string(s []string, ok bool) string {
	val := ""
	if ok {
		val = fmt.Sprint(s)
	}
	return val
}
