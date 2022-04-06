package main

import (
	"opitz-consulting.com/example/container/internal/app/container"
)

func main() {
	app := container.Initialize()

	app.Run(":8081")
}
