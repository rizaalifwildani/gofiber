# GO Rest API With Fiber

### Getting Started

##### Migration
- this framework use [PostgreSQL](https://www.postgresql.org) for database and [golang-migrate](https://github.com/golang-migrate/migrate) for migration
- migrate -database "postgres://user:password@host:port/databaseName?sslmode=disable" -path migrations up

##### Run APP
- go mod tidy
- go install [github.com/cosmtrek/air@latest](https://github.com/cosmtrek/air) (for live reload)
- air
- [postman collection](https://www.postman.com/aviation-astronaut-83989734/workspace/gofiber/collection/19025865-122ffb51-bb7a-4a2d-8415-20ad15126597?action=share&creator=19025865)
