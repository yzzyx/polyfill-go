
# polyfill-go/builder

Helper to generate the code for the following tools:

- `useragent/parse.go` User-agent parser based on [useragent_parser](https://github.com/Financial-Times/useragent_parser)
- `useragent/normalise.go` - User-agent normaliser based on [polyfill-useragent-normaliser](https://github.com/Financial-Times/polyfill-useragent-normaliser)

## Usage

```
Usage of ./builder:
  -normalise
        build user-agent normalising code
  -useragent
        build user-agent parsing code
```

Builder writes it's output to standard output.
To build the files to polyfill-go, run the following commmands:

```
go build
./builder -useragent | gofmt > ../useragent/parse.go
./builder -normalise | gofmt > ../useragent/normalise.go
```