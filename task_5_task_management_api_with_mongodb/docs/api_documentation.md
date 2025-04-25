# Task Management REST API Documentation

## Objective

Enhance the existing Task Management REST API by integrating MongoDB as the persistent data storage solution using the Mongo Go Driver. This replaces the in-memory database to provide data persistence across API restarts.

## Base URL

`http://localhost:8080/api/v1/tasks`

## Postman Documentation

You can find the Postman documentation for this API at the following URL:  
[Postman Documentation](http://localhost:8080/swagger/doc.json)  
[Swagger Documentation](http://localhost:8080/swagger/index.html)

## Environment Setup

### Create a `.env` File

To connect to MongoDB, create a `.env` file in the root directory of the project and add the following line:

```
MONGO_URI=mongodb://<username>:<password>@<host>:<port>/<database>
```

- Replace `<username>`, `<password>`, `<host>`, `<port>`, and `<database>` with your MongoDB credentials and connection details.
- Example:
  ```
  MONGO_URI=mongodb://admin:password@localhost:27017/taskdb
  ```

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

- **DELETE**: Delete a specific task by its ID from MongoDB.
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
- **data/task_service.go**: Contains business logic and MongoDB operations.
- **router/router.go**: Sets up the routes and initializes the Gin router.
- **docs/api_documentation.md**: Contains API documentation and other related documentation.
- **go.mod**: Defines the module and its dependencies.

## Evaluation Criteria

1. **MongoDB Integration**: Successful integration of MongoDB as the persistent data storage solution.
2. **CRUD Operations**: Correct implementation of CRUD operations using the MongoDB Go Driver.
3. **Error Handling**: Proper error handling for MongoDB operations and network/database errors.
4. **Data Verification**: Verification of data correctness by testing API endpoints and querying MongoDB directly.
5. **Documentation**: Clear and comprehensive documentation for MongoDB integration and API changes.
