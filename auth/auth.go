package auth

import (
    "secura/db"
    "golang.org/x/crypto/bcrypt"
    "net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    u := r.FormValue("username")
    p := r.FormValue("password")

    var hash string
    err := db.DB.QueryRow("SELECT password_hash FROM users WHERE username = ?", u).Scan(&hash)
    if err != nil || bcrypt.CompareHashAndPassword([]byte(hash), []byte(p)) != nil {
        http.Error(w, "Invalid credentials", http.StatusUnauthorized)
        return
    }

    w.Write([]byte("Login successful"))
}