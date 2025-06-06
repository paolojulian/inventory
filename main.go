package main

import "paolojulian.dev/inventory/interfaces/rest"

func main() {
	bootstrap := rest.Bootstrap()
	bootstrap.Router.Run()

	bootstrap.DBCleanup()
}
