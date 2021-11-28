package main

import (
	"maple"
	_ "maple/example/base/controller"
)

func main() {
	app := maple.NewApplication()
	app.Run(":8000")
}
