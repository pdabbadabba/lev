package math

import (
	"github.com/huin/chunkymonkey/perlin"
)

func PerlinRandomHeight(x uint16, y uint16) uint16 {
	maxh := float64(65536)

	noisectx := perlin.NewPerlinNoise(1234)
	noise := noisectx.At2d(float64(x) *0.1, float64(y) *0.1)
	return uint16((maxh/2) + ((maxh/2)*noise))
}