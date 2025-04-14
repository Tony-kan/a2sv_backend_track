# Task Manager RESTful API

## Objective

Build a RESTful API to manage tasks, supporting operations to create, retrieve, update, and delete tasks.

## Endpoints

### `/tasks`

- **GET**: Retrieve a list of all tasks.
- **POST**: Create a new task. Accepts a JSON body with the task's title, description, due date, and status.

### `/tasks/:id`

- **GET**: Retrieve the details of a specific task by its ID.
- **PUT**: Update a specific task. Accepts a JSON body with the new details of the task.
- **DELETE**: Delete a specific task by its ID.
