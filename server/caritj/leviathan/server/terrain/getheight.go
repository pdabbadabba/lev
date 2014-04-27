package terrain

import (
	"github.com/huin/chunkymonkey/perlin"
)

func GetHeight (x, y float32) float32 {
	p_scale := float64(0.1)
	noisectx := perlin.NewPerlinNoise(1234)
	noise := noisectx.At2d(float64(x) * p_scale, float64(y) * p_scale)
	return float32(noise)
	
}