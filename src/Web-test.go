package main

import (
	"cube/web"
)

func main() {
	app := web.GetApp("lin",8080)
	app.Start()
}
