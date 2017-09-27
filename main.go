package main

import (
	"nana/conf"
	db "nana/flow"
)

func main() {
	defer db.SqlDB.Close()
	router := conf.InitRouter()
	router.Run(":8000")
}
