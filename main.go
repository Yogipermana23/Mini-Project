package main

func main() {
	initConfig()
	initDb()
	router := InitRouter()
	router.Run(":8290")
}
