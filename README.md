# Firman's Sport Rental

A RESTful API backend designed to manage sports facility rentals efficiently. This project focuses on high performance, scalability, and secure automated payment integration using the Midtrans Payment Gateway.

## Technologies Used

| Technology                                      | Description                                                           |
| ----------------------------------------------- | --------------------------------------------------------------------- |
| [Go (programming language)](https://go.dev/)    | Main programming language, known for efficiency and concurrency.      |
| [Echo](https://echo.labstack.com/)              | Minimalist web framework for building high-performance REST APIs.     |
| [Postgres](https://www.postgresql.org/)         | Primary Relational Database Management System (RDBMS).                |
| [Supabase](https://supabase.com/)               | Cloud-based PostgreSQL platform for reliable data storage.            |
| [GORM](https://gorm.io/index.html)              | The fantastic ORM library for Golang database management.             |
| [Midtrans](https://midtrans.com/)               | Payment Gateway integration for automated transaction handling.       |
| [Railway](https://railway.com/)                 | Full-stack cloud platform for deployment and monitoring.              |

## Key Features

- User Management: Secure user registration and authentication system.
- Automated Top-Up: Seamless balance top-up system using **Midtrans Snap** integration.
- Real-time Webhook: Automated synchronization of transaction status and user balance via asynchronous *Notification Handlers*.
- Clean Architecture: Implemented using the **Controller-Service-Repository** pattern for highly structured and maintainable code.
- Cloud Deployment: Automatically deployed and hosted on the **Railway** platform with CI/CD integration.

## Project Structure

```
sport-rental/
├── app/                  # Application entry point
│   └── main.go           # Server initialization and routes
├── config/               # Configuration and database setup
│   └── db.go             # GORM & PostgreSQL connection logic
├── controller/           # Request handlers
│   ├── equipment_controller.go
│   ├── rental_controller.go
│   ├── topup_controller.go
│   └── user_controller.go
├── models/               # Database Schema
│   ├── sport.go
│   └── topup.go
├── repository/           # Data access layer
│   ├── equipment_repo.go
│   ├── rental_repo.go
│   ├── topup_repo.go
│   └── user_repo.go
├── service/              # Business logic layer
│   ├── equipment_service.go
│   ├── rental_service.go
│   ├── topup_service.go
│   └── user_service.go
├── .env.example          # Template for environment variables
├── .gitignore            # Files to ignore in Git
├── go.mod                # Go module dependencies
└── go.sum                # Go module checksums
```

## Architecture Flow & Payment Integration
This project follows the Repository-Service-Controller pattern to ensure clean separation of concerns alongside automated payment flow integrated with Midtrans:
- Controller: Handles incoming HTTP requests and sends responses.
- Service: Contains the core business logic (e.g., validating users, processing top-up logic).
- Repository: Directly interacts with the database via GORM SQL expressions.

## Environment Variables
To run this project locally, create a .env file in the root directory and configure the following variables:
```
# Database Configuration
DB_HOST=your_supabase_host
DB_USER=your_database_user
DB_PASSWORD=your_database_password
DB_NAME=your_database_name
DB_PORT=5432

# Midtrans Configuration
MIDTRANS_SERVER_KEY=your_midtrans_server_key
MIDTRANS_IS_PRODUCTION=false
```

## Installation & How to Run
1. Clone the repository
```
git clone https://github.com/Firmeteran/sport-rental.git
cd sport-rental
```

2. Install dependencies
```
go mod tidy
```

3. Run the application
```
go run app/main.go
```

## API Endpoints

| Method    | Endpoint                  | Description                                             |
| --------- | ------------------------- | ------------------------------------------------------- |
| `POST`    | `/register`               | Register a new user account                             |
| `POST`    | `/login`                  | Authenticate user and receive access token              |
| `POST`    | `/topup`                  | Create a top-up request and receive Midtrans Snap URL   |
| `POST`    | `/midtrans/notifications` | Webhook handler for automated payment status updates    |
