package filterexercise

func MapFilterByKey[K comparable, V any](input map[K]V, filterFunc func(V) bool) map[K]V {
	output := make(map[K]V)
	for key, val := range input {
		if filterFunc(val) {
			output[key] = val
		}
	}

	return output
}

func MapFilterByVal[V any](input map[string]V, filterFunc func(V) bool) map[string]any {
	output := make(map[string]any)
	for key, val := range input {
		if filterFunc(val) {
			output[key] = val
		}
	}
	return output
}

func MapFilterByStruct[T any, V string](input []T, name V, filterFunc func(T, V) bool) bool {
	for _, val := range input {
		return filterFunc(val, name)
	}
	return false
}

func MapFilterByMap[T any](input map[string]T, filterFunc func(map[string]T) bool) bool {
	return filterFunc(input)
}

func MapFilterAny[V any](input map[string]V, filterFunc func(V) bool) map[string]V {
	output := make(map[string]V)
	for key, val := range input {
		if filterFunc(val) {
			output[key] = val
		}
	}
	return output
}
