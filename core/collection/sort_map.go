package collection

import (
	"cmp"
	"sort"
)

// SortedMapTraversal 排序后遍历
func SortedMapTraversal[T cmp.Ordered, R any](m map[T]R, reverse bool, function func(T, R)) {
	keys := make([]T, 0)
	for k := range m {
		keys = append(keys, k)
	}

	sort.Slice(keys, func(i, j int) bool {
		if reverse {
			return keys[i] > keys[j]
		}
		return keys[i] < keys[j]
	})

	for _, v := range keys {
		function(v, m[v])
	}
}
