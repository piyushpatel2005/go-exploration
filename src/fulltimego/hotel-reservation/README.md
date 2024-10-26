# Hotel Reservation API


## Outline

- users: book a room from a hotel API
- admins: manage rooms, bookings, and users
- hotels: CRUD API
- rooms: CRUD API
- authentication and authorization: using JWT tokens
- scripts: database management, seeding, and migrations

## Set up

```shell
go mod init github.com/piyushpatel2005/hotel-reservation
go get github.com/gofiber/fiber/v2
```
## Requirements

- Go 1.21
- Fiber v2
- Docker
- Docker Compose

```shell
docker run --name mongodb -d mongo:latest -p 27017:27017
```

```shell
go get go.mongodb.org/mongo-driver/mongo
```