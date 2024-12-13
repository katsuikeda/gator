# gator

An RSS feed aggregator in Go.

## Overview

Gator is a command-line tool for aggregating and managing RSS feeds. It allows users to register, follow, and browse RSS feeds, as well as manage user accounts.

## Features

- Register and manage user accounts
- Add and follow RSS feeds
- Fetch and aggregate RSS feed data
- Browse aggregated posts
- List followed feeds and users

## Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/katsuikeda/gator.git
    cd gator
    ```

2. Install dependencies:

    ```sh
    go mod download
    ```

3. Set up the database:

    ```sh
    # Ensure PostgreSQL is running and create a database
    createdb gator

    # Apply the schema
    goose -dir sql/schema postgres "user=youruser dbname=gator sslmode=disable" up
    ```

## Configuration

Create a configuration file at `~/.gatorconfig.json` with the following content:

```json
{
    "db_url": "postgres://youruser:yourpassword@localhost/gator?sslmode=disable",
    "current_user_name": "default_user"
}

## Usage
Run the `gator` command-line tool with the following commands:

- Register a new user:

```sh
gator register <username>
```

- Login as an existing user:

```sh
gator login <username>
```

- Add and follow a new RSS feed:

```sh
gator addfeed <feed_name> <feed_url>
```

- List all feeds:

```sh
gator feeds
```

- Follow a feed:

```sh
gator follow <feed_url>
```

- Unfollow a feed:

```sh
gator unfollow <feed_url>
```

- List followed feeds:
  
```sh
gator following
```

- Aggregate feeds:

```sh
gator agg <time_between_requests>
```

- Browse aggregated feed posts

```sh
gator browse [limit]
```
