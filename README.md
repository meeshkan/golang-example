# Example of using Unmock server for testing HTTP traffic in Go

## Instructions

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

## Contributing

If you notice an error or you'd like to add something new to this example, please [open an issue](https://github.com/unmock/golang-example/issues). We really appreciate the feedback and support! 

Please note that this project is governed by the [Unmock Community Code of Conduct](https://github.com/unmock/code-of-conduct). By participating in this project, you agree to abide by its terms.
