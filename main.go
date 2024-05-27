package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var TODO_FILE string = "/home/nestor/Dev/todo/todo.txt"

func addToDo(todo string) {
	file, err := os.OpenFile(TODO_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	if _, err := file.WriteString(todo + "\n"); err != nil {
		log.Fatalf("Failed to write to file: %v", err)
	}
}

func completeToDo() {
	inputFile, err := os.Open(TODO_FILE)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}

	tempFileName := TODO_FILE + ".tmp"
	tempFile, err := os.Create(tempFileName)
	if err != nil {
		log.Fatalf("Failed horrifically %v", err)
	}
	defer tempFile.Close()

	scanner := bufio.NewScanner(inputFile)
	scanner.Scan()
	for scanner.Scan() {
		_, err := tempFile.WriteString(scanner.Text() + "\n")
		if err != nil {
			log.Fatalf("what the heck happnd?, %v", err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("oh my gaaaawddd!! %v", err)
	}

	inputFile.Close()
	err = os.Remove(TODO_FILE)
	if err != nil {
		log.Fatalf("snapple dawg $v", err)
	}

	err = os.Rename(tempFileName, TODO_FILE)
	if err != nil {
		log.Fatalf("renamin ain wut it usd 2b %v", err)
	}

}

func whatsNext() string {
	file, err := os.Open(TODO_FILE)
	if err != nil {
		fmt.Println("Error:", err)
		return "Error"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return scanner.Text()
	} else {
		return ""
	}
}

func deferToDo() {
	addToDo(whatsNext())
	completeToDo()
}

func main() {
	if len(os.Args) > 1 {
		command := os.Args[1]
		switch command {
		case "--add":
			if len(os.Args) > 2 {
				addToDo(os.Args[2])
			} else {
				fmt.Println("Nothing to add...")
			}
		case "--complete":
			completeToDo()
		case "--defer":
			deferToDo()
		case "--whatnext", "-n":
			fmt.Println(whatsNext())
		default:
			fmt.Println("invalid argument...")
			return
		}
	}
}
