# Example of using Unmock server for testing HTTP traffic in Go

## Instructions

### Adding dependencies

```bash
go get github.com/joho/godotenv
// or
make deps
```

### Creating a build

```bash
go build -o example-golang -v
// or
make build
```

### Running tests

To run tests, you either need to use the Unmock server and use that as network proxy in your tests or create a GitHub access token. The access token should be added as the `GITHUB_TOKEN` environment variable in `.env` file. The token only needs `public_repo` access.

Run tests:

```bash
go test
// or
make test
```
