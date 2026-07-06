# Job Portal Backend (Go)

**Current Progress**

A scalable Job Portal Backend API built with Go using Clean Architecture.

## ✅ Done So Far

**Project Setup:**
- Go Module Initialization
- Clean Architecture Structure

**Job Module:**
- Job Domain Model
- DTOs (CreateJobRequest, JobResponse)
- Job Repository
- Job Service
- Job Handlers (GetAll, GetByID)
- Manual 10 Job Posts for Testing

**Testing:**
- Postman API Testing completed for Job endpoints

## 🛠 Tech Stack
- Go (Golang)
- PostgreSQL
- JWT Authentication
- Clean Architecture


**Features Added:**
- POST /jobs endpoint for creating new jobs
- JSON request body support
- Automatic ID generation
- CORS support with preflight check
- Error handling for invalid requests

**Tested with Postman:**
- Create Job with valid JSON
- Error handling for invalid JSON
- Method not allowed check
## ✅ Delete Job Implemented

- DELETE /jobs/{id} endpoint
- Removes job from list
- Returns success message
- Tested with Postman
- Complete Full CRUD Operations

**Middleware:**
- CORS Support
- Logger Middleware

**Tested with Postman:**
- All CRUD operations tested
- CORS Preflight checked
- Error handling verified
- JSON response validation

 **User Module:**
- User Domain Model
- Create User Handler (POST /users)
  **Tested with Postman:**
- Create User tested successfully

  
## Next Steps
- User Authentication
- Job Application System
- Rate Limiting

**Status**: In Progress

---
