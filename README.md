# Enterprise Go Inventory System - Tech Stack Overview

## ✅ Your Final Tech Stack

### 🧠 Architecture

- **Clean Architecture / DDD**: Domain-driven design with layered architecture
- **TDD-first**: Testing-first development (no Postman/Swagger)
- **Modular Package Structure**:

  - `domain/` - domain layer, entities an stuff
  - `usecase/` - business logics/application layet
  - `interfaces/` - endpoints like rest/gql(in the future)
  - `infrastructure/` - things like databases and repositories
  - `config/` - constants like env files
  - `tests/` - e2e tests

---

### 🧱 Backend

| Layer      | Technology                         |
| ---------- | ---------------------------------- |
| Language   | **Go** (Golang)                    |
| Framework  | **Gin** (fast HTTP routing)        |
| Auth       | **Session-based** (cookie + Redis) |
| Validation | `go-playground/validator`          |
| Config     | `github.com/spf13/viper` or `.env` |

---

### 𞳂 Data & Persistence

| Role          | Technology                             |
| ------------- | -------------------------------------- |
| Primary DB    | **PostgreSQL** (relational, ACID)      |
| Session Store | **Redis** (key-value, fast)            |
| ORM/Driver    | `pgx` (recommended), `sqlc`, or `gorm` |

---

### 🧪 Testing

| Purpose               | Tool                             |
| --------------------- | -------------------------------- |
| Unit & use case tests | Go `testing` pkg                 |
| HTTP handler tests    | `httptest` + Gin                 |
| Integration tests     | `testcontainers-go` for DB/Redis |

---

### 🐳 DevOps / Infra

| Component        | Stack                            |
| ---------------- | -------------------------------- |
| Containerization | **Docker** + `docker-compose`    |
| Deployment-ready | Dockerfile with multistage build |
| DB service       | PostgreSQL container             |
| Session service  | Redis container                  |

---

### 📱 Mobile-Friendly Backend

- API returns lightweight JSON responses
- Session-based login with optional mobile-persistent cookies
- Optional CORS setup for hybrid apps / PWA

---

Let me know when you're ready to scaffold the initial codebase or begin defining the first entity/use case.
