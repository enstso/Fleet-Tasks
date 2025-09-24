Got it ✅ Here’s a clean **README.md** for your **Fleet-Task** project in English:

---

# Fleet-Task

Fleet-Task is a lightweight **Go REST API** to manage *tasks* and *users*.
It exposes simple HTTP endpoints to create, list, retrieve, and delete resources.

---

## 🚀 Getting Started

### Prerequisites

* [Go](https://go.dev/dl/) ≥ 1.20 installed
* `git` installed

### Installation & Run

```bash
# Clone the repo
git clone https://github.com/enstso/Fleet-Tasks.git
cd Fleet-Tasks/cmd

# Run the server
go run main.go
```

By default, the server runs on [http://localhost:8080](http://localhost:8080).

---

## 📚 API Endpoints

### 🔹 Tasks

* `GET /tasks` → returns all tasks
* `POST /task` → creates a new task (expects JSON body)
* `GET /task/{id}` → returns a task by its ID
* `DELETE /task/{id}` → deletes a task by its ID

### 🔹 Users

* `GET /users` → returns all users
* `POST /user` → creates a new user (expects JSON body)
* `GET /user/{id}` → returns a user by ID

---

## 📦 Project Structure

```
Fleet-Tasks/
├── cmd/
│   └── main.go            # Server entry point
├── internal/
│   ├── handler/           # HTTP handlers (Tasks, Users, etc.)
│   ├── service/           # Service Business Logical (Task, User)
│   ├── domain/            # Domain Entities (Task, User)
│   └── utils/             # Utilities and constants
```

## 📌 Possible Improvements

* Use a more advanced router (e.g., [chi](https://github.com/go-chi/chi), [gorilla/mux](https://github.com/gorilla/mux)) for cleaner route & method handling.
* Add a database (SQLite, Postgres, etc.) instead of storing data in memory.
* Write unit and integration tests.
* Add `Docker` support for easier deployment.

---

👉 Do you want me to also add **example JSON schemas** for tasks and users in the README (so consumers of the API know what payloads to send)?
