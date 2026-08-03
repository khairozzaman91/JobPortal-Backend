# Job Portal Backend (Go)

A scalable and production-oriented Job Portal Backend built with **Go (Golang)** using the standard `net/http` package. The project follows a clean layered architecture and is being developed incrementally. Every feature is implemented, tested with Postman, verified, and then refactored before moving to the next development phase.

---

# Tech Stack

- Go (Golang)
- Standard Library (`net/http`)
- PostgreSQL
- SQLX
- JWT (Custom Implementation)
- bcrypt
- Postman
- Environment Variables (`.env`)

---

# Project Architecture

```
HTTP Request
      │
      ▼
Middleware
      │
      ▼
Handler
      │
      ▼
Service
      │
      ▼
Repository
      │
      ▼
PostgreSQL
      │
      ▼
HTTP Response
```

---

# Development Journey

---

# Step 1 — Project Initialization

## Development

- Initialized Go Module
- Organized project structure
- Configured HTTP Server
- Built modular folder structure
- Implemented basic routing

### Tested

- Server startup
- API routing
- JSON responses

---

# Step 2 — Job Management

## Development

Implemented complete Job CRUD.

### APIs

- POST `/jobs`
- GET `/jobs`
- PUT `/jobs/{id}`
- DELETE `/jobs/{id}`

### Tested

- Create Job
- List Jobs
- Update Job
- Delete Job
- Invalid JSON
- Invalid HTTP Methods

---

# Step 3 — User Management

## Development

Implemented User Management APIs.

### APIs

- POST `/users`
- GET `/users`
- POST `/login`

### Tested

- User Registration
- User Listing
- Login
- Invalid Credentials

---

# Step 4 — Environment Configuration

## Development

Added centralized application configuration.

### Added

- .env support
- HTTP Port
- JWT Secret
- Service Name
- Version

### Tested

- Environment loading
- Server configuration

---

# Step 5 — Middleware System

## Development

Built a reusable Middleware Manager.

### Implemented

- Global Middleware
- Route Middleware
- Middleware Chaining

### Middlewares

- CORS
- Logger

### Tested

- Request Logging
- CORS
- OPTIONS Request
- Middleware Execution

---

# Step 6 — Custom JWT Authentication

## Development

Implemented JWT completely from scratch without using third-party JWT libraries.

### Features

- Base64URL Encoding
- HMAC SHA256 Signature
- JWT Payload
- Token Expiration
- Issued Time
- Signature Verification

### Payload

- User ID
- First Name
- Last Name
- Email

### Tested

- Login
- Token Generation
- Token Validation

---

# Step 7 — Authorization Middleware

## Development

Protected private APIs using JWT Authorization.

### Features

- Bearer Token Validation
- Signature Verification
- Expiration Validation
- Protected Routes

### Protected APIs

- POST `/jobs`
- PUT `/jobs/{id}`
- DELETE `/jobs/{id}`

### Tested

- Valid Token
- Missing Token
- Invalid Token
- Expired Token

---

# Step 8 — Role-Based Access Control (RBAC)

## Development

Implemented role-based authorization.

### Roles

- Admin
- Employer
- Job Seeker

### Features

- RequireRole Middleware
- Role Validation
- Protected Endpoints

### Permissions

Admin

- Full Access

Employer

- Create Job
- Update Own Job
- Delete Own Job

Job Seeker

- View Jobs

### Tested

- Employer Access
- Admin Access
- Forbidden Access
- 403 Responses

---

# Step 9 — Job Ownership Authorization

## Development

Implemented ownership verification.

### Features

- Automatically assign logged-in user as job owner
- Employer can update only own jobs
- Employer can delete only own jobs
- Admin can manage all jobs

### Tested

- Owner Update
- Owner Delete
- Unauthorized Access

---

# Step 10 — Project Refactoring

## Development

Improved project architecture.

### Refactoring

- Server Struct
- Dependency Injection
- Constructor Functions
- Handler Structs
- Better Routing
- Cleaner Folder Structure

---

# Step 11 — Repository Pattern

## Development

Separated data access from business logic.

### Implemented

- Repository Interface
- Job Repository
- User Repository
- Dependency Injection

