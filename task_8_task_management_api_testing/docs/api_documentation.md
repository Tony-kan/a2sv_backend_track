# Task Management REST API Documentation

### Objective

Implement comprehensive unit tests for the Task Management API using the `testify` library to ensure the correctness of individual components and increase confidence in the stability of the application.

## Base URL

`http://localhost:8080/api/v1`

## Postman Documentation

You can find the Postman documentation for this API at the following URL:  
[Postman Documentation](http://localhost:8080/swagger/doc.json)  
[Swagger Documentation](http://localhost:8080/swagger/index.html)

## How to Run the Application & Environment Setup

Follow these steps to run the Task Management API:

1. **Navigate to the Root Directory**:
   Open a terminal and navigate to the root directory of the project:

   ```bash
   cd /path/to/task_7_task_management_api_refactoring
   ```

2. **Create a `.env` File**:
   Create a `.env` file in the root directory and add the following environment variables:

   ```
   MONGO_URI=mongodb://<username>:<password>@<host>:<port>/<database>
   JWT_SECRET=your_jwt_secret
   ```

   - Replace `<username>`, `<password>`, `<host>`, `<port>`, and `<database>` with your MongoDB credentials.
   - Replace `your_jwt_secret` with a secure secret key for signing JWT tokens.

3. **Run the Application**:
   Use the following command to start the application:

   ```bash
   go run delivery/main.go
   ```

4. **Access the API**:
   Once the application is running, you can access the API at `http://localhost:8080/api/v1`.

## Project Structure

```
   task_8_task_management_api_testing/
   ├── delivery/
   │   ├── main.go
   │   ├── controllers/
   │   │   └── controller.go
   │   └── routers/
   │       └── router.go
   ├── domain/
   │   ├── task_test.go
   │   └── user_test.go
   |   └── task.go
   |   └── user.go
   |   └── login.go
   |   └── error_response.go
   ├── infrastructure/
   |   └── auth_middleWare_test.go
   │   ├── jwt_service_test.go
   │   └── password_service_test.go
   │   ├── auth_middleWare.go
   │   ├── jwt_service.go
   │   └── password_service.go
   ├── repositories/
   |   └──  mocks/
   │        ├── task_repository.go
   │        └── user_repository.go
   │   ├── task_repository.go
   │   └── user_repository.go
   └── usecases/
       ├── task_usecases_test.go
       └── user_usecases_test.go
       ├── task_usecases.go
       └── user_usecases.go
```

## Unit Testing with Testify

### Test Suite Setup

1. **Install the `testify` Library**:
   Use the following command to install the `testify` library:

   ```bash
   go get github.com/stretchr/testify
   ```

2. **Mocking Dependencies**:
   Use `testify/mock` to mock dependencies and isolate components during testing. For example:

   - Mock the database layer for repository tests.
   - Mock external services like JWT and password hashing.

3. **Setup and Teardown**:
   Implement setup and teardown functions to maintain a clean state between test cases.

4. **Run Tests**:
   Use the following command to run all tests:
   ```bash
   go test ./... -v
   ```

### Test Coverage

- **Domain Models**: Test the behavior of core business entities like `Task` and `User`.
- **Use Cases**: Test business logic for tasks and users, including edge cases.
- **Repositories**: Test data access logic with mocked database interactions.
- **Infrastructure**: Test external services like JWT and password hashing.

### CI Integration

Integrate unit tests into the CI pipeline to automate testing and ensure code quality with each commit. Use a tool like GitHub Actions or Jenkins to run tests on every push or pull request.

### Documentation

- **Running Tests**: Use `go test ./... -v` to run all tests.
- **Test Coverage**: Use `go test ./... -cover` to generate a test coverage report.
- **Mocking**: Use `testify/mock` for mocking dependencies.

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

## Evaluation Criteria

1. **Adherence to Clean Architecture**: Clear separation of concerns and dependency inversion.
2. **Code Organization**: Well-structured layers with distinct responsibilities.
3. **Domain Models and Use Cases**: Decoupled and reusable business logic.
4. **Abstraction of Dependencies**: Interfaces for external dependencies to enable easy substitution and testing.
5. **Backward Compatibility**: Maintains existing API functionality.
6. **Documentation**: Comprehensive documentation of the refactored architecture and design decisions.
7. **Testing**: Unit tests for critical components to ensure correctness and maintainability.
8. **Unit Testing**: Comprehensive unit tests with high coverage using `testify`.
9. **Mocking**: Effective use of mocking to isolate components and ensure test independence.
10. **CI Integration**: Automated testing in the CI pipeline to ensure code quality.
