# RBAC Project with Golang and Gin

## Introduction

This project is an implementation of Role-Based Access Control (RBAC) using Golang and Gin framework. It provides APIs for managing events, permissions, roles, and users with authentication and authorization features.

## Table of Contents

- [Project Structure](#project-structure)
- [Installation](#installation)
- [Configuration](#configuration)
- [Database Setup](#database-setup)
- [API Endpoints](#api-endpoints)
- [Usage](#usage)
- [Swagger Documentation](#swagger-documentation)

## Project Structure

```plaintext
│   .env
│   another-main-go.txt
│   go.mod
│   go.sum
│   index.html
│   main.go
│   RBAC-project
│   RBAC-project.exe
│   swagger.yaml
│
├───controllers
│       PermissionController.go
│       ResourceController.go
│       RoleController.go
│       UserController.go
│
├───dist
│       favicon-16x16.png
│       favicon-32x32.png
│       index.css
│       oauth2-redirect.html
│       swagger-initializer.js
│       swagger-ui-bundle.js
│       ...
│
├───docs
│       docs.go
│       swagger.json
│       swagger.yaml
│
├───initializers
│       database.go
│       loadEnvVariables.go
│       syncDatabase.go
│
├───middlewares
│       authMiddleware.go
│       authorizationMiddleware.go
│
└───models
        eventModel.go
        permissionModel.go
        roleModel.go
        userModel.go

nstallation
Clone the repository:

bash
Copy code
git clone https://github.com/reztimo/rbac-project.git
cd rbac-project
Install dependencies:

bash
Copy code
go mod download
Configuration
Copy the .env.example file to .env and configure the necessary environment variables.
Database Setup
Run the database initialization:

bash
Copy code
go run initializers/database.go
API Endpoints
Event Endpoints
POST /event: Create a new resource (Admin access required).
GET /event: Get all resources (Authenticated access required).
GET /event/:id: Get a specific resource by ID (Authenticated access required).
PUT /event/:id: Update a resource by ID (Admin access required).
DELETE /event/:id: Delete a resource by ID (Admin access required).
Permission Endpoints
POST /permission: Create a new permission (Admin access required).
GET /permission: Get all permissions (Admin access required).
GET /permission/:id: Get a specific permission by ID (Admin access required).
PUT /permission/:id: Update a permission by ID (Admin access required).
DELETE /permission/:id: Delete a permission by ID (Admin access required).
Role Endpoints
POST /role: Create a new role (Admin access required).
GET /role: Get all roles (Admin access required).
GET /role/:id: Get a specific role by ID (Admin access required).
PUT /role/:id: Update a role by ID (Admin access required).
DELETE /role/:id: Delete a role by ID (Admin access required).
User Endpoints
POST /register: Register a new user.
POST /login: User login.
GET /validate: Validate user credentials (Authenticated access required).
GET /profile: Get all user profiles (Authenticated access required).
GET /profile/:id: Get a specific user profile by ID (Authenticated access required).
GET /logout: Logout (Authenticated access required).
Usage
Run the application:

bash
Copy code
go run main.go
Access the API at http://localhost:5000.

Swagger Documentation
Swagger documentation is available at http://localhost:5000/docs. Explore and test the API endpoints interactively.

javascript
Copy code

Make sure to replace placeholders such as `<URL>` and `<PORT>` with your actual values
```
