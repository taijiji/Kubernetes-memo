package main

import (
  "io"
  "net/http"
  "os"
)

func main() {
  message := os.Getenv("WDPRESS_MESSAGE")
  if message == "" {
    message = "Hello World!"
  }

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/plain")
    io.WriteString(w, message)
  })

  http.ListenAndServe(":8000", nil)
}
