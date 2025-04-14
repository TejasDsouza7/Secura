```markdown
# Secura Local DB

Secura Local DB is a Go-based local database server that allows users to:
- Log in with a username and password (whitelisted users only).
- Upload files to the server.
- View a list of uploaded files.
- Download files from the server.

---

## Features
- User authentication with hashed passwords.
- File upload and download functionality.
- Cross-platform support.
- Simple and responsive web interface.

---

## Setup Instructions

### 1. Clone the Repository
Clone the repository to your local machine:
```bash
git clone https://github.com/TejasDsouza7/Secura-Local-DB.git
cd Secura-Local-DB
```

---

### 2. Add Your Username and Password
To add your own username and password:
1. Open the `db/db.go` file.
2. Locate the `InitDB` function and modify it to include your username and password:
   ```go
   var exists int
   DB.QueryRow("SELECT COUNT(*) FROM users WHERE username = 'your_username'").Scan(&exists)
   if exists == 0 {
       hash, _ := bcrypt.GenerateFromPassword([]byte("your_password"), bcrypt.DefaultCost)
       DB.Exec("INSERT INTO users (username, password_hash) VALUES (?, ?)", "your_username", string(hash))
   }
   ```
3. Replace `your_username` and `your_password` with your desired credentials.

---

### 3. Update the Configuration File
1. Open the `config.json` file.
2. Update the `storage_path` with the location where you want to store uploaded files. For example:
   ```json
   {
       "storage_path": "Your Storage Location (example: D:/storage)",
       "server_port": "8080"
   }
   ```
3. Replace `Your Storage Location (example: D:/storage)` with your desired file storage path.

---

### 4. Install Dependencies
Run the following command to install the required Go modules:
```bash
go mod tidy
```

---

### 5. Run the Server
Start the server by running:
```bash
go run main.go
```

---

### 6. Access the Application
#### On the Same Machine:
Open your browser and navigate to:
```
http://localhost:8080
```

#### On Other Devices in the Same Network:
1. Find your server's IP address:
   - On **Windows**:
     - Open Command Prompt and run:
       ```bash
       ipconfig
       ```
     - Look for the `IPv4 Address` under your active network connection.
   - On **Linux**:
     - Open a terminal and run:
       ```bash
       hostname -I
       ```
     - The first IP address in the output is your server's IP.
2. Ensure the device you want to access the server from is connected to the same Wi-Fi network as the server.
3. Open your browser and navigate to:
   ```
   http://yourserverip:8080
   ```
   Replace `yourserverip` with the IP address you found (e.g., `192.168.1.100`).

---

## Folder Structure
```
Secura Local DB/
├── main.go
├── config.json
├── go.mod
├── db/
│   └── db.go
├── storage/
│   └── storage.go
├── auth/
│   └── auth.go
├── handlers/
│   └── handlers.go
├── static/
│   └── index.html
```

---

## Features Overview
1. **Login**:
   - Users log in with their credentials.
   - Credentials are validated against the database.

2. **Upload**:
   - Authenticated users can upload files to the server.
   - Files are stored in the `uploads` directory, organized by username.

3. **View Files**:
   - Users can view all files in the database after logging in.

4. **Download**:
   - Users can download files directly from the database.


## Notes
- **The login functionality above the upload files section on the website is only for testing purposes.** It is used to verify that the username and password you entered in db.go are working correctly.
- Ensure that the device accessing the server is connected to the same Wi-Fi network as the server.
- Use the server's IP address to access the application from other devices.

---
