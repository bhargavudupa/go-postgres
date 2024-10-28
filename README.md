
# Using Postgres in Go using GORM

This project is a web application designed to manage Users, Employers, and Employees with a focus on CRUD (Create, Read, Update, Delete) operations. Users is a standalone table while One to Many relationship exists between Employer and Employee. It is built using the Go programming language and the Gin framework, it leverages a PostgreSQL database for data storage and retrieval. It offers various API endpoints that allow users to easily interact with the system, making it suitable for integration into larger applications or for use as a standalone service. The project is structured to promote scalability and maintainability, making it a robust solution for managing personnel data.

## Setup
**Docker Command to Create a Postgres Container**
```
docker run --name go-postgres-test -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres
```
**Docker Command to Access Postgre Database**
```
docker run --name go-postgres-test -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres
```
**SQL Command to Create Database**

Execute the following command in Postgres Shell
```
create database test_1;
```
**Install Go Dependencies**

Navigate to Project Repository and run the following command
```
go mod tidy
```
**Run the Migration Command**
```
migrate -path migrations/ -database "postgresql://postgres:mysecretpassword@localhost:5432/test_1?sslmode=disable" -verbose up
```