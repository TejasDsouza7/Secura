package main

import (
    "log"
    "net/http"
    "secura/auth"
    "secura/db"
    "secura/handlers"
    "secura/storage"
)

func main() {
    db.InitDB()

    config := storage.LoadConfig()

    storage.EnsureStoragePath(config.StoragePath)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "static/index.html") 
    })
    http.HandleFunc("/login", auth.LoginHandler)
    http.HandleFunc("/upload", handlers.UploadHandler)
    http.HandleFunc("/list-files", handlers.ListFilesHandler)
    http.HandleFunc("/download", handlers.DownloadHandler)

    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    log.Printf("Server started on port %s", config.ServerPort)
    log.Fatal(http.ListenAndServe(":"+config.ServerPort, nil))
}