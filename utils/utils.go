package utils

func ValueToPointer[V any](v V) *V {
	return &v
}
