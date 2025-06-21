# CRUD API with Go and MySQL

A clean and simple REST API using Go (Golang), MySQL, and an HTML frontend for user management.

## 📦 Features

- Create, list, and update users  
- RESTful API with clean architecture (separated in layers)  
- MySQL database integration  
- CORS middleware  
- HTML form frontend using Fetch API  

## ⚙️ Requirements

- [Download Go](https://golang.org/dl/) (v1.22 or newer)  
- MySQL installed and running  
- Git  
- Python 3 (optional, for running the HTML locally)

## 📁 Project Structure

crud-go-with-api-and-db/
├── config/ # Database connection
├── handler/ # HTTP handlers
├── model/ # Data model
├── repository/ # Data access (SQL)
├── service/ # Business logic
├── router/ # Route registration
├── middleware/ # CORS middleware
├── view/ # HTML frontend
├── main.go # Application entrypoint
├── go.mod # Go module definition
├── .env # Environment variables (not committed)

## 🧪 MySQL Setup

Create the database and table:

```sql
CREATE DATABASE crud_go;

USE crud_go;

CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100)
);
```
🚀 How to Run the Project
1. Clone the repository

git clone https://github.com/gabrielmartins0312/crud-go-with-api-and-db.git
cd crud-go-with-api-and-db

2. Install Go dependencies

go mod tidy

3. Create a .env file

DB_USER=root
DB_PASS=your_password
DB_HOST=127.0.0.1
DB_PORT=3306
DB_NAME=crud_go

⚠️ Make sure to not commit the .env file. It is listed in .gitignore.

4. Run the API

go run main.go

The API will be available at:

http://localhost:8080

🌐 Frontend (HTML Form)
1. Open the frontend

Navigate into the view/ folder and run a simple server:

cd view
python3 -m http.server 5500

2. Open in browser

http://localhost:5500/form.html

✅ How the Form Works

    Fill name and email, then click Submit to create a user

    Fill ID, name and email, then click Submit to update a user

    Uses POST if ID is empty and PUT if ID is filled

📫 Author

Gabriel Martins – github.com/gabrielmartins0312
