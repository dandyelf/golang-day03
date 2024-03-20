package main

import (
	"myHttp/db"
	"myHttp/htmlWriter"
)

func main() {
	var db db.PlaceStore
	htmlWriter.HttpServ(&db)
}
