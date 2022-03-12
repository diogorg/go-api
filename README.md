# Simple Golang API
## Propose
- Learn
- Practice

## Main Features
- Clean code with clean architecture
- SOLID
- Isolated repository
- Postgres DB
- Simple CRUD of Products with soft delete
- DDD
- Strategy
- Docker for Postgres
- Models for each context

## How to use
- Create `.env` from `.env.example` (cp .env.example .env) 
- Change `.env` to use your postgres configuration
- run `docker-compose up`
- run `go run main.go`

## ENDPOINTS
- FindAll Users GET `localhost:8000`
- Store User POST `localhost:8000`
- FindBYId User GET `localhost:8000/{id}`
- DeleteById User DELETE `localhost:8000/{id}`
- Update User PATCH `localhost:8000/{id}`

### Thanks
- By Diogo Gutierre