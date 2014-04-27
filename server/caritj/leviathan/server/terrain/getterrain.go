package terrain

func GetTerrain(x float32, y float32) [][]float32 {
	return GetChunk(x, y, 100)
}