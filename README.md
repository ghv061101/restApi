# RestApi (Users with DOB and Age)

This project is a simple RESTful API in Go that manages users with `name` and `dob` and returns a dynamically computed `age`.

## Run with Docker (Postgres + App)

1. Copy the example env and modify if needed:

```powershell
cp .env.example .env
# edit .env as needed
```

2. Build and run with docker-compose:

```powershell
docker-compose up --build
```

- The API will be available on http://localhost:8080/api
- Postgres runs on port 5432 and data is persisted in a named volume `db_data`.

## Use a managed Postgres (Neon) via DATABASE_URL

This project supports providing a full Postgres connection URL via the `DATABASE_URL` environment variable (recommended for managed providers such as Neon). Example:

```env
DATABASE_URL=postgresql://user:password@host:port/dbname?sslmode=require
```

After setting `DATABASE_URL` (or copying it into `.env`), start the app:

```powershell
go run ./cmd/server
```

The app will run migrations on startup and connect to the database from the URL.

## Notes
- The app expects DB environment variables as shown in `.env.example`.
- `MigrateUsers` runs automatically on startup to create the `users` table.
