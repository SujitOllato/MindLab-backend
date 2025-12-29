# MindLab Auth-Service

This is the **Auth Service** for the MindLab platform. It provides user authentication via:

* **Google OAuth 2.0**
* **Email/password signup and login**
* **JWT-based session management**

The service is built using **Go (Gin framework)** and **MySQL**.

---

## Table of Contents

1. [Folder Structure](#folder-structure)
2. [Environment Variables](#environment-variables)
3. [Database Setup](#database-setup)
4. [Install Dependencies](#install-dependencies)
5. [Run Service](#run-service)
6. [API Routes](#api-routes)
7. [Testing](#testing)
8. [Google OAuth Setup](#google-oauth-setup)
9. [JWT Authentication](#jwt-authentication)
10. [Admin APIs](#admin-apis)
11. [License](#license)

---

## Folder Structure

```
auth-service/
│
├── cmd/
│   └── main.go               # Entry point
│
├── internal/
│   ├── config/
│   │   └── config.go         # Environment config
│   │
│   ├── database/
│   │   └── mysql.go          # MySQL connection
│   │
│   ├── models/
│   │   ├── user.go           # User model
│   │   └── auth_provider.go  # Auth provider info
│   │
│   ├── handlers/
│   │   ├── google_auth.go    # Google OAuth handlers
│   │   └── email_auth.go     # Email/password handlers
│   │
│   ├── services/
│   │   ├── google_oauth.go   # Google OAuth validation
│   │   ├── jwt_service.go    # JWT token generation/validation
│   │   └── password_service.go # Password hashing & verification
│   │
│   ├── middleware/
│   │   └── jwt_middleware.go # JWT protected routes
│   │
│   └── routes/
│       └── routes.go         # All API routes
│
├── go.mod
├── go.sum
├── .env                      # Environment variables
└── README.md
```

---

## Environment Variables

Create a `.env` file in the root:

```env
DB_URL="user:password@tcp(localhost:3306)/authdb"
GOOGLE_CLIENT_ID="your-google-client-id"
JWT_SECRET="your-secret-key"
PORT=8081
```

* `DB_URL` – MySQL connection string
* `GOOGLE_CLIENT_ID` – OAuth client ID from Google
* `JWT_SECRET` – Secret key for JWT
* `PORT` – Server port

---

## Database Setup

Create MySQL database and table:

```sql
CREATE DATABASE authdb;

USE authdb;

CREATE TABLE users (
  id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  email VARCHAR(255) UNIQUE NOT NULL,
  name VARCHAR(255),
  provider ENUM('google','local') NOT NULL,
  password_hash VARCHAR(255),
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

---

## Install Dependencies

```bash
go mod tidy
```

Dependencies used:

* `github.com/gin-gonic/gin` → Web framework
* `golang.org/x/oauth2` → OAuth2 support
* `google.golang.org/api/idtoken` → Google ID token validation
* `github.com/dgrijalva/jwt-go` → JWT token handling
* `github.com/go-sql-driver/mysql` → MySQL driver

---

## Run Service

```bash
go run cmd/main.go
```

Service will start on `http://localhost:8081`.

---

## API Routes

### Auth Routes

| Method | Route          | Description             | Body (JSON)                                   |
| ------ | -------------- | ----------------------- | --------------------------------------------- |
| POST   | `/auth/google` | Login with Google OAuth | `{ "id_token": "<token>" }`                   |
| POST   | `/auth/signup` | Email signup            | `{ "email": "", "password": "", "name": "" }` |
| POST   | `/auth/login`  | Email login             | `{ "email": "", "password": "" }`             |
| GET    | `/health`      | Health check            | —                                             |

---

## Testing

### Using Postman

#### Google Login

1. Generate a Google ID token via frontend (see Google OAuth Setup below)
2. POST to `/auth/google` with JSON:

```json
{
  "id_token": "<YOUR_GOOGLE_ID_TOKEN>"
}
```

#### Email Signup

POST `/auth/signup`:

```json
{
  "email": "test@example.com",
  "password": "123456",
  "name": "Test User"
}
```

#### Email Login

POST `/auth/login`:

```json
{
  "email": "test@example.com",
  "password": "123456"
}
```

---

## Google OAuth Setup

1. Go to **Google Cloud Console → APIs & Services → Credentials**.
2. Create **OAuth Client ID** → Web Application.
3. Add **Authorized JavaScript Origins**:

   ```
   http://localhost:3000
   ```
4. Add your Gmail to **Test Users** (OAuth consent screen → Test Users).
5. Use the client ID in `.env` as `GOOGLE_CLIENT_ID`.

---

## JWT Authentication

* On successful login/signup, backend returns a **JWT token**:

```json
{
  "token": "JWT_TOKEN",
  "user": { "id": 1, "email": "user@gmail.com", "name": "User", "provider": "google" }
}
```

* Use `Authorization: Bearer <JWT_TOKEN>` for protected routes.

---

## Admin APIs (Optional)

* Add, edit, delete users or manage OAuth providers
* Protected by **JWT middleware**

---

## License

MIT License


