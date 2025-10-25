package main

import "paolojulian.dev/inventory/interfaces/rest"

func main() {
	bootstrap := rest.Bootstrap()
	bootstrap.Router.Run("0.0.0.0:8080")

	bootstrap.DBCleanup()
}
