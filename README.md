# Neversitup Test

## Prerequisites
Ensure you have the following installed:

- [Docker](https://www.docker.com/)

## Installation

### 1. Clone the Repository

```terminal
git clone <repository_url>
cd <project_directory>
```
### 2. Configure Environment
```terminal
cp .env.example .env
```
Edit the .env file to add your configurations.

### 3. Build and Start the Services
To build and start the services using Docker Compose, run:
```terminal
docker-compose up -d
```

## Project Structure
    ├── cmd
    │   ├── cli
    │   │   └── main.go
    │   └── server
    │       └── main.go
    ├── internal
    │   ├── core
    │   │   ├── implements
    │   │   │   └── services
    │   │   │       ├── counter
    │   │   │       │   └── counter_service_implement.go
    │   │   │       ├── odd
    │   │   │       │   └── odd_service_implement.go
    │   │   │       └── permutations
    │   │   │           └── permutations_service_implement.go
    │   │   └── interfaces
    │   │       └── services
    │   │           ├── counter_service_interface.go
    │   │           ├── odd_service_interface.go
    │   │           └── permutations_service_interface.go
    │   └── handlers
    │       ├── counter_hanler.go
    │       ├── odd_handler.go
    │       └── permutations_handler.go
    └── tests
        ├── counter_service_test.go
        ├── odd_service_test.go
        └── permutations_service_test.go

## Conventions

1. **cmd/**:
    - Contains the main entry points for the application.
2. **internal/core/**:
    - The core logic of the application is split into:
        - **implements/services/**: Contains implementations of the business logic. Each service has its own subdirectory.
        - **interfaces/services/**: Contains the interfaces that define the contracts for these services, ensuring clear separation between implementation and definition.
3. **tests/**:
    - Contains unit tests for each service implementation.
4. **handlers/**:
    - Contains handlers for each service.

## Naming Conventions
 - Services: Each service has a specific folder and uses descriptive names, such as counter_service_implement.go for the implementation and counter_service_interface.go for the interface.
 - Test Files: Test files follow the same naming pattern as their corresponding service implementations but with _test.go appended.    
 - Handlers: Each handler has a specific folder and uses descriptive names, such as user_handler.go for the HTTP handler. If the handler has a specific purpose, it may be reflected in the name, such as auth_handler.go for authentication-related routes.