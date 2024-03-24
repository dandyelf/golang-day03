package main

import (
	"myHttp/db"
	"myJson/jsonWriter"
)

func main() {
	var db db.PlaceStore
	jsonWriter.HttpServ(&db)
}
