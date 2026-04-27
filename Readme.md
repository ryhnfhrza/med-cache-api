![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)

# Med-Cache API

A high-performance **RESTful API** built with **Golang**, designed to manage clinical medical records efficiently. This project is a technical showcase of modern backend architecture, focusing on **Database Migrations**, **Distributed Caching with Redis**, and **Container Orchestration with Docker**.

---

## 🚀 Key Features

* **RESTful CRUD:** Clean and efficient API design for managing medical records.
* **Performance Caching:** Integrated **Redis** to reduce database load and achieve sub-millisecond response times.
* **Database Versioning:** Automated schema management using **golang-migrate**.
* **Containerization:** Fully Dockerized environment for seamless "one-command" deployment.
* **Service Health Monitoring:** Built-in health checks ensuring the API only starts when MySQL and Redis are ready.

---

## 🛠 Tech Stack

* **Language:** Golang (Go 1.22+)
* **Database:** MySQL 8.0
* **Cache:** Redis 7.2
* **DevOps:** Docker & Docker Compose
* **Migration Tool:** Golang-migrate

---

## 🏗 Project Structure

This project follows the **Layered Architecture** pattern to ensure maintainability and scalability:

```text
.
├── backend/
│   ├── controller/         # Handles HTTP requests & responses
│   ├── service/            # Contains business logic & cache orchestration
│   ├── repository/         # Handles direct database (MySQL) interactions
│   ├── model/              # Database structs and entities
│   ├── db/migrations/      # SQL migration files (.sql)
│   ├── main.go             # Application entry point
│   └── Dockerfile          # Multi-stage build recipe
├── .env                    # Environment variables (Ignored by Git)
├── .env.example            # Template for environment configuration
└── docker-compose.yml      # Service orchestration
```

## ⚙️ Configuration (Environment Variables)

Before running the application, create a `.env` file in the root directory. You can use the provided `.env.example` as a template. These variables allow the backend to communicate with the database and cache services within the Docker network.

| Variable | Description | Recommended Value |
| :--- | :--- | :--- |
| `APP_PORT` | Port where the API will run | `8080` |
| `DB_HOST` | Database host (use `db` for Docker service name) | `db` |
| `DB_PORT` | Internal MySQL port | `3306` |
| `DB_HOST_PORT` | External port for DBeaver/Workbench access | `3307` |
| `DB_USER` | MySQL Username | `medical_record_user` |
| `DB_PASSWORD` | MySQL Password | `{password}` |
| `DB_NAME` | MySQL Database Name | `db_medical_records` |
| `REDIS_HOST` | Redis host (use `redis` for Docker service name) | `redis` |
| `REDIS_PORT` | Internal Redis port | `6379` |
| `REDIS_PASS` | Redis Password | `{password}` |

---

## 🔄 Application Workflow & Layered Architecture

This project implements a **Layered Architecture** (Clean Architecture principles) to decouple business logic from technical implementations:

1.  **Controller Layer**: Receives HTTP requests, parses JSON input, and calls the appropriate Service.
2.  **Service Layer (The "Brain")**: 
    * Manages the **Cache-Aside** logic.
    * When a request comes in, it checks **Redis** first.
    * If data is found (**Cache Hit**), it returns immediately.
    * If not (**Cache Miss**), it requests data from the Repository and saves a copy to Redis.
3.  **Repository Layer (Data Access)**: 
    * **MySQL Implementation**: Handles all persistent data storage and complex queries.
    * **Redis Implementation**: Handles high-speed caching logic. 
    * The Repository layer is responsible for the **Cache-Aside** strategy: it checks Redis for existing data and falls back to MySQL on a cache miss.
4.  **Model Layer**: Defines the data structures and schema for the medical records.

---

## ⚡ Quick Start

### 1. Prerequisites
* [Docker Desktop](https://www.docker.com/products/docker-desktop/) installed and running.
* [Postman](https://www.postman.com/) for API testing.

### 2. Installation & Deployment
Clone the repository and enter the directory:
```bash
git clone [https://github.com/ezioauditore2/med_cache_api.git](https://github.com/ezioauditore2/med_cache_api.git)
cd med_cache_api 
```

Prepare your environment variables:
```bash
cp .env.example .env
```
Launch all services (API, Database, Redis, and Migrations) with a single command:
```bash
docker-compose up -d
```

---

## 🔗 API Endpoints

| Method | Endpoint | Description | Caching Strategy |
| :--- | :--- | :--- | :--- |
| `GET` | `/api/health` | Check service health status | No Cache |
| `GET` | `/api/records` | Fetch all medical records | **Read Cache** |
| `GET` | `/api/records/:id` | Fetch record by ID | **Read Cache** |
| `POST` | `/api/records` | Add new medical record | Invalidate Cache |
| `PUT` | `/api/records/:id` | Update existing record | Invalidate Cache |
| `DELETE` | `/api/records/:id` | Remove a record | Invalidate Cache |

---

## 🧪 Testing the Implementation

1.  **Database Check**: Access MySQL via DBeaver or MySQL Workbench on `localhost:3307` using the credentials defined in your `.env`.
2.  **Redis Check**: Access the Redis CLI via Docker Desktop Exec tab or terminal:
    ```bash
    docker exec -it med_cache_redis redis-cli -a adminStrongPass
    ```
    Then run `KEYS *` to see the cached medical records.
3.  **Performance Benchmark**: 
    * Use **Postman** to send a `GET` request. 
    * Observe the response time (Latency). 
    * The first request (Cache Miss) will be served by **MySQL**.
    * Subsequent requests (Cache Hit) will be served by the **Redis Repository**, resulting in significantly lower response times.

---

## 🏗 Backend Architecture Details

This project demonstrates proficiency in:
* **Dependency Injection**: Decoupling layers for better testability.
* **Environment Configuration**: Securely managing credentials.
* **Orchestration**: Handling multi-container dependencies and startup sequences (health checks).
* **Data Consistency**: Ensuring the Redis cache is invalidated whenever data is modified in MySQL.

---

## 🛡 License

This project was developed for educational purposes, focusing on backend engineering best practices, system optimization, and containerization.

---
