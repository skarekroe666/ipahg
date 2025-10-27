# IPAHG - GitHub User Info CLI

A command-line tool written in Go that fetches and displays basic information about any GitHub user.

## Overview

IPAHG (presumably "IP API GitHub") is a CLI application built with Go and Cobra that interacts with the GitHub API to retrieve user information including their profile URL, username, number of public repositories, account creation date, and social handles.

## Features

- Fetch GitHub user information via GitHub API
- Rate limiting (5 requests per 5 seconds)
- Clean, formatted output
- Built with Cobra CLI framework

## Prerequisites

Before running this project, ensure you have the following installed:

- **Go 1.25.0 or higher** - [Download Go](https://golang.org/dl/)
- **Git** (for cloning the repository)
- Internet connection (to access GitHub API)

## Project Structure

```
ipahg/
├── cmd/
│   ├── root.go          # Root command definition
│   └── cmd.go           # Fetch command implementation
├── internal/
│   ├── api.go           # API interaction logic
│   ├── api_test.go      # API tests
│   └── models.go        # Data models
├── main.go              # Application entry point
├── go.mod               # Go module dependencies
├── go.sum               # Dependency checksums
├── .env                 # Environment configuration
├── .gitignore           # Git ignore rules
└── README.md            # This file
```

## Step-by-Step Execution Guide

### Step 1: Clone or Navigate to the Project

If you haven't already, navigate to the project directory:

```bash
cd /home/skarekroe/code/go/cli/ipahg
```

### Step 2: Install Dependencies

Download and install all required Go modules:

```bash
go mod download
```

This will install:
- `github.com/spf13/cobra` - CLI framework
- `github.com/joho/godotenv` - Environment variable loader
- `golang.org/x/time/rate` - Rate limiting

### Step 3: Configure Environment Variables

The `.env` file is already configured with the GitHub API URL. Verify it contains:

```
URL="https://api.github.com/users/"
```

**Note:** The `.env` file is listed in `.gitignore` for security purposes. If you clone this repository, you'll need to create this file manually.

### Step 4: Build the Application

Build the binary executable:

```bash
go build -o ipahg .
```

This creates an executable named `ipahg` in the current directory.

### Step 5: Run the Application

#### Option A: Using the Built Binary

To fetch information about a GitHub user:

```bash
./ipahg fetch <github-username>
```

**Example:**

```bash
./ipahg fetch torvalds
```

#### Option B: Using `go run`

Alternatively, run without building:

```bash
go run main.go fetch <github-username>
```

**Example:**

```bash
go run main.go fetch octocat
```

### Step 6: View the Output

The application will display:
- **User URL**: GitHub profile link
- **Username**: GitHub login name
- **Public Repos**: Number of public repositories
- **Created At**: Account creation date
- **Social Handles**: Twitter username (if available)

**Example Output:**

```
User URL: https://github.com/torvalds
Username: torvalds
Public Repos: 8
Created At: 2011-09-03 15:26:22 +0000 UTC
Social Handles: 
```

## Available Commands

### Root Command

```bash
./ipahg
```

Displays help information about the CLI tool.

### Fetch Command

```bash
./ipahg fetch [username]
```

Fetches and displays basic information about the specified GitHub user.

**Arguments:**
- `username` (optional): GitHub username to fetch information for

## Testing

Run the tests:

```bash
go test ./...
```

To run tests with verbose output:

```bash
go test -v ./...
```

To run tests with coverage:

```bash
go test -cover ./...
```

## Development

### Running in Development Mode

```bash
go run main.go fetch <username>
```

## Rate Limiting

The application implements rate limiting to respect GitHub API guidelines:
- **Limit**: 5 requests per 5 seconds

This prevents API abuse and ensures reliable operation.

## Troubleshooting

### Error: "could not load .env"

**Solution:** Ensure the `.env` file exists in the project root with:
```
URL="https://api.github.com/users/"
```

### Error: "API not available"

**Possible causes:**
- Invalid GitHub username
- GitHub API is down
- Network connectivity issues
- Rate limit exceeded

**Solution:** Check the username spelling and try again after a few moments.

### Error: "command not found: ipahg"

**Solution:** Run `go build` first, or use `./ipahg` with the `./` prefix.

## Dependencies

- **cobra v1.10.1** - CLI framework
- **godotenv v1.5.1** - Environment variable management
- **golang.org/x/time v0.14.0** - Rate limiting utilities

## License

See `LICENSE` file for details.

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## API Reference

This tool uses the GitHub REST API:
- **Endpoint**: `https://api.github.com/users/{username}`
- **Method**: GET
- **Authentication**: Not required (public data)
- **Rate Limit**: 60 requests per hour (unauthenticated)

For more information, visit: https://docs.github.com/en/rest/users/users

## Quick Start Summary

```bash
# 1. Navigate to project
cd /home/skarekroe/code/go/cli/ipahg

# 2. Install dependencies
go mod download

# 3. Build the application
go build -o ipahg .

# 4. Run the application
./ipahg fetch octocat
```

---

**Happy GitHub user hunting! 🚀**