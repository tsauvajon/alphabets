package main

import "os"

func main() {
	app := App{}

	app.Initialize()

	port := os.Getenv("PORT")

	if port == "" {
		port = "3333"
	}

	app.Run(":" + port)
}
