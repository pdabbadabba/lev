package perlin

import (
	"math/rand"
	"math"
)
 
const PI = 3.1415926535
 
type Vec2 struct {
	X, Y float32
}
 
func lerp(a, b, v float32) float32 {
	return a * (1 - v) + b * v
}
 
func smooth(v float32) float32 {
	return v * v * (3 - 2 * v)
}
 
func random_gradient(r *rand.Rand) Vec2 {
	v := r.Float64() * PI * 2
	return Vec2{
		float32(math.Cos(v)),
		float32(math.Sin(v)),
	}
}
 
func gradient(orig, grad, p Vec2) float32 {
	sp := Vec2{p.X - orig.X, p.Y - orig.Y}
	return grad.X * sp.X + grad.Y * sp.Y
}
 
type Noise2DContext struct {
	rgradients []Vec2
	permutations []int
	gradients [4]Vec2
	origins [4]Vec2
}
 
func NewNoise2DContext(seed int) *Noise2DContext {
	rnd := rand.New(rand.NewSource(int64(seed)))
 
	n2d := new(Noise2DContext)
	n2d.rgradients = make([]Vec2, 256)
	n2d.permutations = rand.Perm(256)
	for i := range n2d.rgradients {
		n2d.rgradients[i] = random_gradient(rnd)
	}
 
	return n2d
}
 
func (n2d *Noise2DContext) get_gradient(x, y int) Vec2 {
	idx := n2d.permutations[x & 255] + n2d.permutations[y & 255]
	return n2d.rgradients[idx & 255]
}
 
func (n2d *Noise2DContext) get_gradients(x, y float32) {
	x0f := math.Floor(float64(x))
	y0f := math.Floor(float64(y))
	x0 := int(x0f)
	y0 := int(y0f)
	x1 := x0 + 1
	y1 := y0 + 1
 
	n2d.gradients[0] = n2d.get_gradient(x0, y0)
	n2d.gradients[1] = n2d.get_gradient(x1, y0)
	n2d.gradients[2] = n2d.get_gradient(x0, y1)
	n2d.gradients[3] = n2d.get_gradient(x1, y1)
 
	n2d.origins[0] = Vec2{float32(x0f + 0.0), float32(y0f + 0.0)}
	n2d.origins[1] = Vec2{float32(x0f + 1.0), float32(y0f + 0.0)}
	n2d.origins[2] = Vec2{float32(x0f + 0.0), float32(y0f + 1.0)}
	n2d.origins[3] = Vec2{float32(x0f + 1.0), float32(y0f + 1.0)}
}
 
func (n2d *Noise2DContext) Get(x, y float32) float32 {
	p := Vec2{x, y}
	n2d.get_gradients(x, y)
	v0 := gradient(n2d.origins[0], n2d.gradients[0], p)
	v1 := gradient(n2d.origins[1], n2d.gradients[1], p)
	v2 := gradient(n2d.origins[2], n2d.gradients[2], p)
	v3 := gradient(n2d.origins[3], n2d.gradients[3], p)
	fx := smooth(x - n2d.origins[0].X)
	vx0 := lerp(v0, v1, fx)
	vx1 := lerp(v2, v3, fx)
	fy := smooth(y - n2d.origins[0].Y)
	return lerp(vx0, vx1, fy)
}