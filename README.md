# Storage App

Storage App is a simple and efficient RESTful API server that manages promotion data.

## Features

- [x] RESTful API (using Fiber framework)
- [x] CRUD operations for promotion data
- [x] CSV data import
- [x] Data persistence with PostgreSQL
- [x] Dockerized application
- [x] Batch data processing
- [x] Redis caching for efficient data retrieval

## Prerequisites

- Go (v1.17 or newer)
- Docker & Docker Compose
- Redis
- PostgreSQL

## Installation & Running

1. Clone the repository:

   ```bash
   git clone https://github.com/hajbabaeim/storage-app.git
   cd storage-app
   ```

2. Build and run the application with Docker Compose:

   ```bash
   docker-compose up --build
   ```

The application will be accessible at `localhost:1321`.

## API Endpoints

### Get Promotion by ID

- `GET /promotions/:id` - Get promotion by ID.

## License

MIT
