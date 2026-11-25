# Gator

Gator is a command-line RSS feed aggregator written in Go. It enables multi-user feed management, background scraping, and terminal-based reading of published blog posts using PostgreSQL.

## Features

- **User Management**: Multi-user support with registration, login, and active session tracking stored in a local config.
- **Feed Management**: Add, follow, and unfollow standard RSS feeds per user.
- **Continuous Aggregation**: Long-running background worker that scrapes followed feeds at configurable intervals and saves new posts to the database.
- **Post Browsing**: View recent posts directly in the terminal from all feeds the active user follows.

## Tech Stack

- **Language**: Go
- **Database**: PostgreSQL
- **Database Tooling**: Goose (migrations), sqlc (type-safe Go generation)

## Prerequisites

- [Go](https://golang.org/dl/) (1.22+)
- [PostgreSQL](https://www.postgresql.org/)

## Installation & Setup

1. Clone and install the binary:
   ```bash
   git clone https://github.com/your-username/gator.git
   cd gator
   go install .

    {
      "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
      "current_user_name": ""
    }
    
    goose -dir sql/schema postgres "<YOUR_DB_URL>" up
```
    # User commands
    gator register <username>
    gator login <username>
    gator users
    
    # Feed commands
    gator addfeed <name> <url>
    gator feeds
    gator follow <url>
    gator following
    gator unfollow <url>
    
    # Aggregation & Browsing
    gator agg 1m          # Scrapes feeds every 1 minute
    gator browse [limit]   # Read recent posts (default limit: 2)
    
    
  
