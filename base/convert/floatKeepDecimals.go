package convert

import "math"

type FloatKeepDecimals struct{}

func NewFloatKeepDecimals() *FloatKeepDecimals {
	return new(FloatKeepDecimals)
}
func (s FloatKeepDecimals) KeepDecimals(f float64) float64 {
	return math.Trunc((f)*1e2) / 1e2
}
