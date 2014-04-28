package main

import (
    "net/http"
    "terrain"
    "encoding/json"
    "strconv"
    "core"
    "strings"
)

type Manager struct {
    Resources map[string]core.Resource
}

func getTerrainHandler(manager Manager) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        handleTerrainRequest(w,r, manager)
    }
}

func handleTerrainRequest(w http.ResponseWriter, r *http.Request, manager Manager) {
    x, _ := strconv.ParseFloat(r.FormValue("x"), 32)
    y, _ := strconv.ParseFloat(r.FormValue("y"), 32)
    t := terrain.GetTerrain(float32(x),float32(y))

    encoder := json.NewEncoder(w)
    w.Header().Set("Content-Type: application/json", "*")
    w.Header().Set("Access-Control-Allow-Origin", "*")
    encoder.Encode(t)

    manager.Resources[t.Id] = t.HeightMap
}

func getResourceHandler(manager Manager) func(w http.ResponseWriter, r *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
        handleResourceRequest(w,r, manager)
    }
}

func handleResourceRequest(w http.ResponseWriter, r *http.Request, manager Manager) {

    id := strings.TrimLeft(r.URL.String(), "/resource/")
    w.Write(manager.Resources[id].Data)
}

func main() {

    manager := Manager{}
    manager.Resources = make(map[string]core.Resource)

    http.HandleFunc("/entity/terrain", getTerrainHandler(manager))
    http.HandleFunc("/resource/", getResourceHandler(manager))
    http.ListenAndServe(":8080", nil)
}