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
3. **.devcontainer/devcontainer.json**: Ensures a consistent development environment within VS Code.

## Development Steps
1. **Set up new GitHub repository**.
2. **Clone the repository locally**.
3. **Create and switch to the `working` branch**.
4. **Add and commit the `.gitignore` file**.
5. **Add and commit the Dockerfile**.
6. **Add and commit the `.devcontainer/devcontainer.json`**.
7. **Create a new Codespace and ensure it builds successfully**.
8. **Proceed with backend development**.

## Additional Considerations
- **User Authentication**: Implement ID.me for verifying veteran status.
- **Testing**: Set up unit and integration tests.
- **Load Testing**: Ensure the application can handle concurrent users effectively.

