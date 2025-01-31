package convert

import (
	"strconv"

	"golang.org/x/exp/constraints"
)

func StringToInt(s string) (int, error) {
	return strconv.Atoi(s)
}

func IntSafeToUint(i int) (res uint, err error) {
	if i < 0 {
		return 0, strconv.ErrRange
	}
	return uint(i), nil
}

func StringSafeToUint(s string) (res uint, err error) {
	var i int
	if i, err = StringToInt(s); err != nil {
		return
	}
	return IntSafeToUint(i)
}

func Itoa[T constraints.Integer](i T) string {
	return strconv.Itoa(int(i))
}
