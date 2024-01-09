package main

import (
	"log"
	"os"
	"sipget/app"
)

func main() {
	application := app.Generate()

	if error := application.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}
