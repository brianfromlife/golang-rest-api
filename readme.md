# Golang API with MongoDB

This repository is for a [Medium Article.](https://brianfromlife.medium.com/build-and-test-an-api-with-golang-and-mongodb-authentication-with-jwt-dependency-mocking-data-c9e549ffac40)

## Get Started

1. Create a .env file at the root of the project and add the following:
   ```
   DB_HOST=localhost:27017
   DB_PASS=password
   DB_USER=root
   DB_NAME=golang_ecs
   ENV=development
   JWT_SECRET=supersecretstring
   JWT_EXPIRES=5
   ```
2. Run `Go mod download`.
3. Run `docker-compose up -d`.
4. Run `make run`.
