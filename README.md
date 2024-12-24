# Project Shopper

Shopper is a RESTful API for an e-commerce application. This API handles basic CRUD operations for products and orders, and provides user management and authentication.

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.22.2 or later)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/IgweDaniel/shopper.git
   cd shopper
   ```

2. Copy the example environment file and update the environment variables:

   ```bash
   cp .env.example .env
   ```

3. Build the application:

   ```bash
   make build
   ```

4. Run the application:

   ```bash
   make run
   ```

### Running the Tests

Run the test suite:

    ```bash
    make test
    ```

### Database Migrations

Apply all up database migrations:

    ```bash
    make migrations/up
    ```

Rollback all database migrations:

    ```bash
    make migrations/reset
    ```

Move to a specified database migration version:

    ```bash
    make migrations/goto version=<version>
    ```

Force a migration to the given version:

    ```bash
    make migrations/force version=<version>
    ```

Display current database migration version:

    ```bash
    make migrations/version
    ```

### Seeding the Database

Seed the database with initial data:

    ```bash
    make seed
    ```

### Live Reload

Live reload the application:

    ```bash
    make watch
    ```

### Swagger Documentation

Generate Swagger documentation:

    ```bash
    make swagger
    ```

Format Swagger comments:

    ```bash
    make swagger/fmt
    ```

### Docker

Create and run the DB container:

    ```bash
    make docker-run
    ```

Shutdown the DB container:

    ```bash
    make docker-down
    ```

### API Documentation

The API documentation is generated using Swagger. You can access the Swagger UI at `/swagger/index.html` after running the application.

## Contributing

Please read `CONTRIBUTING.md` for details on our code of conduct, and the process for submitting pull requests.

## License

This project is licensed under the Apache 2.0 License - see the `LICENSE` file for details.

## Acknowledgments

- Swaggo for generating Swagger documentation
- Echo for the web framework
- Golang-JWT for JWT authentication
