package convert

import "golang.org/x/exp/constraints"

func Signed[T constraints.Unsigned, K constraints.Signed](unsigned T) K {
	return K(unsigned)
}

func Unsigned[T constraints.Signed, K constraints.Unsigned](signed T) K {
	if signed < 0 {
		return 0
	}
	return K(signed)
}
