Here's the cleaned up version of your README without using ```bash code blocks:

# Authentication API with JWT (Echo v3 + dgrijalva/jwt-go)

## Description  
Simple authentication API using JWT with Echo framework v3 and github.com/dgrijalva/jwt-go library.  
Supports login and JWT-protected profile endpoints.

---

## Prerequisites  
- Go 1.18+  
- MySQL (optional, for database integration)  
- Environment variable: JWT_SECRET (your JWT secret key)  

---

## Installation & Run  

1. Clone repository:  
   git clone <repository-url>  

2. Install dependencies:  
   go mod tidy  

3. Set JWT secret:  
   export JWT_SECRET="your_secret_key"  

4. Run the application:  
   go run main.go  


---

## Endpoints 
- POST /register - Register User 
- POST /login - Get JWT token  
- GET /profile - Protected endpoint (requires JWT)  

---

## License  
MIT License © 2025