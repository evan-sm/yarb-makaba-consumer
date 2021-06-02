package main

func main() {
	go PullMsgsSync("yarb-313112", "yarb-makaba")
	r := setupRouter()
	r.Run(":" + ginPortMakabaConsumer)

}
