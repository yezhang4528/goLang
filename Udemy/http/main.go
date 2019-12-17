package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
)

type logWriter struct {}

func main() {
    resp, err := http.Get("http://google.com")

    if err != nil {
        fmt.Println("Error: ", err)
        os.Exit(-1)
    }

/*
    // Create a byte slice with size 1<<14 (1024*16),
    // and initialize its element to be "empty space"?
    bs := make([]byte, 1<<14)
    resp.Body.Read(bs)
    fmt.Println(string(bs))
*/
    // This line does exactly the same as above commented out code
    lw := logWriter{}
    io.Copy(lw, resp.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
    fmt.Println(string(bs))
    fmt.Printf("logWriter Write function wrote %d bytes\n", len(bs))
    return len(bs), nil
}
