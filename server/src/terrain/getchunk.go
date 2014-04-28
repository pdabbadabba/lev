package terrain

import (
	"image"
	"image/color"
)

func  GetChunk(originx float32, originy float32, size int16) [][]float32 {

	chunk := make([][]float32, size)	

	for x := range chunk {
		chunk[x] = make([]float32, size)
		for y, _ := range chunk[x] {
			
			chunk[x][y] = GetHeight(float32(x)+originx, float32(y)+originy)
		}
	}

	return chunk
	
}

func MapFloat32HeightToGray(height float32) uint16 {
	maxh := float32(65536)

	return uint16((height * (maxh /2)) + (maxh/2))
}

func GetChunkImage(chunk [][]float32) *image.Gray16 {

	h := len(chunk)
	w := len(chunk[0])

	img := image.NewGray16(image.Rect(0,0,h,w))

	for x, r := range chunk {
		for y, v := range r {
			img.SetGray16(x, y, color.Gray16{MapFloat32HeightToGray(v)})
	
		}
	}

	return img
}