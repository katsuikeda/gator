# gator

A multi-user command line tool for aggregating RSS feeds and viewing the posts.

## Installation

Make sure you have the latest [Go toolchain](https://golang.org/dl/) installed as well as a local PostgreSQL database. You can then install `gator` with:

```bash
go install github.com/katsuikeda/gator@latest
```

## Configuration

Create a configuration file `.gatorconfig.json` in your home directory with the following structure:

```json
{
    "db_url": "postgres://username:@localhost:5432/database?sslmode=disable"
}
```

Replace the values with your database connection string.

## Usage

Run the `gator` command-line tool with the following commands:

- Create a new user:

```bash
gator register <username>
```

- Add and follow a new RSS feed:

```bash
gator addfeed <feed_url>
```

- Start the aggregator:

```bash
gator agg 30s
```

- View aggregated feed posts

```bash
gator browse [limit]
```

There are a few other commands you'll need as well:

- `gator login <name>` - Log in as a user that already exists
- `gator users` - List all users
- `gator feeds` - List all feeds
- `gator follow <url>` - Follow a feed that already exists in the database
- `gator unfollow <url>` - Unfollow a feed that already exists in the database
