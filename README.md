# CookBooked App API

A RESTful API to manage recipes and ingredients, allowing users to create, read, update, and delete (CRUD) recipes and ingredients, and categorize them. The app supports user authentication and allows users to manage their recipes.

## Table of Contents
- [CookBooked App API](#cookbooked-app-api)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
  - [API Documentation](#api-documentation)
  - [Contributing](#contributing)

## Features
- User Registration and Authentication.
- CRUD operations for Recipes.
- CRUD operations for Ingredients.
- Assign Ingredients to Recipes.
- Filter and Search recipes and ingredients.
- [Additional Features]

## Getting Started

### Prerequisites
- Go [version]
- PostgreSQL [version]
- [Other Technologies/Dependencies]

### Installation
1. Clone this repository:
   ```sh
   git clone [your_repo_url]
   ```
2. Install dependencies:
   ```sh
   go mod download
   ```
3. Create a PostgreSQL database:
   ```sh
   createdb recipeapp
   ```
4. Create a `.env` file in the root directory of the project and add the following environment variables:
   ```sh
   PGUSER=[your_db_username]
   PGPASSWORD=[your_db_password]
   PGDATABASE=[your_db_name]
   PGHOST=[localhost|your_db_host]
   PGPORT=[your_db_port|5432]
   DATABASE_URL=[your_db_url]
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
API documentation is available at [CookBooked API]().

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