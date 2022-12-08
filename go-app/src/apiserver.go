package main
// Importing the required packages
import (
	"fmt"
	"net/http"
	"log"
)
// Defining the messageHandler function to handle the requests
func messageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Response from Go server")
}
// Defining the main function to start the server
func main() {
	server := http.Server {
		Addr: ":8080",
	}
// Registering the messageHandler function to handle the requests on ip:8080/message path
	http.HandleFunc("/message", messageHandler)
	log.Println("Starting server on port 8080")
	server.ListenAndServe()
}
