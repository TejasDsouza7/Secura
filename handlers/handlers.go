package handlers

import (
    "encoding/json"
    "golang.org/x/crypto/bcrypt" 
    "io"
    "net/http"
    "os"
    "path/filepath"
    "secura/db"
    "secura/storage"
    "time"
)

type FileInfo struct {
    Username   string `json:"username"`
    Filename   string `json:"filename"`
    UploadedAt string `json:"uploaded_at"`
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", 405)
        return
    }

    username := r.FormValue("username")
    password := r.FormValue("password")
    var hash string
    err := db.DB.QueryRow("SELECT password_hash FROM users WHERE username = ?", username).Scan(&hash)
    if err != nil || bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil {
        http.Error(w, "Unauthorized", 401)
        return
    }

    r.ParseMultipartForm(10 << 20) 
    file, header, err := r.FormFile("file")
    if err != nil {
        http.Error(w, "Invalid file", 400)
        return
    }
    defer file.Close()

    destPath := filepath.Join(storage.LoadConfig().StoragePath, username)
    os.MkdirAll(destPath, 0755)
    out, err := os.Create(filepath.Join(destPath, header.Filename))
    if err != nil {
        http.Error(w, "Error saving file", 500)
        return
    }
    defer out.Close()

    io.Copy(out, file)
    db.DB.Exec("INSERT INTO files (username, filename) VALUES (?, ?)", username, header.Filename)
    w.Write([]byte("Upload successful"))
}

func ListFilesHandler(w http.ResponseWriter, r *http.Request) {
    username := r.URL.Query().Get("username")
    password := r.URL.Query().Get("password")
    var hash string
    err := db.DB.QueryRow("SELECT password_hash FROM users WHERE username = ?", username).Scan(&hash)
    if err != nil || bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil {
        http.Error(w, "Unauthorized", 401)
        return
    }

    rows, err := db.DB.Query("SELECT username, filename, uploaded_at FROM files")
    if err != nil {
        http.Error(w, "Error fetching files", 500)
        return
    }
    defer rows.Close()

    var files []FileInfo
    for rows.Next() {
        var file FileInfo
        var uploadedAt time.Time
        rows.Scan(&file.Username, &file.Filename, &uploadedAt)
        file.UploadedAt = uploadedAt.Format(time.RFC3339)
        files = append(files, file)
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(files)
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
    username := r.URL.Query().Get("username")
    password := r.URL.Query().Get("password")
    var hash string
    err := db.DB.QueryRow("SELECT password_hash FROM users WHERE username = ?", username).Scan(&hash)
    if err != nil || bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil {
        http.Error(w, "Unauthorized", 401)
        return
    }

    name := r.URL.Query().Get("filename")
    path := filepath.Join(storage.LoadConfig().StoragePath, username, name)
    f, err := os.Open(path)
    if err != nil {
    	http.Error(w, "File not found", 404)
        return
    }
    defer f.Close()

    w.Header().Set("Content-Disposition", "attachment; filename="+name)
    w.Header().Set("Content-Type", "application/octet-stream")
    io.Copy(w, f)
}