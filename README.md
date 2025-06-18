# Go-CRUD-Application

A simple CRUD (Create, Read, Update, Delete) application built with Go, using the Gin framework and GORM for database interactions.

## Table of Contents

* [Getting Started](#getting-started)
* [Endpoints](#endpoints)
* [Database Setup](#database-setup)
* [Running the Application](#running-the-application)

## Getting Started

This application is built using Go 1.17 and the following dependencies:

* Gin framework for routing and middleware
* GORM for database interactions
* PostgreSQL as the database management system

## Endpoints

The following endpoints are available:

* `GET /`: Welcome message
* `GET /api/tags`: Retrieve all tags
* `GET /api/tags/:tagId`: Retrieve a tag by ID
* `POST /api/tags`: Create a new tag
* `PATCH /api/tags/:tagId`: Update a tag
* `DELETE /api/tags/:tagId`: Delete a tag

## Database Setup

The application uses a PostgreSQL database. To set up the database, create a new database with the following credentials:

* Host: `localhost`
* Port: `5432`
* User: `postgres`
* Password: `postgres`
* Database name: `test`

## Running the Application

To run the application, navigate to the project directory and execute the following command:

```bash
go run main.go
