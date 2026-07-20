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

## (Update) Role Based Authorization (RBAC)

Implemented role-based authorization using JWT.

### What Changed

- Added user role support.
- Added role information inside JWT payload.
- Created `RequireRole` middleware.
- Protected job routes based on user roles.

### Supported Roles

- employer
- jobseeker
- admin


### Authorization Flow


Request
|
v
Authorization Middleware
|
v
JWT Verification
|
v
Extract User Claims
|
v
Role Checking Middleware
|
v
Handler Access

### Protected Routes

Job creation, update, and delete routes now require authentication and role permission.


### Testing

Tested with Postman.

- Employer token → Can access job management routes ✅
- Jobseeker token → Access denied with 403 Forbidden ❌


### 🚀 Recent Updates
- JWT Authentication
- Implemented custom JWT authentication using HMAC SHA256.
- Generated JWT token after successful user login.
- Protected private routes using Authorization middleware.
- Validated JWT signature and token expiration.
- Role-Based Access Control (RBAC)

### Implemented three user roles:

- Admin
- Employer
- Jobseeker
- Protected Routes
- GET /users → Admin only
- POST /jobs → Employer only
- PUT /jobs/{id} → Employer only
- DELETE /jobs/{id} → Employer only
- Postman Testing

### Successfully tested:

- User Registration
- User Login
- JWT Authentication
- Role-Based Authorization
- Route Protection


## Latest Update

### Role-Based Job Creation

- Added role-based access control for job creation.
- Only **Admin** and **Employer** can create job posts.
- **Jobseeker** users receive **403 Forbidden** when attempting to create jobs.
- Automatically assign the logged-in user's ID to the `posted_by` field using JWT claims.
- Protected the `POST /jobs` endpoint with Authorization and RequireRole middleware.


### Job Ownership

- Added ownership validation for updating jobs.
- Employers can update only their own jobs.
- Admin can update any job.
- Unauthorized updates return 403 Forbidden.

# Next Development Steps

- Password Hashing (bcrypt)
- PostgreSQL Integration
- Job Application System
- Rate Limiting
- Refresh Token
- Unit Testing
- Docker Support

---

# Current Status



Project Status: In Progress 