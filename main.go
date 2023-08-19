package main

import "assignment-project/router"

func main() {
	var PORT = ":8080"

	router.StartServer().Run(PORT)
}
