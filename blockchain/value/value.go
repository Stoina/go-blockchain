package value

// Value exported
// ...
type Value struct {
	Value float64
	Unit  string
}

// New exported
// ...
func New(value float64, unit string) *Value {
	return &Value{
		Value: value,
		Unit:  unit}
}
