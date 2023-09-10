package handler

import (
	"BikeCH/src/projection"
	"fmt"
	"net/http"
)

func Osm(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>OSM %v</h1>", projection.ContainsEN(projection.MAX_E, projection.MAX_N))
}
