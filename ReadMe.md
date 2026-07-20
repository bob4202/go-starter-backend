# Go Starter Backend

> *I got tired of creating the same Go project for the 47th time, so I finally made Future Me do the work.*

This is **my personal Go backend starter**. It's the project I clone whenever I want to build a new API instead of spending an hour wiring up the same authentication, database connection, middleware, and folder structure over and over again.

It's not trying to be the next enterprise framework. It's just my "clone, code, ship" template.

## What's already done?

* Gin because it just works.
* PostgreSQL with `sqlx`.
* JWT authentication.
* Password hashing.
* CORS configuration.
* Environment variable loading.
* S3/MinIO support.
* A user module to steal from for future projects.
* A project structure that doesn't make me question my life choices.

Basically all the boring setup is out of the way before I write a single line of actual business logic.

---

## Stack

* Go
* Gin
* PostgreSQL
* SQLX
* JWT
* MinIO / S3
* Docker Compose

---

## Running it

Clone it.

```bash
git clone https://github.com/bob4202/go-starter-backend.git
cd go-starter-backend
```

Copy the environment file.

```bash
cp .env.example .env
```

Start the local services.

```bash
docker compose up -d
```

Run the API.

```bash
go run ./cmd/api
```

Done.

---

## Current features

* User registration
* Login
* JWT authentication
* Current user endpoint
* Password change
* PostgreSQL connection
* MinIO object storage
* Health check endpoint
* Shared response helpers
* Middleware setup

---

## Folder layout

```text
cmd/                # Application entry point
internal/
    config/         # Config loading
    db/             # Database
    middleware/     # JWT, CORS
    server/         # Router
    user/           # User module
pkg/
    response/
    storage/
```

Simple enough that I can find things six months later.

---

## Why this exists

Because every new Go project starts like this:

> "I'll just make a quick API."

Three hours later...

* creating folders
* configuring PostgreSQL
* adding JWT
* writing middleware
* making response helpers
* forgetting CORS
* Googling how to load `.env` files **again**

Not anymore.

Now I clone this repo, rename the module, pretend I'm productive, and start building the actual project.

---

## Future additions

Whenever I need something in another project, it'll probably end up here.

Things that might get added:

* migrations
* refresh tokens
* role-based authorization
* file uploads
* logging
* rate limiting
* email support
* background jobs
* tests (maybe... let's not get carried away)

---

## License

MIT.

Use it if you want.

If it saves you from writing the same boilerplate for the hundredth time, then it already did its job.
