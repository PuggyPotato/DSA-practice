package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
)

func main() {
    file, err := os.Open("file.txt")
    if err != nil {
        log.Fatalf("Failed to open file: %v", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    lineNumber := 1
    var line10 string
    for scanner.Scan() {
        if lineNumber == 10 {
            line10 = scanner.Text()
            break
        }
        lineNumber++
    }

    if err := scanner.Err(); err != nil {
        log.Fatalf("Error reading file: %v", err)
    }

    if line10 != "" {
        fmt.Println(line10)
    } else {
        fmt.Println("File has less than 10 lines.")
    }
}
