package main

import (
	"fmt"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to text/html
	w.Header().Set("Content-Type", "text/html")

	// Get the current time
	currentTime := time.Now().Format("Mon Jan 2 15:04:05 2006")

	// HTML content with the specified message and CSS styling
	htmlContent := `
		<!DOCTYPE html>
		<html>
		<head>
			<title>Hello Gig Group!</title>
			<style>
				body {
					font-family: Arial, sans-serif;
					background-color: #f5f5f5;
					text-align: center;
				}
				.container {
					margin-top: 50px;
				}
				.message {
					font-size: 36px;
					color: #333;
				}
				.timestamp {
					font-size: 24px;
					color: #666;
				}
			</style>
		</head>
		<body>
			<div class="container">
				<div class="message">Hello Gig Group! This confirms that task#3 is complete!</div>
				<div class="timestamp">Current Server Time: <strong>` + currentTime + `</strong></div>
			</div>
		</body>
		</html>
	`

	// Write the HTML content to the response
	fmt.Fprintf(w, htmlContent)
}

func main() {
	// Handle requests to the root URL ("/") with the handler function
	http.HandleFunc("/", handler)

	// Start the web server on port 8080
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}