### Benefits

- Loose Coupling
- Easy Testing
- Database Independent
- Cleaner Architecture

---

# Step 12 — Service Layer

## Development

Added Service Layer between Handler and Repository.

### Implemented

- Job Service
- User Service
- Business Logic Layer

### Benefits

- Clean Separation
- Easier Maintenance
- Better Validation
- Future Scalability

---

# Step 13 — PostgreSQL Integration

## Development

Migrated repositories from in-memory storage to PostgreSQL using SQLX.

### Completed

- PostgreSQL Connection
- SQLX Integration
- User Repository Migration
- Job Repository Migration

### SQL Operations

- Create
- Read
- Update
- Delete

### Tested

- CRUD Operations
- SQL Queries
- Database Connectivity

---

# Step 14 — Authentication Improvements

## Password Hashing

Implemented secure authentication using bcrypt.

### Features

- Password Hashing
- Password Verification
- Secure Login
- Plain Password Protection

### Tested

- Registration
- Login
- Password Verification

---

# Step 15 — Job Seeker Profile Module

## Development

Created a dedicated Job Seeker Profile module.

### Features

- Create Profile
- View Profile
- Update Profile
- Delete Profile
- JWT Authentication
- PostgreSQL Storage

### Architecture

Handler

↓

Service

↓

Repository

↓

PostgreSQL

---

# Testing

Every feature was verified using **Postman** before moving to the next development phase.

### Verified

- User Registration
- User Login
- JWT Authentication
- Authorization
- RBAC
- Job CRUD
- Job Ownership
- User CRUD
- PostgreSQL CRUD
- Job Seeker Profile
- Password Hashing
- Error Handling

---

# Current Features

### Authentication

- User Registration
- Login
- Custom JWT
- Authorization
- Password Hashing

### Authorization

- RBAC
- Ownership Validation
- Protected Routes

### Jobs

- Create Job
- Update Job
- Delete Job
- View Jobs

### User

- Register
- Login
- CRUD

### Job Seeker

- Create Profile
- Update Profile
- Delete Profile
- View Profile

### Architecture

- Middleware
- Handler
- Service
- Repository
- PostgreSQL
- SQLX
- Dependency Injection

---

# Upcoming Development

The following features are planned before the project reaches its first stable production-ready release.

---

## 1. Pagination

Implement efficient pagination for listing APIs.

### Planned Features

- Page & Limit parameters
- Total records
- Total pages
- Current page metadata
- SQL LIMIT & OFFSET
- Standard paginated response

---

## 2. Rate Limiting

Protect APIs from abuse and excessive traffic.

### Planned Features

- IP-based request limiting
- Configurable request window
- Configurable request limits
- HTTP 429 responses
- Middleware integration
- Automatic counter reset

---

## 3. Payment Integration

Introduce premium payment support.

### Planned Features

- Payment initialization
- Payment verification
- Transaction management
- Premium job posting
- Subscription-ready design
- Secure payment workflow

---

## 4. Database Migration

Implement version-controlled database migrations.

### Planned Features

- SQL Migration Files
- Up Migration
- Down Migration
- Version Tracking
- Automatic Migration Execution
- Easy Environment Setup

---

# Current Status

**Project Status:** 🚧 In Progress

## Completed

- Project Initialization
- Job CRUD
- User Management
- Environment Configuration
- Middleware Manager
- CORS Middleware
- Logger Middleware
- Custom JWT Authentication
- Authorization Middleware
- Role-Based Access Control (RBAC)
- Job Ownership Authorization
- Repository Pattern
- Service Layer
- PostgreSQL Integration (SQLX)
- User Repository Migration
- Job Repository Migration
- Password Hashing (bcrypt)
- Job Seeker Profile Module

## Remaining

- Pagination
- Rate Limiting
- Payment Integration
- Database Migration

---

# Future Goal

The goal of this project is to become a production-ready backend by following modern backend engineering practices including:

- Clean Architecture
- SOLID Principles
- Repository Pattern
- Service Layer
- Secure Authentication
- PostgreSQL
- SQLX
- Scalable API Design
- Production Middleware
- Database Migration
- Payment Integration
- Rate Limiting
- Pagination
