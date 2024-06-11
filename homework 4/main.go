package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/google/uuid"
)

// task represents a single task with an ID and a description.
type task struct {
	Id          uuid.UUID `json:"id"`          // Unique identifier for the task
	Description string    `json:"description"` // Description of the task
}

// addTask adds a new task to the task list stored in the specified file.
func (t task) addTask(file *os.File) {
	// Read existing task list from the file
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}
	// If file is empty, initialize with an empty task list
	if string(content) == "" {
		content = []byte("[]")
	}
	// Unmarshal existing task list into tasks slice
	var tasks []task
	if err := json.Unmarshal(content, &tasks); err != nil {
		log.Fatal("Error unmarshalling JSON:", err)
		return
	}

	// Append the new task to the task list
	tasks = append(tasks, t)

	// Marshal the updated task list to JSON
	newContent, err := json.Marshal(tasks)
	if err != nil {
		log.Fatalf("Failed to marshal tasks: %s", err)
	}

	// Truncate the file to remove old content
	err = file.Truncate(0)
	if err != nil {
		fmt.Println("Error truncating file:", err)
		return
	}

	// Move the file pointer to the beginning of the file
	_, err = file.Seek(0, 0)
	if err != nil {
		fmt.Println("Error seeking file:", err)
		return
	}

	// Write the updated JSON data to the file
	writer := bufio.NewWriter(file)
	_, err = writer.Write(newContent)
	if err != nil {
		fmt.Println("Error writing JSON data:", err)
		return
	}

	// Flush the buffer to ensure all data is written to the file
	err = writer.Flush()
	if err != nil {
		fmt.Println("Error flushing buffer:", err)
		return
	}

}

// listTasks reads and prints the task list from the specified file.
func listTasks(file *os.File) {
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}
	fmt.Println(string(content))
}

func main() {
	// Open file for read/write operations
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("Failed to open or create file: %s", err)
	}
	defer file.Close()

	// Parse command-line arguments
	if len(os.Args) < 2 {
		log.Fatalf("No command provided")
	}
	cmd := os.Args[1]

	// Execute command based on user input
	switch cmd {
	case "add":
		// Check if task description is provided
		if len(os.Args) < 3 {
			log.Fatalf("No task provided")
		}
		description := os.Args[2]
		t := task{
			Id:          uuid.New(), // Generate a new UUID for the task
			Description: description,
		}
		// Add the task to the task list
		t.addTask(file)
	case "list":
		// List all tasks in the task list
		listTasks(file)
	default:
		log.Fatalf("Unknown command: %s", cmd)
	}
}
