package main

import (
	"errors"
	"fmt"

	"github.com/amiltimore2016/gowebservice/models"
)

func main() {
	fmt.Println("Hello World")
	u := models.User{
		ID:        2,
		FirstName: "Tricia",
		LastName:  "Munster",
	}
	fmt.Println(u.FirstName, u.FirstName, u.ID)
	models.TestPointer()
	/*CarUsers := new(models.users[])*/
	port := 3000
	startWebServer(port, 2)
	_, status := startWebServer(port, 2)
	fmt.Println(status)
}

func startWebServer(port int, numRetries int) (int, error) {
	fmt.Println("Start retries", numRetries)
	fmt.Println("Starting Server on port ", port)
	// The cookie
	fmt.Println("Server Started", port)
	// Getting data back
	return port, errors.New("Something went wrong")
}
