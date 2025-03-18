package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
)

// Handler function to serve the HTML page
func handler(w http.ResponseWriter, r *http.Request) {
	// Read the HTML file
	htmlContent, err := ioutil.ReadFile("index.html")
	if err != nil {
		http.Error(w, "Could not read HTML file", http.StatusInternalServerError)
		return
	}

	// Set content type to text/html and write the HTML content as the response
	w.Header().Set("Content-Type", "text/html")
	w.Write(htmlContent)
}

func main() {
	// Handle requests to the root URL by calling the handler
	http.HandleFunc("/", handler)

	// Print the URL where the server is running
	fmt.Println("Server is running at http://localhost:8080")

	// Start the web server
	log.Fatal(http.ListenAndServe(":8080", nil))
}




// package main

// import (
//     "fmt"
//     "log"
//     "net/http"
// )

// func handler(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintf(w, "Hi there, I love you %s!", r.URL.Path[1:])
// }

// func main() {
//     http.HandleFunc("/", handler)
// 	fmt.Println("http://localhost:8080")
//     log.Fatal(http.ListenAndServe(":8080", nil))
// }