package main

import (
	"myHttp/db"
	"myHttp/htmlW"
)

func main() {
	var db db.PlaceStore
	htmlW.HttpServ(&db)
}
