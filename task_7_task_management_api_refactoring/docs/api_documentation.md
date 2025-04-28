# Task Management REST API Documentation

## Objective

Enhance the existing Task Management REST API by integrating MongoDB as the persistent data storage solution using the Mongo Go Driver. This replaces the in-memory database to provide data persistence across API restarts.

## Base URL

`http://localhost:8080/api/v1`

## Postman Documentation

You can find the Postman documentation for this API at the following URL:  
[Postman Documentation](http://localhost:8080/swagger/doc.json)  
[Swagger Documentation](http://localhost:8080/swagger/index.html)

## Environment Setup

### Create a `.env` File

To connect to MongoDB, create a `.env` file in the root directory of the project and add the following line:

```
MONGO_URI=mongodb://<username>:<password>@<host>:<port>/<database>
JWT_SECRET=your_jwt_secret
```

- Replace `<username>`, `<password>`, `<host>`, `<port>`, and `<database>` with your MongoDB credentials.
- Replace `your_jwt_secret` with a secure secret key for signing JWT tokens.

Ensure the `.env` file is included in your `.gitignore` file to avoid exposing sensitive information.

## MongoDB Integration

### Enhancements

- **Persistent Data Storage**: MongoDB is now used as the database for storing tasks, replacing the in-memory database.
- **MongoDB Go Driver**: The API uses the official MongoDB Go Driver (`go.mongodb.org/mongo-driver`) for database operations.
- **Backward Compatibility**: The API maintains the same endpoint structure and behavior as the previous version.

### CRUD Operations with MongoDB

The following operations are now implemented using MongoDB:

1. **Create Task**: Insert a new task into the MongoDB collection.
2. **Retrieve All Tasks**: Fetch all tasks from the MongoDB collection.
3. **Retrieve Task by ID**: Fetch a specific task by its ID from MongoDB.
4. **Update Task**: Update an existing task in MongoDB.
5. **Delete Task**: Remove a task from MongoDB.

### Testing MongoDB Integration

- Use Postman or curl to test the API endpoints.
- Verify data persistence by querying the MongoDB database directly or using tools like MongoDB Compass.
- Ensure proper error handling for scenarios such as invalid requests, network errors, and database errors.

## Endpoints

### `/tasks`

- **GET**: Retrieve a list of all tasks from MongoDB.

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

- **POST**: Create a new task in MongoDB.
  - **Request**:
    ```json
    {
      "title": "Task Title",
      "description": "Task Description",
      "due_date": "2023-12-31",
      "status": "Pending" // optional : default pending if not specified
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

- **GET**: Retrieve the details of a specific task by its ID from MongoDB.

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

- **PUT**: Update a specific task in MongoDB.

  - **Request**:
    ```json
    {
      "title": "Updated Task Title", // optional
      "description": "Updated Task Description", // optional
      "due_date": "2024-01-15", // optional
      "status": "Completed" // optional
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

- **DELETE**: Delete a specific task by its ID from MongoDB.
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

## Folder Structure

```
task_manager/
├── main.go
├── controllers/
│   └── controller.go
├── models/
│   ├── task.go
│   └── user.go
├── data/
│   ├── task_service.go
│   └── user_service.go
├── middleware/
│   └── auth_middleware.go
├── router/
│   └── router.go
├── docs/
│   └── api_documentation.md
└── go.mod
```

- **main.go**: Entry point of the application.
- **controllers/controller.go**: Handles HTTP requests for both tasks and user authentication.
- **models/task.go**: Defines the `Task` struct.
- **models/user.go**: Defines the `User` struct and role types.
- **data/task_service.go**: Contains business logic and data manipulation functions for tasks.
- **data/user_service.go**: Contains business logic and data manipulation functions for users, including password hashing.
- **middleware/auth_middleware.go**: Implements middleware to validate JWT tokens for authentication and authorization.
- **router/router.go**: Sets up the routes and initializes the Gin router.
- **docs/api_documentation.md**: Contains API documentation and other related documentation.
- **go.mod**: Defines the module and its dependencies.

## Evaluation Criteria

## Evaluation Criteria

1. **User Management**: Successful implementation of user registration and login.
2. **JWT Authentication**: Proper generation and validation of JWT tokens.
3. **Access Control**: Correct enforcement of role-based access control for protected routes.
4. **Security**: Secure storage and transmission of user credentials using encryption and hashing.
5. **Testing**: Verification of authentication and authorization functionality through API testing.
6. **Documentation**: Clear and comprehensive documentation for authentication and authorization processes.
