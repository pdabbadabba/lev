package terrain

import (
	"strconv"
	"core"
	"math"
	"image/png"
	"bytes"
 )

type Terrain struct {
	Id string
	HeightMap 	core.Resource
    MinX float32
    MinY float32
}

func (t Terrain) MarshalJSON() ([]byte, error){
//     return []byte(`{"minCoord":{"x":`+strconv.FormatFloat(float64(t.MinX), 'f', -1, 32)+`,"y":`+strconv.FormatFloat(float64(t.MinY), 'f', -1, 32)+`}, "heightMap":` + t.HeightMap.Id +`}`), nil
	xStr := strconv.FormatFloat(float64(t.MinX), 'f', -1, 32)
	yStr := strconv.FormatFloat(float64(t.MinY), 'f', -1, 32)
	jsonOut := `{"id":"`+t.Id+`", "mincoord":{"x":`+xStr+`, "y":`+yStr+`}, "heightmapid":"`+t.HeightMap.Id+`"}`;
	return []byte(jsonOut), nil	
}


func GetChunkCoords(x float32, y float32) (float32, float32) {
	chunkSize := float64(100)

	xRoot,_ := math.Modf(float64(x)/chunkSize)
	yRoot,_ := math.Modf(float64(y)/chunkSize)


	return float32(xRoot*chunkSize), float32(yRoot*chunkSize)
}

func GetTerrain(x float32, y float32) Terrain {

	x,y = GetChunkCoords(x,y)

	id := GetTerrainId(x,y)
	t := Terrain{id, GetHeightmap(x,y), x, y}
	return t
}



func GetHeightmap(x float32, y float32) core.Resource {
	buf := new(bytes.Buffer)

	png.Encode(buf, GetChunkImage(GetChunk(x,y,100)))
	return core.Resource{GetTerrainId(x,y), buf.Bytes()}
}

func GetTerrainId(x float32, y float32) string {
	return "terrain/"+strconv.FormatFloat(float64(x), 'f', -1, 32)+"/"+strconv.FormatFloat(float64(y), 'f', -1, 32)
}

