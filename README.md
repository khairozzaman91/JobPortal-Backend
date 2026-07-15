# Job Portal Backend (Go)

A scalable Job Portal Backend API built with Go using the standard `net/http` package. This project is being developed incrementally, where each feature is implemented, tested with Postman, and verified before moving to the next development step.

---

# Development Progress

## Step 1: Project Setup

### Development

- Initialized Go Module
- Organized project structure
- Configured HTTP server using Go standard library
- Established clean and modular project architecture

### Testing

- Server startup verified
- Basic API routing tested

---

## Step 2: Job Management

### Development

- Implemented Job domain model
- Created POST `/jobs`
- Implemented GET `/jobs`
- Implemented PUT `/jobs/{id}`
- Implemented DELETE `/jobs/{id}`
- Completed full CRUD operations for job management

### Testing (Postman)

- Create Job with valid JSON
- Retrieve all jobs
- Update job information
- Delete job by ID
- Invalid JSON handling
- Method validation
- JSON response verification

---

## Step 3: User Management

### Development

- Implemented User domain model
- Created POST `/users`
- Implemented GET `/users`
- Implemented POST `/login`
- User authentication using email and password

### Testing (Postman)

- User registration tested
- Retrieve users tested
- Login authentication tested
- Invalid login credentials tested

---

## Step 4: Environment Configuration

### Development

- Added `.env` file support
- Implemented centralized configuration loader
- Loaded application configuration from environment variables

### Configuration

- Service Name
- Version
- HTTP Port
- JWT Secret

### Testing

- Environment variables loaded successfully
- Server started using configured port

---

## Step 5: Middleware Implementation

### CORS Middleware

- Added CORS support
- Configured allowed origins
- Configured allowed HTTP methods
- Configured allowed headers
- Added preflight request handling (OPTIONS)

### Logger Middleware

- Implemented request logging middleware
- Logs HTTP method
- Logs request path
- Tracks request execution time

### Testing (Postman)

- CORS headers verified
- OPTIONS preflight request tested
- API request logging verified

---

## Step 6: Custom JWT Authentication

### Development

Implemented JWT authentication from scratch without using any third-party JWT library.

- Custom JWT generation
- HMAC-SHA256 signature
- Base64URL encoding
- JWT payload creation
- Issued At (`iat`)
- Expiration Time (`exp`)

### JWT Payload

- User ID (`sub`)
- First Name
- Last Name
- Email

### Testing (Postman)

- JWT generated after successful login
- JWT returned to client
- Token structure verified

---

## Step 7: Authorization Middleware

### Development

Implemented custom JWT Authorization Middleware.

- Bearer Token validation
- JWT signature verification
- Payload decoding
- Token expiration validation
- Protected routes support

### Protected Endpoints

- POST `/jobs`
- PUT `/jobs/{id}`
- DELETE `/jobs/{id}`

### Testing (Postman)

- Valid JWT accepted
- Missing token rejected
- Invalid token rejected
- Expired token rejected
- Protected routes verified

---

# Tech Stack

- Go (Golang)
- Standard Library (`net/http`)
- Custom JWT Authentication
- Environment Configuration (`.env`)
- Postman

---

## Middleware

Implemented a custom Middleware Manager to simplify middleware registration and execution. The manager supports both global middleware and route-specific middleware, making the request pipeline clean, modular, and easy to extend.

**Features Added:**
- Middleware Manager
- Global Middleware Registration (`Use`)
- Route-Specific Middleware (`With`)
- Middleware Chaining
- CORS Middleware
- Logger Middleware

**Tested with Postman:**
- Verified Global Middleware execution
- Verified Route-Specific Middleware execution
- Verified protected routes require JWT authentication
- Verified public routes are accessible without authentication
- Verified CORS preflight requests
- Verified request logging for all incoming requests

# Next Development Steps

- Password Hashing (bcrypt)
- PostgreSQL Integration
- Role-Based Authorization (RBAC)
- Job Application System
- Rate Limiting
- Refresh Token
- Unit Testing
- Docker Support

---

# Current Status



Project Status: In Progress 