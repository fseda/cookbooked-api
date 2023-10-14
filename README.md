# CookBooked App API

The Cookbooked API is a backend service designed specifically to support the Cookbooked application. It provides structured endpoints for managing user data, recipes, and ingredients, ensuring data integrity and secure management for the application’s functionalities.


## Table of Contents
- [CookBooked App API](#cookbooked-app-api)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Technology Stack](#technology-stack)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
  - [API Documentation](#api-documentation)
  - [Contributing](#contributing)

## Features
- **User Registration and Authentication**: Secure user spaces, protecting culinary secrets.
- **Manage Recipes**: Organize, store, and retrieve your cherished recipes.
- **Ingredient Management**: Log and categorize essential cooking components.
- **Searching**: Filter and Search recipes and ingredients.

## Technology Stack

- **[Go](https://go.dev/)**: For scalable and efficient application development.
- **[GORM](https://gorm.io/)**: Ensuring smooth database interactions.
- **[Fiber](https://docs.gofiber.io/)**: A web framework for Go, utilized for efficient API development.
- **[Goose](https://github.com/pressly/goose)**: A database migration tool, managing schema changes and versioning.
- **[PostgreSQL](https://www.postgresql.org/)**: Offering a reliable, open-source database system.
- **[Docker](https://www.docker.com/)**: Guaranteeing consistent API functionality across various computing environments.
- **[Railway](https://railway.app/)**: Deploying the API to the cloud.
- **[Swagger](https://swagger.io/)**: Documenting the API endpoints.

## Getting Started

### Prerequisites
- Go 1.21+
- PostgreSQL 
- Docker (optional)

### Installation
1. Clone this repository:
   ```sh
   git clone https://github.com/fseda/cookbooked-api.git
   ```
2. Install dependencies:
   ```sh
   go mod download
   ```
3. Create a PostgreSQL database:
   ```sh
   createdb [your_db_name] or create a remote psql database
   ```
4. Create a `.env` file in the root directory of the project and add the following environment variables:
   ```sh
   PGUSER=[your_db_username]
   PGPASSWORD=[your_db_password]
   PGDATABASE=[your_db_name]
   PGHOST=[localhost|your_db_host]
   PGPORT=[5432|your_db_port]
   DATABASE_URL=[your_db_url]
   PORT=[3000|your_port]
   JWT_SECRET_KEY=[your_jwt_secret_key]
   GO_ENV=[development|production] (optional)
   ```
5. Run the migrations:
   ```sh
   goose -dir internal/infra/database/migrations postgres "[you_database_url]" up
   ```
6. Run the app:
   ```sh
   go run cmd/http/main.go
   ```

## API Documentation
API documentation is available at [CookBooked API](https://cookbooked-api-deploy-170e.up.railway.app/docs). Or in your [localhost](http://localhost:3000/docs) if you are running the app locally.

## Contributing
Contributions are warmly welcomed! If you'd like to contribute, please follow these steps:

1. Fork the Repository
2. Create your Feature Branch 
    ```sh 
   git switch -c feature/your-feature
   ```
3. Commit your Changes
      ```sh 
      git commit -m 'Add some feature/your-feature'
      ```
4. Push to the Branch
      ```sh 
      git push origin feature/your-feature
      ```
5. Open a Pull Request and Describe the Changes you Made

Your pull request will be reviewed, and we appreciate your patience in waiting for that review.