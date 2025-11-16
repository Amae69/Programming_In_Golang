# Ticket CLI (Go + Cobra)

A simple command-line tool to track daily tickets and store them in a CSV file.

---

## Features
- Add tickets with title, customer, status, priority, description
- Store data locally in `tickets.csv` (inside user config folder)
- List all tickets or filter by date
- Uses Go + Cobra CLI

---

## Install

```bash
git clone <this-project>
cd ticket-cli
go mod tidy
go build -o ticket-cli .
```

## Usage

### Add a ticket:

```
./ticket-cli add \
  --title "Login failure" \
  --customer "Acme Inc" \
  --priority high \
  --description "User cannot login after reset"
```

### List all tickets:

```
./ticket-cli list
```

### Filter by date

```
./ticket-cli list --date 2025-11-15

```

## CSV Storage

A CSV file is created automatically at:

Our Project directory, and it keeps getting updated, once a new ticket is added using the `.ticket-cli add --flags`

### Columns :

```
ID, Date, Title, Customer, Priority, Status, Description

```