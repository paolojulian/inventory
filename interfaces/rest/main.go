package rest

func main() {
	bootstrap := Bootstrap()
	bootstrap.Router.Run()

	bootstrap.DBCleanup()
}
