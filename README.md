# Carbon Offsets Awareness Program

A gamified referral platform promoting carbon offset awareness.

## Table of Contents

1. [Features](#features)
2. [Architecture](#architecture)
3. [Project Structure](#project-structure)
4. [Prerequisites](#prerequisites)
5. [Quick Start](#quick-start)
6. [API Endpoints](#api-endpoints)
7. [Make Commands](#make-commands)
8. [Technologies](#technologies)

## Features

- Points-based gamification system
- User referral tracking
- Automated email notifications
- Real-time leaderboard
- Social sharing integration
- Responsive design
- Search and filtering capabilities
- Pagination support

## Architecture
```
┌─────────────┐         ┌──────────┐
│   Next.js   │ ──────► │   Gin    │
│    Front    │         │  Backend │
└─────────────┘         └──────────┘
                             │
                             ▼
                       ┌──────────┐
                       │ Docker   │
                       └──────────┘
                             │
                             ▼
                       ┌──────────┐
                       │PostgreSQL│
                       └──────────┘
```

## Project Structure

```
vbio-test/
├── client/          # Next.js front
│   ├── app/         # Pages and routing
│   ├── components/  # Shared components
│   ├── services/    # API services
│   └── types/       # TypeScript definitions
├──   server /          #  GoGin backend
│   ├── api/         # HTTP handlers & routes
│   ├── internal/    # Business logic
│   └── cmd/         # Entry points
├── docker/
└── Makefile
```

## Prerequisites

- Go 1.19+
- Node.js 18+
- Docker and Docker Compose
- PostgreSQL
- Make

## Quick Start

1. Clone and setup:
```bash
git clone [repository-url]
cd VBIO-test
cp .env.example .env
```

2. Configure environment variables:
```env
DB_HOST=localhost
DB_USER=your_user
DB_PASSWORD=your_password
DB_NAME=your_db_name
DB_PORT=5432
```

3. Start services:
```bash
make install  # Installs dependencies and Docker containers
make api     # Starts Go server
make cli     # Starts Next.js dev server
```

## API Endpoints

### Register User
```http
POST /api/register
Body: {
    "name": "string",
    "email": "string",
    "phone_number": "string",
    "referral_code": "string" (optional)
}
```

### Get Leaderboard
```http
GET /api/leaderboard
Query: {
    sort: "points" | "name" | "email",
    search: string,
    page: number
}
```

### Get Share Link
```http
GET /api/share/:id
```

## Technologies

### Front
- Next.js 13+
- React 18
- TypeScript
- TailwindCSS
- React Hooks

###   Backend 
-  Go1.19+
- Gin Framework
- GORM
- PostgreSQL

### Tools
- Docker
- Make
- Git

### Dependencies

#### Front Dependencies
```json
{
  "dependencies": {
    "@types/node": "^20.0.0",
    "@types/react": "^18.0.0",
    "@types/react-dom": "^18.0.0",
    "next": "^13.0.0",
    "react": "^18.0.0",
    "react-dom": "^18.0.0",
    "tailwindcss": "^3.0.0",
    "typescript": "^5.0.0"
  }
}
```

#### Back Dependencies
```go
require (
    github.com/gin-gonic/gin v1.10.0
    github.com/google/uuid v1.6.0
    github.com/joho/godotenv v1.5.1
    gorm.io/driver/postgres v1.5.11
    gorm.io/gorm v1.25.12
    github.com/golang-jwt/jwt/v5 v5.2.2
    github.com/lib/pq v1.10.9
)
```

#### Database
- PostgreSQL 15+

#### Development Tools
- Docker Engine 24+
- Docker Compose v2.22+
- Make 4.3+
- Git 2.34+

