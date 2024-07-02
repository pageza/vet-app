# Vet-App

## Project Overview
`vet-app` is a backend-only application designed to serve application information to various front-end platforms (web, Android, iOS).

### Tech Stack
- **Programming Language**: Go (Golang)
- **Database**: PostgreSQL
- **Caching**: Redis

### Key Features
- **Veteran Verification**: Uses ID.me for user authentication and verification.
- **Backend Services**: Provides APIs for front-end applications.
- **Front-end Agnostic**: Designed to work with multiple front-end platforms.

## Repository Setup
### Branch Management
- `main`: Stable production-ready branch.
- `working`: Development branch for ongoing work.

### Configuration Files
1. **.gitignore**: Ensures unnecessary files are not committed.
2. **Dockerfile**: Configures the development environment with necessary packages.



## Additional Considerations
- **User Authentication**: Implement ID.me for verifying veteran status.
- **Testing**: Set up unit and integration tests.
- **Load Testing**: Ensure the application can handle concurrent users effectively.

