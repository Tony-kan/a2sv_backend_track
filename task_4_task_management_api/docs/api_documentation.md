# Task Management REST API Documentation

## Objective

Develop a simple Task Management REST API using Go and the Gin Framework to perform basic CRUD operations for managing tasks.

## Base URL

`http://localhost:8080/api/v1/tasks`

## Endpoints

### `/tasks`

- **GET**: Retrieve a list of all tasks.

  - **Response**:
    ```json
    [
      {
        "id": 1,
        "title": "Task Title",
        "description": "Task Description",
        "due_date": "2023-12-31",
        "status": "Pending"
      }
    ]
    ```

- **POST**: Create a new task.
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
      "id": 1,
      "title": "Task Title",
      "description": "Task Description",
      "due_date": "2023-12-31",
      "status": "Pending"
    }
    ```

### `/tasks/:id`

- **GET**: Retrieve the details of a specific task by its ID.

  - **Response**:
    ```json
    {
      "id": 1,
      "title": "Task Title",
      "description": "Task Description",
      "due_date": "2023-12-31",
      "status": "Pending"
    }
    ```

- **PUT**: Update a specific task.

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
      "id": 1,
      "title": "Updated Task Title",
      "description": "Updated Task Description",
      "due_date": "2024-01-15",
      "status": "Completed"
    }
    ```

- **DELETE**: Delete a specific task by its ID.
  - **Response**:
    ```json
    {
      "message": "Task deleted successfully"
    }
    ```

## Folder Structure

```
task_manager/
├── main.go
├── controllers/
│   └── task_controller.go
├── models/
│   └── task.go
├── data/
│   └── task_service.go
├── router/
│   └── router.go
├── docs/
│   └── api_documentation.md
└── go.mod
```

- **main.go**: Entry point of the application.
- **controllers/task_controller.go**: Handles incoming HTTP requests and invokes the appropriate service methods.
- **models/task.go**: Defines the `Task` struct.
- **data/task_service.go**: Contains business logic and data manipulation functions.
- **router/router.go**: Sets up the routes and initializes the Gin router.
- **docs/api_documentation.md**: Contains API documentation and other related documentation.
- **go.mod**: Defines the module and its dependencies.

## Testing

Use Postman or curl to test the API endpoints. Ensure proper error handling and response codes for scenarios such as:

- Invalid requests.
- Resources not found.
- Successful operations.

## Evaluation Criteria

1. **Implementation**: All required endpoints are implemented according to specifications.
2. **HTTP Methods**: Correct handling of various HTTP methods and response codes.
3. **Error Handling**: Proper validation of input data and error handling.
4. **Code Quality**: Efficient and well-structured code following Go best practices.
5. **Documentation**: Clear and comprehensive documentation of API endpoints.
