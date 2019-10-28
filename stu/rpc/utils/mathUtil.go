package utils

import "math"

type MathUtil struct {
}

func (mu *MathUtil) CircleArea(req float32, resp *float32) error {
	*resp = math.Pi * req * req
	return nil
}
