package main 

import (
	"fmt"
	"github.com/caritj/leviathan/server/math/perlin"
)

func main() {
	n := perlin.NewNoise2DContext(1234)
	for i := 0.0; i < 100; i++ {
		fmt.Println(uint16((65536/2) + (65536* n.Get(float32(i)*0.1, float32(i)*0.1))))
	}
}