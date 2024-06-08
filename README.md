# Go CRUD App

A simple CRUD application built with Golang, using PostgreSQL for the database, and Docker for containerization. This app includes authentication, user management, and todo management functionalities.

## Features

- User Authentication
- Users CRUD (Create, Read, Update, Delete)
- Todos CRUD (Create, Read, Update, Delete)
- Soft Delete for Users
- Validation of inputs
- Secure user handling

## Prerequisites

- Docker
- Docker Compose
- Golang
- PostgreSQL

## Getting Started

### Set Up Your Development Environment

1. **Clone the repository:**

   ```sh
   git clone https://github.com/Dilip-Valiya/go-crud-app.git
   cd go-crud-app
   ```

2. **Start Docker containers:**

   ```sh
   docker-compose up -d
   ```

3. **Run the application:**

   ```sh
   go run main.go
   ```

## API Endpoints

### Authentication

#### Sign Up

- **URL:** `/signup`
- **Method:** `POST`
- **Request Body:**

  ```json
  {
    "name": "John Doe",
    "email": "john.doe@example.com"
  }
  ```

- **Response:**

  ```json
  {
    "message": "Sign up success!",
    "data": {
      "ID": 1,
      "CreatedAt": "2024-05-29T12:34:56.789Z",
      "UpdatedAt": "2024-05-29T12:34:56.789Z",
      "DeletedAt": null,
      "name": "John Doe",
      "email": "john.doe@example.com",
      "todos": []
    }
  }
  ```

#### Login

- **URL:** `/login`
- **Method:** `POST`
- **Request Body:**

  ```json
  {
    "email": "john.doe@example.com"
  }
  ```

- **Response:**

  ```json
  {
    "message": "Login success!"
  }
  ```

### Users CRUD

#### Create User

- **URL:** `/users`
- **Method:** `POST`
- **Request Body:**

  ```json
  {
    "name": "Jane Doe",
    "email": "jane.doe@example.com"
  }
  ```

- **Response:**

  ```json
  {
    "message": "success",
    "user": {
      "ID": 2,
      "CreatedAt": "2024-05-29T12:34:56.789Z",
      "UpdatedAt": "2024-05-29T12:34:56.789Z",
      "DeletedAt": null,
      "name": "Jane Doe",
      "email": "jane.doe@example.com",
      "todos": []
    }
  }
  ```

#### Get All Users

- **URL:** `/users`
- **Method:** `GET`
- **Response:**

  ```json
  [
    {
      "ID": 1,
      "CreatedAt": "2024-05-29T12:34:56.789Z",
      "UpdatedAt": "2024-05-29T12:34:56.789Z",
      "DeletedAt": null,
      "name": "John Doe",
      "email": "john.doe@example.com",
      "todos": []
    },
    {
      "ID": 2,
      "CreatedAt": "2024-05-29T12:34:56.789Z",
      "UpdatedAt": "2024-05-29T12:34:56.789Z",
      "DeletedAt": null,
      "name": "Jane Doe",
      "email": "jane.doe@example.com",
      "todos": []
    }
  ]
  ```

#### Get User by ID

- **URL:** `/users/{id}`
- **Method:** `GET`
- **Response:**

  ```json
  {
    "ID": 1,
    "CreatedAt": "2024-05-29T12:34:56.789Z",
    "UpdatedAt": "2024-05-29T12:34:56.789Z",
    "DeletedAt": null,
    "name": "John Doe",
    "email": "john.doe@example.com",
    "todos": []
  }
  ```

#### Update User

- **URL:** `/users/{id}`
- **Method:** `PUT`
- **Request Body:**

  ```json
  {
    "name": "John Smith"
  }
  ```

- **Response:**

  ```json
  {
    "ID": 1,
    "CreatedAt": "2024-05-29T12:34:56.789Z",
    "UpdatedAt": "2024-05-29T12:45:00.123Z",
    "DeletedAt": null,
    "name": "John Smith",
    "email": "john.doe@example.com",
    "todos": []
  }
  ```

#### Delete User (Soft Delete)

- **URL:** `/users/{id}`
- **Method:** `DELETE`
- **Response:**

  ```json
  {
    "message": "User deleted successfully"
  }
  ```

### Todos CRUD

#### Create Todo

- **URL:** `/todos`
- **Method:** `POST`
- **Request Header:**

  ```
  Authorization: user@example.com
  ```

- **Request Body:**

  ```json
  {
    "title": "Buy groceries",
    "description": "Milk, Bread, Cheese, Eggs"
  }
  ```

- **Response:**

  ```json
  {
    "ID": 1,
    "CreatedAt": "2024-05-29T12:34:56.789Z",
    "UpdatedAt": "2024-05-29T12:34:56.789Z",
    "DeletedAt": null,
    "title": "Buy groceries",
    "description": "Milk, Bread, Cheese, Eggs",
    "userId": 1
  }
  ```

#### Get All Todos

- **URL:** `/todos`
- **Method:** `GET`
- **Response:**

  ```json
  [
    {
      "ID": 1,
      "CreatedAt": "2024-05-29T12:34:56.789Z",
      "UpdatedAt": "2024-05-29T12:34:56.789Z",
      "DeletedAt": null,
      "title": "Buy groceries",
      "description": "Milk, Bread, Cheese, Eggs",
      "userId": 1
    }
  ]
  ```

#### Get Todo by ID

- **URL:** `/todos/{id}`
- **Method:** `GET`
- **Response:**

  ```json
  {
    "ID": 1,
    "CreatedAt": "2024-05-29T12:34:56.789Z",
    "UpdatedAt": "2024-05-29T12:34:56.789Z",
    "DeletedAt": null,
    "title": "Buy groceries",
    "description": "Milk, Bread, Cheese, Eggs",
    "userId": 1
  }
  ```

#### Update Todo

- **URL:** `/todos/{id}`
- **Method:** `PUT`
- **Request Header:**

  ```
  Authorization: user@example.com
  ```

- **Request Body:**

  ```json
  {
    "title": "Buy groceries and fruits",
    "description": "Milk, Bread, Cheese, Eggs, Apples, Oranges"
  }
  ```

- **Response:**

  ```json
  {
    "ID": 1,
    "CreatedAt": "2024-05-29T12:34:56.789Z",
    "UpdatedAt": "2024-05-29T12:45:00.123Z",
    "DeletedAt": null,
    "title": "Buy groceries and fruits",
    "description": "Milk, Bread, Cheese, Eggs, Apples, Oranges",
    "userId": 1
  }
  ```

#### Delete Todo (Soft Delete)

- **URL:** `/todos/{id}`
- **Method:** `DELETE`
- **Request Header:**

  ```
  Authorization: user@example.com
  ```

- **Response:**

  ```json
  {
    "message": "Todo deleted successfully"
  }
  ```

## Validation

- `title` in todos is mandatory and should not exceed 50 characters.
- `description` in todos is optional and should not exceed 250 characters.
- `name` and `email` in users are mandatory.
- `email` in users must be unique and valid.

## Database Models

### User

```go
type User struct {
    gorm.Model
    Name  string `json:"name" validate:"required"`
    Email string `json:"email" gorm:"uniqueIndex" validate:"required,email"`
    Todos []Todo `json:"todos"`
}
```

```go
type Todo struct {
    gorm.Model
    Title       string `json:"title" gorm:"not null;size:50" validate:"required,max=50"`
    Description string `json:"description" gorm:"size:250" validate:"omitempty,max=250"`
    UserID      uint   `json:"userId" gorm:"not null"`
}
```
