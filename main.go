package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/amiltimore2016/gowebservice/controllers"
)

func main() {

	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}

func startWebServer(port int, numRetries int) (int, error) {
	fmt.Println("Start retries", numRetries)
	fmt.Println("Starting Server on port ", port)
	// The cookie
	fmt.Println("Server Started", port)
	// Getting data back
	return port, errors.New("Something went wrong")
}
