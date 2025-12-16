# SQLC
SQLC tutorial. Go + SQLC leaderboard API. Raw SQL → generated type-safe queries. Clean, modern backend design.
This repository hosts a tutorial demonstrating how to use sqlc to generate type-safe Go code from raw SQL for a simple game's backend. The game features a persistent [leaderboard/score tracking/inventory system] managed entirely through the generated database code, eliminating boilerplate and ensuring type safety.

## Demo

![Image](https://github.com/user-attachments/assets/0a91c7e0-6ee8-4a0f-9b5b-1f35770ef03b)

## Features

- Type-Safe Database Access: Harness the power of sqlc to interact with your database without writing manual boilerplate.
- Simple Game Logic: [Briefly describe the game's core logic, e.g., "Allows players to submit scores and view the top 10 rankings."]
- Clean Separation of Concerns: SQL queries are kept in dedicated .sql files, separate from application logic.
- Easy Setup: Get up and running quickly with [e.g., SQLite, PostgreSQL, or MySQL].


## Run Locally

Clone the project

```bash
  git clone https://github.com/JeffreyOmoakah/SQLC.git
```

Go to the project directory

```bash
  cd SQLC
```

Install dependencies

```bash
  go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

Configure and Generate Code: Ensure your database schema is defined in schema.sql. Your queries are defined in query.sql (or similar)
 Run the generation command: 

```bash
  sqlc generate
```

## Resources
- sqlc GitHub Repository: https://github.com/sqlc-dev/sqlc on GitHub
- sqlc Official Documentation: https://docs.sqlc.dev/
