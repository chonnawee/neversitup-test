# Neversitup Test

## Project Structure
    ├── cmd
    │   ├── cli
    │   │   └── main.go
    ├── internal
    │   └── core
    │       ├── implements
    │       │   └── services
    │       │       └── permutations
    │       │           └── permutations_service_implement.go
    │       └── interfaces
    │           └── services
    │               └── permutations_service_interface.go
    └── tests
        └── permutations_service_test.go

## Conventions

1. **cmd/ Directory**:
    - Contains the main entry points for the application.
2. **internal/core/**:
    - The core logic of the application is split into:
        - **implements/services/**: Contains implementations of the business logic for various services such as counter, odd, and permutations. Each service has its own subdirectory.
        - **interfaces/services/**: Contains the interfaces that define the contracts for these services, ensuring clear separation between implementation and definition.
3. **tests/**:
    - Contains unit tests for each service implementation.

