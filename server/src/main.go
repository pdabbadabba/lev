package main

import (
    "net/http"
    "github.com/caritj/leviathan/server/terrain"
    "image/png"
    "encoding/json"
    "strconv"
)

type Resource struct {
    id string
}





func terrainhandler(w http.ResponseWriter, r *http.Request)
{
    x, _ := strconv.ParseFloat(r.FormValue("x"), 32)
    y, _ := strconv.ParseFloat(r.FormValue("x"), 32)
    terrain.GetTerrain(float32(x),float32(y))

    encoder := json.NewEncoder(w)
    w.Header().Set("Access-Control-Allow-Origin", "*")
    encoder.Encode(t)

}

// func pnghandler(w http.ResponseWriter, r *http.Request) {

//     x, _ := strconv.ParseFloat(r.FormValue("x"), 32)
//     y, _ := strconv.ParseFloat(r.FormValue("x"), 32)
//     t := terrain.GetChunkImage(terrain.GetTerrain(float32(x),float32(y)))
//     w.Header().Set("Access-Control-Allow-Origin", "*")
//     png.Encode(w, t)
// }

func main() {



    http.HandleFunc("/entity/terrain", terrainhandler)
    http.ListenAndServe(":8080", nil)
}