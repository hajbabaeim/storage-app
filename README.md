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

2. Download the project dependencies:

   ```bash
   go mod download
   ```

3. Compile the application:

   ```bash
   go build -o storage-app .
   ```

4. Run the application:

   ```bash
   ./storage-app
   ```

5. Or, simply you can build and run the application with Docker Compose:

   ```bash
   docker-compose up --build
   ```

The application will be accessible at `localhost:1321`.

## API Endpoints

### Get Promotion by ID

- `GET /promotions/:id` - Get promotion by ID.

## Production Operations

### Deployment

Use container orchestration like Kubernetes or managed services like Amazon ECS or Google Cloud Run. Manage environment variables securely via services like Kubernetes Secrets.

### Scaling

Leverage horizontal scaling due to application's stateless nature. For database, consider managed SQL service with read replicas. For Redis, use partitioning or sharding as usage grows.

### Monitoring

Use Prometheus and Grafana for performance insights and alerts. Adopt structured logging with a centralized logging service for log management. Utilize Jaeger or Zipkin for distributed tracing to understand request flows and latencies.

## Performance in Peak Periods

This application employs various technologies and techniques to ensure optimal performance during peak periods, even under a load of millions of requests per minute:

- **Concurrent Processing:** Written in Go, this application leverages Goroutines for concurrent processing, allowing for the handling of multiple requests concurrently, thereby reducing latency.

- **Efficient Data Handling:** The application utilizes batch operations for inserting data into the database, enhancing efficiency in data handling.

- **Caching Mechanism:** The application incorporates Redis cache, reducing database load and expediting data retrieval for frequently accessed data.

- **Database Indexing:** Efficient indexing is implemented in PostgreSQL, significantly speeding up data access times.

- **Rate Limiting:** The application makes use of rate limiter middleware provided by the Fiber framework, limiting the number of requests an individual client can make in a given timeframe, thus preventing individual clients from overloading the server.

- **Horizontal Scaling:** The application is designed to be stateless, greatly simplifying horizontal scaling when the load increases.

Leveraging these technologies and techniques, the application is equipped to efficiently handle high-load scenarios, including peak periods with millions of requests per minute.

## License

MIT
