# RESTful API with Go and Gin

## Objective

Build an API that provides access to a store selling vintage recordings on vinyl. The API will allow clients to retrieve and add albums.

## Endpoints

### `/albums`

- **GET**: Retrieve a list of all albums, returned as JSON.
- **POST**: Add a new album using request data sent as JSON.

### `/albums/:id`

- **GET**: Retrieve an album by its ID, returning the album data as JSON.

- **main.go**: Entry point of the application.
- **controllers/album_controller.go**: Handles HTTP requests and responses for album-related operations.
- **models/album.go**: Defines the `Album` struct.
- **services/album_service.go**: Contains business logic for managing albums.
- **docs/documentation.md**: Contains system documentation and other related information.
- **go.mod**: Defines the module and its dependencies.

## Steps to Set Up the Project

1. **Create a Project Directory**:

   - Open a terminal and navigate to your home directory.
   - Create a directory for the project:
     ```
     $ mkdir web-service-gin or <chosen directory>
     $ cd web-service-gin or <chosen directory>
     ```

2. **Initialize a Go Module**:

   - Run the following command to create a `go.mod` file:
     ```
     $ go mod init example/web-service-gin or <chosen directory>
     ```

3. **Develop the API**:
   - Implement the endpoints as described above using the Gin framework.

## Evaluation Criteria

1. **Correct Implementation**: Ensure that the endpoints are correctly defined and implemented.
2. **Functionality**: Verify that the API supports retrieving and adding albums as specified.
3. **Code Structure**: Ensure the code follows the provided folder structure and is organized for maintainability.
4. **Documentation**: Provide clear and concise documentation for the API and its usage.
