package main


const port = ":8080"

type Config struct{}

func main() {
	app := Config{}
	router := app.routes()
	router.Run(port)
}
