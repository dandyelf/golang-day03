package main

import (
	"myGeoserv/db"
	"myGeoserv/mServer"
)

func main() {
	var db db.PlaceStore
	mServer.HttpServ(&db)
}
