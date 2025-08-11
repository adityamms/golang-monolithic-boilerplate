
# Golang Monolithic Boilerplate

Golang Monolithic Boilerplate using Domain Driven Design handles authentication and authorization, and beside of this features, It also handles user CRUD using :
* [**Gin Gonic**](https://github.com/gin-gonic/gin)
* [**MongoDB**](https://www.mongodb.com)
* [**Redis**](https://redis.io)
* [**Docker**](https://www.docker.com)
* [**Paseto**](https://paseto.io)

## Terminology

- __User__ &mdash; User is the main entity of this microservice which contain the primary information of user, all the actions of user list below 

- __Authentication__ &mdash; Each user can register and login in the system and take advantage of our platform, therefore we assign a token for handling this operation. (PASETO)

- __Authorization__ &mdash; Each user access defined permissions and take actions on specific entities, therefore we control this permissions through role based authorization.

- __Permission__ &mdash; One the main entities of this microservice and it is key value pair to know which role has which permission.

- __Role__ &mdash; Each user can have one role, and each role can have multiple permissions, we handle user's access to Polaris project 


## Structure of the Code

```
 ┌───┐
 │ / │
 └─┬─┘
   │
   ├───────▶ Common ─ ─ ─ ─ Configuration(s), Env, Middleware, Token
   │
   ├───────▶ DTO ─ ─ ─ ─ Shared Request/Response Structs
   │
   ├───────▶ Components ─ ─ ─ ─ All Components have Handler + Repository
   │                   │ 
   │                   ├──▶ Auth ─ ─ ─ ─ Login User, NewAccessToken
   │                   │      ├── handler.go   // Controller + Service merged
   │                   │      └── repository.go
   │                   │
   │                   ├──▶ User ─ ─ ─ ─ Create, Read, Update User
   │                   │      ├── handler.go
   │                   │      └── repository.go
   │                   │
   │                   ├──▶ Role ─ ─ ─ ─ Manage Roles
   │                   │      ├── handler.go
   │                   │      └── repository.go
   │                   │
   │                   └──▶ ... other features ...
   │
   ├───────▶ Docs ─ ─ ─ ─ ─ Swagger / API Spec
   │
   ├───────▶ Router ─ ─ ─ ─ HTTP Routes
   │
   ├───────▶ Test ─ ─ ─ ─ ─ Test DB & Unit Tests

```

## Setting up the environment

The system needs to have a database (mongodb).

You have to complete your `.env` file based on the `.test.env` template file provided within the root directory of the project.

You can run the docker environment

For start Mongodb :
```
$ make createdb
```

For start Redis :
```
$ make redis
```

For start Swagger :
```
$ make swagger
```

For start server :
```
$ make server
```