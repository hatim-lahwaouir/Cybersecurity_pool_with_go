package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func inspectRequestHandler(w http.ResponseWriter, r *http.Request) {
	// Only allow GET requests
	if r.Method != http.MethodGet {
		http.Error(w, "Only GET requests are allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("=== New Request Received ===")
	fmt.Println()

	// Print Method and URL
	fmt.Printf("Method: %s\n", r.Method)
	fmt.Printf("URL: %s\n", r.URL.String())
	fmt.Printf("Host: %s\n", r.Host)
	fmt.Printf("Remote Address: %s\n", r.RemoteAddr)
	fmt.Printf("Request URI: %s\n", r.RequestURI)
	fmt.Printf("Protocol: %s\n", r.Proto)
	fmt.Println()

	// Print Headers
	fmt.Println("--- Headers ---")
	for name, values := range r.Header {
		for _, value := range values {
			fmt.Printf("%s: %s\n", name, value)
		}
	}
	fmt.Println()

	// Print Query Parameters
	fmt.Println("--- Query Parameters ---")
	if len(r.URL.Query()) > 0 {
		for key, values := range r.URL.Query() {
			fmt.Printf("%s: %s\n", key, strings.Join(values, ", "))
		}
	} else {
		fmt.Println("No query parameters")
	}
	fmt.Println()

	// Print Body (even though GET requests typically don't have bodies)
	fmt.Println("--- Body ---")
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("Error reading body: %v\n", err)
	} else if len(body) > 0 {
		fmt.Printf("%s\n", string(body))
	} else {
		fmt.Println("Empty body")
	}
	fmt.Println()

	// Print Cookies
	fmt.Println("--- Cookies ---")
	if len(r.Cookies()) > 0 {
		for _, cookie := range r.Cookies() {
			fmt.Printf("%s: %s\n", cookie.Name, cookie.Value)
		}
	} else {
		fmt.Println("No cookies")
	}
	fmt.Println()
	fmt.Println("=== End of Request ===")
	fmt.Println()

	// Send response back to client
	response := map[string]interface{}{
		"message": "Request details printed to console",
		"method":  r.Method,
		"url":     r.URL.String(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	// Register the handler
	http.HandleFunc("/", inspectRequestHandler)

	// Start the server
	port := ":8080"
	fmt.Printf("Server starting on http://localhost%s\n", port)
	fmt.Println("Try: curl http://localhost:8080/inspect?param=value -H 'Custom-Header: test'")
	fmt.Println()

	log.Fatal(http.ListenAndServe(port, nil))
}
