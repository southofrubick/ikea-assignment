# ikea-assignment
Tech interview prep for [Martin Rimbro Åstrand](https://se.linkedin.com/in/martin-rimbro-%C3%A5strand)

[Trello board](https://trello.com/b/8ku9FoK8/ikea-assignment)

## REQUIREMENTS
Docker
Golang
Bun

### Languages used
- Golang
- Typescript
- `raw sql`

### Techstack
- PostgreSQL
- PGX for go integration
- Echo webframework
- React

### Other packages used
- react bootstrap
- styled components
- axios
- react-router
- react-router-dom

## How to run
### Start Processes
- Run `docker compose up -d` in the terminal, this:
  - Starts the postgres docker container on port 13927.
- Run the `main.go` file with go, this:
  - initializes the database, by:
    - Creating tables
    - Populating product_type and colour from .txt files
  -  Starts the echo server on port 8080.
-  Run `bun run dev` in the terminal, this:
  - Starts the React server on port 3002.

In order to attach to the database in order to write queries,
run `docker exec -it ikea-assignment-db-1 psql -U user -d db` from the terminal.

## Not yet implemented
- Unit tests for
  - DB logic
  - APIs
  - Pages
  - Components
- E2E testing using playwright
- Responsivity in all components
- Setup script in bash
Had I only and endless supply of time..

## Thoughts
This was a fun assignment, bringing me back to when we created fullstack applications in school.
But with all of the knowledge I have now. And React being the only framework in common.

I've only ever worked with PostgreSQL using Django's ORM in Python, and thought it would be a fun challenge to set it up for a GO application, and I've learned a lot.


