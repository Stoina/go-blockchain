package encoding

// ConvertToSlice exported
// ...
func ConvertToSlice(hashCode [32]byte) []byte {
	bytes := make([]byte, 32)
	copy(bytes, hashCode[:])
	return bytes
}
