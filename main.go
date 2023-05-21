package main

import (
	"fmt"
	"log"
	"os"
	"net/http"
)

// import "github.com/milvus-io/milvus-sdk-go/milvus" // Import Milvus Go SDK

// Handler for the home route "/"
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

// Handler for the about route "/about"
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the about page.")
}

func main() {
	// milvus_port := os.Getenv("MILVUS_PORT")
	// // If port is not set, default to 8080
	// if milvus_port == "" {
	// 	milvus_port = "19530"
	// }

	// Get port from environment variables
	port := os.Getenv("GO_MILVUS_PORT")
	// If port is not set, default to 8080
	if port == "" {
		port = "8080"
	}

	// Define your routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)
	
	// Start the server on port 8080
	fmt.Println("Server listening on port " + port)
	log.Fatal(http.ListenAndServe(":" + port, nil))


	// // Connect to Milvus server
	// client, err := client.NewGrpcClient(context.Background(), "localhost" + milvus_port)
	// if err != nil {
	// 	// handle error
	// }

	// if err != nil {
	// 	log.Fatalf("Failed to connect to Milvus: %v", err)
	// }
	// defer client.Close()

	// // Create a collection
	// collectionName := "my_collection"
	// dimension := 128 // Dimensionality of vectors
	// indexType := "IVF_FLAT" // Index type for faster search

	// err = client.CreateCollection(collectionName, dimension, milvus.WithIndexType(indexType))
	// if err != nil {
	// 	log.Fatalf("Failed to create collection: %v", err)
	// }

	// vectors := [][]float32{
	// 	{
	// 		1.0
	// 		, 2.0
	// 		, 3.0, 
	// 		}, 
	// }

	// ids := []int64{1, 2} // Replace with unique IDs for each vector

	// err = client.Insert(collectionName, vectors, ids)
	// if err != nil {
	// 	log.Fatalf("Failed to insert vectors: %v", err)
	// }

	// Perform vector similarity search
	// queryVector := []float32{1.5, 2.5, 3.5, ...} // Replace with your query vector
	// topK := 5 // Number of nearest neighbors to retrieve

	// res, err := client.Search(collectionName, queryVector, topK)
	// if err != nil {
	// 	log.Fatalf("Failed to perform search: %v", err)
	// }

	// // Process search results
	// for _, result := range res.Results {
	// 	fmt.Printf("Vector ID: %d, Distance: %f\n", result.ID, result.Distance)
	// }
}
