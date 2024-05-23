package helper

func ToPointer[k any](x k) *k {
	return &x
}
