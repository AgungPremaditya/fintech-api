
# Fintech-API

This is a project designed to manage financial transactions and records. It provides functionalities to interact with a database to store and retrieve transaction data.


## Features

- Creating Wallets
- Withdraw and Deposit Balances
- Transfer Balances to Other Wallets


## Getting Started

These instructions will help you set up and run the project on your local machine for development and testing purposes.

### Prerequisites

- Go 1.16 or later
- PostgreSQL database
### Installation

**Clone the repository:**

```bash
   git clone git@github.com:AgungPremaditya/fintech-api.git
   cd fintech-api
```

**Set up environment variables:**

Create a .env file in the root directory and specify the following variables:
```
DB_HOST=your_database_host
DB_PORT=your_database_port
DB_USER=your_database_user
DB_PASSWORD=your_database_password
DB_NAME=your_database_name
```

### Run with Builded APP
**Build the project:**

```bash
go build -o fintech-api
```

**Run Builded Application:**
```bash
./fintech-api
```

### Run Locally

**Install Dependencies**
```bash
go mod tidy
```

**Using Air:**

Install Air :
```bash
go install github.com/cosmtrek/air@latest
```

Run with Air for Live Reloading
```bash
air
```
## Documentation

API Documentation is using Postman with link down below:

[Documentation](https://www.postman.com/premadityamandala/workspace/fintech-api/collection/6427079-a0e55079-6e90-459f-85c0-4d61d9f66e13?action=share&creator=6427079&active-environment=6427079-8afcc99b-b081-4d3b-9c96-b2f2083b9b9b)

