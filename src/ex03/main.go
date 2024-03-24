package main

import (
	"myGeoserv/mainServer"
	"myHttp/db"
)

func main() {
	var db db.PlaceStore
	mainServer.HttpServ(&db)
}
