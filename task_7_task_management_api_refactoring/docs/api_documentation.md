# Task Management REST API Documentation

## Objective

Refactor the existing Task Management REST API using Clean Architecture principles to improve maintainability, testability, and scalability.

## Base URL

`http://localhost:8080/api/v1`

## Postman Documentation

You can find the Postman documentation for this API at the following URL:  
[Postman Documentation](http://localhost:8080/swagger/doc.json)  
[Swagger Documentation](http://localhost:8080/swagger/index.html)

## Project Structure

```
task-manager/
├── delivery/
│   ├── main.go
│   ├── controllers/
│   │   └── controller.go
│   └── routers/
│       └── router.go
├── domain/
│   └── task.go
|   └── user.go
|   └── login.go
|   └── error_response.go
├── infrastructure/
│   ├── auth_middleWare.go
│   ├── jwt_service.go
│   └── password_service.go
├── repositories/
│   ├── task_repository.go
│   └── user_repository.go
└── usecases/
    ├── task_usecases.go
    └── user_usecases.go
```

### Layers Overview

- **Delivery/**: Handles incoming HTTP requests and responses.

  - **main.go**: Sets up the HTTP server, initializes dependencies, and defines the routing configuration.
  - **controllers/controller.go**: Handles HTTP requests and invokes the appropriate use case methods.
  - **routers/router.go**: Sets up the routes and initializes the Gin router.

- **Domain/**: Defines the core business entities and logic.

  - **domain.go**: Contains the core business entities such as `Task` and `User` structs.

- **Infrastructure/**: Implements external dependencies and services.

  - **auth_middleWare.go**: Middleware to handle authentication and authorization using JWT tokens.
  - **jwt_service.go**: Functions to generate and validate JWT tokens.
  - **password_service.go**: Functions for hashing and comparing passwords to ensure secure storage of user credentials.

- **Repositories/**: Abstracts the data access logic.

  - **task_repository.go**: Interface and implementation for task data access operations.
  - **user_repository.go**: Interface and implementation for user data access operations.

- **Usecases/**: Contains the application-specific business rules.
  - **task_usecases.go**: Implements the use cases related to tasks, such as creating, updating, retrieving, and deleting tasks.
  - **user_usecases.go**: Implements the use cases related to users, such as registering and logging in.

## Environment Setup

### Create a `.env` File

To connect to MongoDB and configure JWT, create a `.env` file in the root directory of the project and add the following lines:

```
MONGO_URI=mongodb://<username>:<password>@<host>:<port>/<database>
JWT_SECRET=your_jwt_secret
```

- Replace `<username>`, `<password>`, `<host>`, `<port>`, and `<database>` with your MongoDB credentials.
- Replace `your_jwt_secret` with a secure secret key for signing JWT tokens.

Ensure the `.env` file is included in your `.gitignore` file to avoid exposing sensitive information.

## Endpoints

### `/users`

### `/register`

- **POST**: Register a new user.
  - **Request**:
    ```json
    {
      "username": "john_doe",
      "email": "john@example.com",
      "password": "securepassword",
      "role": "user" // default: user if not specified
    }
    ```
  - **Response**:
    ```json
    {
      "id": "60c72b2f9b1d8e001c8e4d5a",
      "username": "john_doe",
      "email": "john@example.com",
      "role": "user",
      "created_at": "2025-04-26T12:00:00Z"
    }
    ```

### `/login`

- **POST**: Authenticate a user and generate a JWT token.
  - **Request**:
    ```json
    {
      "email": "john@example.com",
      "password": "securepassword"
    }
    ```
  - **Response**:
    ```json
    {
      "message": "user logged in successfully",
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
    }
    ```

### `/users`

- **GET**: Retrieve a list of all users (Admin only).
  - **Authorization**: Requires a valid JWT token with the `admin` role.
  - **Response**:
    ```json
    [
      {
        "id": "60c72b2f9b1d8e001c8e4d5a",
        "username": "john_doe",
        "email": "john@example.com",
        "role": "user",
        "created_at": "2025-04-26T12:00:00Z"
      }
    ]
    ```

### `/tasks`

- **GET**: Retrieve a list of all tasks.

  - **URL**: `http://localhost:8080/api/v1/tasks`
  - **Response**:
    ```json
    [
      {
        "id": "60c72b2f9b1d8e001c8e4d5a",
        "title": "Task Title",
        "description": "Task Description",
        "due_date": "2023-12-31",
        "status": "Pending"
      }
    ]
    ```

- **POST**: Create a new task.
  - **URL**: `http://localhost:8080/api/v1/tasks`
  - **Request**:
    ```json
    {
      "title": "Task Title",
      "description": "Task Description",
      "due_date": "2023-12-31",
      "status": "Pending"
    }
    ```
  - **Response**:
    ```json
    {
      "id": "60c72b2f9b1d8e001c8e4d5a",
      "title": "Task Title",
      "description": "Task Description",
      "due_date": "2023-12-31",
      "status": "Pending"
    }
    ```

### `/tasks/:id`

- **GET**: Retrieve the details of a specific task by its ID.

  - **URL**: `http://localhost:8080/api/v1/tasks/:id`
  - **Response**:
    ```json
    {
      "id": "60c72b2f9b1d8e001c8e4d5a",
      "title": "Task Title",
      "description": "Task Description",
      "due_date": "2023-12-31",
      "status": "Pending"
    }
    ```

- **PUT**: Update a specific task.

  - **URL**: `http://localhost:8080/api/v1/tasks/:id`
  - **Request**:
    ```json
    {
      "title": "Updated Task Title",
      "description": "Updated Task Description",
      "due_date": "2024-01-15",
      "status": "Completed"
    }
    ```
  - **Response**:
    ```json
    {
      "id": "60c72b2f9b1d8e001c8e4d5a",
      "title": "Updated Task Title",
      "description": "Updated Task Description",
      "due_date": "2024-01-15",
      "status": "Completed"
    }
    ```

- **DELETE**: Delete a specific task by its ID.

  - **URL**: `http://localhost:8080/api/v1/tasks/:id`
  - **Response**:

    ```json
    {
      "message": "Task deleted successfully"
    }
    ```

    ### `/register`

- **POST**: Register a new user.
  - **Request**:
    ```json
    {
      "username": "john_doe",
      "email": "john@example.com",
      "password": "securepassword",
      "role": "user" // default: user if not specified
    }
    ```
  - **Response**:
    ```json
    {
      "id": "60c72b2f9b1d8e001c8e4d5a",
      "username": "john_doe",
      "email": "john@example.com",
      "role": "user",
      "created_at": "2025-04-26T12:00:00Z"
    }
    ```

### `/login`

- **POST**: Authenticate a user and generate a JWT token.
  - **Request**:
    ```json
    {
      "email": "john@example.com",
      "password": "securepassword"
    }
    ```
  - **Response**:
    ```json
    {
      "message": "user logged in successfully",
      "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
    }
    ```

### `/users`

- **GET**: Retrieve a list of all users (Admin only).
  - **Authorization**: Requires a valid JWT token with the `admin` role.
  - **Response**:
    ```json
    [
      {
        "id": "60c72b2f9b1d8e001c8e4d5a",
        "username": "john_doe",
        "email": "john@example.com",
        "role": "user",
        "created_at": "2025-04-26T12:00:00Z"
      }
    ]
    ```

## Evaluation Criteria

1. **Adherence to Clean Architecture**: Clear separation of concerns and dependency inversion.
2. **Code Organization**: Well-structured layers with distinct responsibilities.
3. **Domain Models and Use Cases**: Decoupled and reusable business logic.
4. **Abstraction of Dependencies**: Interfaces for external dependencies to enable easy substitution and testing.
5. **Backward Compatibility**: Maintains existing API functionality.
6. **Documentation**: Comprehensive documentation of the refactored architecture and design decisions.
7. **Testing**: Unit tests for critical components to ensure correctness and maintainability.
