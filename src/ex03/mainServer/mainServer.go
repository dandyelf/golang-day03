package mainServer

import (
	"log"
	"myGeoserv/hWriter"
	"myGeoserv/jWriter"
	"myHttp/types"
	"net/http"
)

type Store interface {
	// returns a list of items, a total number of hits and (or) an error in case of one
	GetPlaces(limit int, offset int) ([]types.Place, int, error)
}

var store Store

func HttpServ(st Store) {
	store = st
	mux := http.NewServeMux()
	mux.HandleFunc("/api/", jsonHandler)
	mux.HandleFunc("/", htmlHandler)
	log.Fatal(http.ListenAndServe(":8888", mux))
}

func jsonHandler(w http.ResponseWriter, r *http.Request) {
	jWriter.JWriter(w, r, store.GetPlaces)
}

func htmlHandler(w http.ResponseWriter, r *http.Request) {
	hWriter.HWriter(w, r, store.GetPlaces)
}
