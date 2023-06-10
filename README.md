# Filtering Service - Rectangles
This is a RESTful API built in Go that calculates and saves intersections between rectangles. The API allows you to send a JSON payload with a main rectangle and a list of input rectangles. If an input rectangle intersects with the main rectangle, it is saved in a SQLite database. You can also retrieve the saved intersections through a GET request.

Features:
- POST endpoint ("http://localhost:8080") to calculate intersections between a main rectangle and a list of input rectangles.
- Intersecting rectangles are saved in an SQLite database.
- GET endpoint ("http://localhost:8080") to retrieve a list of previously saved intersections.
- The API uses the Gorilla Mux router for handling HTTP requests.

Usage:
1. Send a POST request to the "http://localhost:8080" endpoint with a JSON payload specifying the main rectangle and input rectangles.
2. The API calculates the intersections and saves them in the database.
3. Send a GET request to the "http://localhost:8080" endpoint to retrieve a list of previously saved intersections.

Example JSON payload for POST request:
```json
{
  "main": {
    "x": 1,
    "y": 1,
    "width": 10,
    "height": 10
  },
  "input": [
    {
      "x": 2,
      "y": 2,
      "width": 5,
      "height": 5
    },
    {
      "x": 8,
      "y": 8,
      "width": 3,
      "height": 3
    }
  ]
}
```

The intersections are calculated based on the overlapping areas of the rectangles. The resulting intersections include the coordinates, width, height, and timestamp of the intersection.

This project can be used as a foundation for applications that require rectangle intersection calculations and persistent storage of the intersections.

Note: Make sure to have Go and the necessary dependencies installed before running the API.

## Dependencies
Make sure you have Go and the necessary dependencies installed before running the API. To prepare the dependencies, run the following commands:
```shell
cd src
go mod init rect.ir/r
go mod tidy
```
Please note that the commands assume that the main file containing the code is named main.go and the project is initialized with the module path rect.ir/r. You can adjust these commands according to your project's structure and needs.

## Running the API
To run the web server, execute the following command:
```shell
go run main.go
```
This will start the API server on localhost at port 8080. You can then make POST requests to calculate and save intersections, and GET requests to retrieve the saved intersections.

Note: Make sure the SQLite database file rectangles.db is present in the same directory as the main.go file.

Enjoy intersecting rectangles with ease using this RESTful API!