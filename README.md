Berikut adalah versi yang lebih rapi dan terstruktur untuk file README.md Anda:

```markdown
# Authentication API with JWT (Echo v3 + dgrijalva/jwt-go)

## Description
Simple authentication API using JWT with Echo framework v3 and `github.com/dgrijalva/jwt-go` library.  
Supports login and JWT-protected profile endpoints.

---

## Prerequisites
- Go 1.18+
- MySQL (optional, if you want database integration)
- Environment variable: `JWT_SECRET` (JWT secret key)

---

## Installation & Run

1. Clone repository:
   ```bash
   git clone <repository-url>
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Set JWT_SECRET environment variable:
   ```bash
   export JWT_SECRET="your_secret_key"
   ```

4. Run the application:
   ```bash
   go run main.go
   ```

## Endpoints
- `POST /register` - Register user
- `POST /login` - Authenticate and get JWT token
- `GET /profile` - Protected endpoint (requires valid JWT)

## License
MIT License © 2025
```
