package main

import (
	"log"
	"net/http"
)

// "Hello from Snippetbox as the response body."
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))

}
func main() {
	// User the NewServeMux() (Router) function to initialize a new servemux, then
	// register the home function as the handler for the "URL" pattern
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	/*
		Use the http.ListenAndServe() function to start a new web server.
		We paas in two parameters: The TCP network address to listen on (in this case ":4000")
		and the servemux. If http.ListenAndServe() returns an error we use the log.Fatal() function to log
		the error message and exit. Any Error returned by http.ListenAndServe() is non-nil.
	*/
	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
