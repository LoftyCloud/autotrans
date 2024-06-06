package main

import (
	"autotrans/model"
	"autotrans/routes"
)

func main() {
	model.InitDb()
	routes.InitRouter()
}
