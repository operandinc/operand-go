package operand

// AsRef returns a reference to a value.
// Used to differentiate between an empty value and a null value.
func AsRef[T any](t T) *T {
	return &t
}
