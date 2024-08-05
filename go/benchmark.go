package main

import (
    "bytes"
    "fmt"
    "log"
    "os"
    "regexp"
    "time"
)

var (
    emailRegex = regexp.MustCompile(`[\w\.+-]+@[\w\.-]+\.[\w\.-]+`)
    uriRegex   = regexp.MustCompile(`[\w]+://[^/\s?#]+[^\s?#]+(?:\?[^\s#]*)?(?:#[^\s]*)?`)
    ipRegex    = regexp.MustCompile(`(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])\.){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9])`)
)

func measure(data string, r *regexp.Regexp) {
    start := time.Now()

    matches := r.FindAllString(data, -1)
    count := len(matches)

    elapsed := time.Since(start)

    fmt.Printf("%f - %v\n", float64(elapsed)/float64(time.Millisecond), count)
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: benchmark <filename>")
        os.Exit(1)
    }

    filerc, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer filerc.Close()

    buf := new(bytes.Buffer)
    buf.ReadFrom(filerc)
    data := buf.String()

    // Email
    measure(data, emailRegex)

    // URI
    measure(data, uriRegex)

    // IP
    measure(data, ipRegex)
}
