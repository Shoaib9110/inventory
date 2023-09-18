# Grocery Inventory Management System

This is a simple grocery inventory management system with RESTful API endpoints.

## Should have

1.  golang environment setup golang 1.18

## Running the Application

1. Clone the repository.
2. Navigate to the project directory.
3. Run the following command to start the application:

- go mod tidy
- go run main.go

The application will be available at http://localhost:8080.

## API Endpoints

- **POST http://localhost:8080/add**: Add a grocery item to the inventory.
- **GET http://localhost:8080/view**: Retrieve all grocery items in the inventory.
