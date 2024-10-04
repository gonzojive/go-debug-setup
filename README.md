# go-debug-setup

[![Go Report Card](https://goreportcard.com/badge/github.com/gonzojive/go-debug-setup)](https://goreportcard.com/report/github.com/gonzojive/go-debug-setup)

A command-line tool to help you set up debugging for Go programs with Delve (dlv).

This tool extracts source file names from a compiled Go binary with debug information. This information can be useful for configuring Delve to correctly locate source files when debugging.

## Installation

`go install github.com/gonzojive/go-debug-setup`

## Example

```shell
go-debug-setup files --executable <path-to-go-binary>
```

This will output a list of source files that were used to compile the binary. You can then use this list to configure your Delve launch configuration to ensure that Delve can find the source files.

## Why is this helpful?

When debugging Go programs with Delve, it sometimes struggles to find the correct source files, especially if the binary was built in a different environment or directory than the one you're debugging in. This tool helps you identify the exact source files that were used to compile the binary, making it easier to configure Delve and avoid "file not found" errors during debugging.

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

This project is licensed under the Apache 2 License - see the [LICENSE](LICENSE) file for details.
