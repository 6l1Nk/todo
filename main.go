package main

import (
    "bufio"
    "fmt"
    "os"
    "log"
)


func addToDo(todoFile string, todo string) {
    file, err := os.OpenFile(todoFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("Failed to open file: %v", err)
    }
    defer file.Close()

    if _, err := file.WriteString(todo + "\n"); err != nil {
        log.Fatalf("Failed to write to file: %v", err)
    }
    fmt.Println("added")
}

func completeToDo(todoFile string) {
    inputFile, err := os.Open(todoFile)
    if err != nil {
	    log.Fatalf("Failed to open file: %v", err)
    }

    tempFileName := todoFile + ".tmp"
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
    err = os.Remove(todoFile)
    if err != nil {
	    log.Fatalf("snapple dawg $v", err)
    }

    err = os.Rename(tempFileName, todoFile)
    if err != nil {
	    log.Fatalf("renamin ain wut it usd 2b %v", err)
    }

}

func whatsNext(todoFile string) string {
    file, err := os.Open(todoFile)
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

//func deferToDo(todoFile string) {
     // nextTask := whatsNext(todoFile)
    //addToDo("whats next")
    //completeToDo()
 //   return todoFile
//}

func main() {
    if len(os.Args) > 1 {
        command := os.Args[1]
	todoFile := "/home/nestor/dev/todo/todo.txt"
        switch command {
        case "--add":
            if len(os.Args) > 2 {
                addToDo(todoFile, os.Args[2])
            } else {
                fmt.Println("Nothing to add...")
            }
        case "--complete":
            completeToDo(todoFile)
        case "--defer":
            //deferToDo(todoFile)
        case "--whatnext":
            fmt.Println(whatsNext(todoFile))
        default:
            fmt.Println("invalid argument...")
            return
        }
    }
}


