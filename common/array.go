package common

import "cmp"

func Partition[T any](arr []T, size int) [][]T {
	result := make([][]T, 0)
	length := len(arr)
	if length <= 0 {
		return result
	}
	outSize := length / size
	for i := 0; i <= outSize; i++ {
		innerArr := make([]T, 0)
		for j := i * size; j < (i+1)*size; j++ {
			if j < length {
				innerArr = append(innerArr, arr[j])
			}
		}
		if len(innerArr) > 0 {
			result = append(result, innerArr)
		}
	}
	return result
}

func GroupBy[T any, R cmp.Ordered](arr []T, function func(T) R) map[R][]T {
	result := map[R][]T{}
	for _, ele := range arr {
		k := function(ele)
		result[k] = append(result[k], ele)
	}
	return result
}
