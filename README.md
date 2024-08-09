```markdown
# Golang CRUD with MongoDB and Gin

This project is a simple CRUD (Create, Read, Update, Delete) application built using Golang, MongoDB, and the Gin web framework. It provides a basic structure for managing products in a MongoDB database.

## Features

- **CRUD Operations**: Create, Read, Update, Delete products.
- **MongoDB Integration**: Uses MongoDB for persistent storage.
- **Gin Framework**: Fast and minimalist web framework for building APIs.
- **Cron Jobs**: Setup for periodic tasks using Go's cron package.
- **Middleware**: Custom error handling middleware.

## Project Structure

```bash
.
├── .env                   # Environment variables
├── connections            # Database connection setup
│   └── db.go
├── controllers            # Controllers for handling HTTP requests
│   └── Products.go
├── cronJobs               # Cron jobs configuration
│   └── cron.go
├── middlerware            # Middleware for the application
│   └── errorHandler.go
├── models                 # Models representing the MongoDB collections
│   └── Products.go
├── services               # Business logic
│   └── productService.go
├── main.go                # Entry point of the application
├── go.mod                 # Go module file
└── go.sum                 # Go dependencies
```

## Prerequisites

Before setting up the project, make sure you have the following installed:

- **Go 1.16+**: Ensure Go is installed on your system. [Download Go](https://golang.org/dl/)
- **MongoDB**: Make sure MongoDB is installed and running. [Install MongoDB](https://docs.mongodb.com/manual/installation/)
- **Environment Variables**: A `.env` file with the following variables:
    ```bash
    MONGO_URI=<your-mongodb-uri>
    MONGO_DB=<your-database-name>
    ```

## Setup

Follow these steps to set up the project:

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   ```

2. **Navigate to the project directory**:
   ```bash
   cd <project-directory>
   ```

3. **Install dependencies**:
   ```bash
   go mod tidy
   ```

4. **Run the application**:
   ```bash
   go run main.go
   ```

## Endpoints

The application exposes the following RESTful API endpoints:

- **POST /products**: Create a new product.
- **GET /products**: Retrieve a list of all products.
- **GET /products/:id**: Retrieve a single product by ID.
- **PUT /products/:id**: Update a product by ID.
- **DELETE /products/:id**: Delete a product by ID.

## Running Cron Jobs

Cron jobs are defined in the `cronJobs/cron.go` file. To start cron jobs along with the application, ensure the `StartCronJobs` function is invoked in the `main.go`. This function schedules periodic tasks as defined in your cron job configuration.

## Middleware

Custom error handling middleware is provided in `middlerware/errorHandler.go`. This middleware intercepts and handles errors globally, ensuring consistent error responses across the application.

## Contributing

Contributions are welcome! If you have any ideas for features or improvements, feel free to fork the repository and submit a pull request. Make sure your changes are well-tested and documented.

