# ikea-assignment
Tech interview prep

[Trello board](https://trello.com/b/8ku9FoK8/ikea-assignment)

## REQUIREMENTS
Docker
Golang
Bun

## How to run
So far..
- Run `docker compose up -d` in the terminal, this:
  - Starts the postgres docker container on port 13927.
- Run the `main.go` file with go, this:
  - initializes the database, by:
    - Creating tables
    - Populating product_type and colour from .txt files
  -  Starts the echo server on port 8080.
-  Run `bun run dev` in the terminal, this:
  - Starts the React server on port 3002.

## Thoughts
N/A
