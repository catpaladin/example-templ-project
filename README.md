# example-templ-project

This project is an example to teach templ. Yes, it is cat themed.

## Requirements

This project is intended for Go version `1.24`+

```bash
# Install the Templ CLI
go install github.com/a-h/templ/cmd/templ@latest

# Install the Air CLI
go install github.com/air-verse/air@latest
```

```
my-templ-project/
├── cmd
│   ├── composition-example
│   │   └── main.go
│   ├── conditional-example
│   │   └── main.go
│   ├── greeting-example
│   │   └── main.go
│   └── iteration-example
│       └── main.go
├── go.mod
├── go.sum
├── internal
│   ├── components
│   │   └── layout_templ.go
│   └── views
│       ├── composition_templ.go
│       ├── conditional_templ.go
│       ├── greeting_templ.go
│       └── iteration_templ.go
├── Makefile
├── README.md
├── scripts
│   └── generate-templ.sh
└── templates
    ├── components
    │   └── layout.templ
    └── views
        ├── composition.templ
        ├── conditional.templ
        ├── greeting.templ
        └── iteration.templ

```
## Making templ changes

Make modifications to any `.templ` file and run:

```bash
make generate
```

## Greeting Example

```bash
make dev CMD_DIR=greeting-example
```

## Composition Example

```bash
make dev CMD_DIR=composition-example
```

## Conditional Example

```bash
make dev CMD_DIR=conditional-example
```

## Iteration Example

```bash
make dev CMD_DIR=iteration-example
```

