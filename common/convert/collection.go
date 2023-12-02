package convert

import "golang.org/x/exp/constraints"

// GetMinValue 比较大小 如果集合为空则返回nil
func GetMinValue[T comparable](from []T, valueFunc func(T) *int32) *int32 {
	if len(from) == 0 {
		return nil
	}

	minValue := valueFunc(from[0])

	for _, item := range from {
		value := valueFunc(item)
		if value != nil && (*value) < (*minValue) {
			minValue = value
		}
	}
	return minValue
}

func Int32Sum(a, b int32) int32 {
	return a + b
}

// GetSumValue 是一个泛型函数，接收一个列表 from，一个映射函数 valueFunc 和一个累加器 accumulator。
// valueFunc 用于将 T 类型的元素转换为 V 类型的元素，accumulator 用于将 V 类型的元素累加起来。
func GetSumValue[T any, V comparable](from []T, valueFunc func(T) V, accumulator func(V, V) V) V {
	// 将列表中的元素通过 valueFunc 映射为 V 类型的元素，然后使用 accumulator 进行累加
	sum := valueFunc(from[0])
	for i := 1; i < len(from); i++ {
		sum = accumulator(sum, valueFunc(from[i]))
	}

	return sum
}

func SliceFiltrate[V any](collection []V, filtrate func(V, int) bool) []V {
	var result []V
	for i, v := range collection {
		if filtrate(v, i) {
			result = append(result, v)
		}
	}

	return result
}

func SliceUpdateElement[T any, R any](collection []T, iteratee func(T, int) R) []R {
	result := make([]R, len(collection))

	for i, t := range collection {
		result[i] = iteratee(t, i)
	}

	return result
}

func SliceUniq[T any, U comparable](collection []T, iteratee func(T) U) []T {
	result := make([]T, len(collection))

	seen := make(map[U]struct{}, len(collection))
	for _, item := range collection {
		key := iteratee(item)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
	}

	return result
}

func SliceGroupBy[T any, U comparable](collection []T, iteratee func(T) U) map[U][]T {
	result := map[U][]T{}

	for _, item := range collection {
		key := iteratee(item)

		result[key] = append(result[key], item)
	}
	return result
}

// CheckInSlice  check value in slice
func CheckInSlice[T constraints.Ordered](a T, s []T) bool {
	for _, val := range s {
		if a == val {
			return true
		}
	}
	return false
}

// DelOneInSlice  delete one element of slice  left->right
func DelOneInSlice[T constraints.Ordered](a T, old []T) (new []T) {
	for i, val := range old {
		if a == val {
			new = append(old[:i], old[i+1:]...)
			return
		}
	}
	return old
}

type MapBuilder[T, K, V comparable] struct {
	data      []T
	keyFunc   func(T) K
	valFunc   func(T) V
	mergeFunc func(V, V) V
}

func newMapBuilder[T, K, V comparable](from []T, keyFunc func(T) K, valueFunc func(T) V, mergeFunc func(V, V) V) *MapBuilder[T, K, V] {
	return &MapBuilder[T, K, V]{
		data:      from,
		keyFunc:   keyFunc,
		valFunc:   valueFunc,
		mergeFunc: mergeFunc,
	}
}

func (builder *MapBuilder[T, K, V]) Build() map[K]V {
	result := make(map[K]V)
	for _, item := range builder.data {
		key := builder.keyFunc(item)
		value := builder.valFunc(item)
		if builder.mergeFunc != nil {
			if existingValue, ok := result[key]; ok {
				result[key] = builder.mergeFunc(existingValue, value)
			} else {
				result[key] = value
			}
		} else {
			result[key] = value
		}
	}
	return result
}

// Map 从给定的切片和键提取函数构建映射。
func Map[T, K comparable](from []T, keyFunc func(T) K) map[K]T {
	if len(from) == 0 {
		return make(map[K]T)
	}
	return MapValues(from, keyFunc, func(item T) T { return item })
}

// MapValues 从给定的切片、键提取函数和值提取函数构建映射。
func MapValues[T, K, V comparable](from []T, keyFunc func(T) K, valueFunc func(T) V) map[K]V {
	if len(from) == 0 {
		return make(map[K]V)
	}
	return MapValuesMerge(from, keyFunc, valueFunc, func(v1, v2 V) V { return v1 })
}

// MapValuesMerge 从给定的切片、键提取函数、值提取函数和合并函数构建映射。
func MapValuesMerge[T, K, V comparable](from []T, keyFunc func(T) K, valueFunc func(T) V, mergeFunc func(V, V) V) map[K]V {
	if len(from) == 0 {
		return make(map[K]V)
	}
	return newMapBuilder[T, K, V](from, keyFunc, valueFunc, mergeFunc).Build()
}

func MultiMap[T, K comparable](from []T, keyFunc func(T) K) map[K][]T {
	multiMap := make(map[K][]T)
	for _, item := range from {
		key := keyFunc(item)
		multiMap[key] = append(multiMap[key], item)
	}
	return multiMap
}
