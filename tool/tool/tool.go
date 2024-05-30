package tool

import (
	"fmt"
	"strings"
)

func ArrayToString[T any](array []T) string {
	return strings.Replace(strings.Trim(fmt.Sprint(array), "[]"), " ", ",", -1)
}
