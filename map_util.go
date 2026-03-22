package generic

import (
	"fmt"
	"strings"
)

func CopyMap[K comparable, V any](mp map[K]V) map[K]V {
	var ret = make(map[K]V, len(mp))
	for k := range mp {
		ret[k] = mp[k]
	}

	return ret
}

func AsItems[K comparable, V any](mp map[K]V) Items[K, V] {
	var ret = make([]Item[K, V], 0, len(mp))
	for k, v := range mp {
		ret = append(ret, Item[K, V]{
			Key:   k,
			Value: v,
		})
	}
	return ret
}

func AsMap[K comparable, V any](slice Items[K, V]) map[K]V {
	var ret = make(map[K]V, len(slice))
	for _, item := range slice {
		ret[item.Key] = item.Value
	}
	return ret
}

func (items Items[K, V]) Map() map[K]V {
	return AsMap(items)
}

func (items Items[K, V]) TopN(n int, less func(lhs, rhs Item[K, V]) bool) Items[K, V] {
	n = min(n, len(items))

	if n <= 0 {
		return nil
	}

	var sorted = SortedFunc(items, less)

	return sorted[:n]
}

func (item Item[K, V]) String() string {
	return fmt.Sprintf("%+v: %+v", item.Key, item.Value)
}

func (items Items[K, V]) String() string {
	var r = make([]string, len(items))
	for i, item := range items {
		r[i] = item.String()
	}
	return fmt.Sprintf("{%s}", strings.Join(r, ", "))
}
