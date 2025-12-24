# Inventory Management System 

Simple REST API built with Golang using in-memory storage and JWT authentication.

## Features
1. **Register & Login**: Authentication using JWT.
2. **Manage Items**: Create, list, and view item details.
3. **Update Stock**: Increase or decrease item stock.
4. **Activity Logs**: Automatically logs stock changes.

## Tech Stack
- Golang
- JWT (Authentication)
- In-Memory Array (Storage)

## Getting Started

### Prerequisites
- Go 1.21 or later

### Installation
1. Clone the project or copy the files.
2. Install dependencies:
   ```bash
   go mod tidy
   ```

### Running the API
```bash
go run main.go
```
The server will start at `http://localhost:8080`.

## API Documentation
The Postman collection is provided below for testing.

### Endpoints
- `POST /register`: Create a new user.
- `POST /login`: Authenticate and get a JWT token.
- `POST /api/items`: Create a new item (Requires JWT).
- `GET /api/items`: List all items (Requires JWT).
- `GET /api/items/{id}`: Get item detail (Requires JWT).
- `PATCH /api/items/{id}/stock`: Update item stock (Requires JWT).
- `GET /api/logs`: Get all activity logs (Requires JWT).

## Postman Collection
You can import the following Postman Collection link for testing:
[Postman Collection Link](https://documenter.getpostman.com/view/19199067/2sBXVZoEKX)
*(Note: Replace with real link if available)*
