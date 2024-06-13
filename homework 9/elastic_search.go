package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

type Student struct {
	Name         string  `json:"name"`
	Age          int64   `json:"age"`
	AverageScore float64 `json:"average_score"`
}

func main() {
	// Create a new client
	cert, err := os.ReadFile("elasticsearch-8.14.1/config/certs/http_ca.crt")
	if err != nil {
		log.Fatalf("Error reading the certificate: %s", err)
	}
	ELASTIC_PASSWORD := os.Getenv("ELASTIC_PASSWORD")

	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
		},
		Username: "elastic",
		Password: ELASTIC_PASSWORD,
		CACert:   cert,
	}
	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	es.Indices.Create("students")

	studentInstance := Student{
		Name:         "Gopher doe",
		Age:          10,
		AverageScore: 99.9,
	}

	// Index a document
	data, err := json.Marshal(studentInstance)
	if err != nil {
		log.Fatalf("Error marshalling the student instance: %s", err)
	}
	index, err := es.Index("students", bytes.NewReader(data))
	if err != nil {
		log.Fatalf("Error indexing the document: %s", err)
	}
	defer index.Body.Close()

	// Get a document
	fmt.Println(es.Get("students", "Q3fNEZABT5QbEXQjqkHW")) // hardocded id --> should parse the id from the index response

	// Search for a document
	query := `{"query" : {
        "match" : { "name" : "doe" }}}`
	found, err := es.Search(
		es.Search.WithIndex("students"),
		es.Search.WithBody(strings.NewReader(query)),
	)
	if err != nil {
		log.Fatalf("Error searching for the document: %s", err)
	}
	fmt.Println(found)
}
