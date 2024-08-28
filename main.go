package main

import (
    "fmt"
    "net/http"
)

// Handler untuk route "/"
func haloDuniaHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Halo dunia")
}

func main() {
    // Mendaftarkan handler untuk route "/"
    http.HandleFunc("/", haloDuniaHandler)

    // Menjalankan server pada port 8080
    fmt.Println("Server berjalan pada http://localhost:8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Gagal menjalankan server:", err)
    }
}

