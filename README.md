## Go-Challenge-Hub

Practical, bite-sized Go examples and mini-challenges covering core language building blocks: control flow, functions, syntax, error handling, and pointers/memory. Each folder contains small, focused programs you can run individually.

This repo is great for learners who prefer hands-on exploration with short, runnable snippets.

### What’s inside

- Control Flow: if/else, for loops (all forms), switch (with fallthrough, expressionless switch), real-world examples (HTTP status, file type checks, simple calculator, grading, sum of numbers)
- Functions: positional and variadic params, slice spread, named returns, single/multiple return values, basic error returns, practice problems (Area, Max)
- Go Syntax: variables, types, zero values, scope, operators — short demonstrations for fundamentals
- Error Handling: sentinel errors, custom error types, wrapping/unwrapping, practical examples (file I/O, DB errors, retry pattern)
- Pointers & Memory: pointer basics, dereferencing, zero value, pointers with funcs and structs, heap alloc via new, swap/array/slice mutation examples
- Practice Sets: small problems (Books/Shapes, FizzBuzz, Even/Odd Sum, Prime check)

## Prerequisites

- Go (the module lists `go 1.24.1`; use the latest stable Go you have — 1.22+ works for these examples)
- Windows PowerShell (repo paths below assume Windows). If you’re on another OS, adapt paths accordingly.

## Project layout

```
go.mod                  # module name: practice
Control Flow/           # if/else, loops, switch, small exercises
Error Handling/         # sentinel errors, custom types, wrapping + examples (I/O, DB, retry)
Functions/              # parameters, returns, variadic, small utilities (divide, Max)
Go Syntax/              # variables, data types, scope, operators (reference demos)
Pointers&Memory/        # pointers, memory, struct pointers, new, examples
practice_1/             # Book/Shape tasks (structs, interfaces)
practice_2/             # FizzBuzz, SumEvenOdd, PrimeNumber
```

## How to run examples (Windows PowerShell)

Run one file at a time. Many files use `package main` and are self-contained. Some files contain multiple commented examples — make sure only one `main()` is active before running.

Examples:

```powershell
# Control flow demos
go run ".\Control Flow\if_else.go"
go run ".\Control Flow\HTTP.go"
go run ".\Control Flow\switch_cases.go"   # simple calculator; follows prompts

# Functions
go run ".\Functions\function.go"
go run ".\Functions\practice.go"

# Error handling
go run ".\Error Handling\Sentinel_Errors.go"
go run ".\Error Handling\Error_Types.go"
go run ".\Error Handling\Error_Wrapping.go"
# Examples folder (file I/O, DB, retry)
go run ".\Error Handling\Examples\Examples_1.go"
go run ".\Error Handling\Examples\Example_2.go"   # requires DB driver & a DB; see notes below
go run ".\Error Handling\Examples\Example_3.go"

# Pointers & memory
go run ".\Pointers&Memory\Example_1.go"
go run ".\Pointers&Memory\example_2.go"
go run ".\Pointers&Memory\example_3.go"

# Practice sets
go run ".\practice_1\main.go"           # Book slice + PrintBookDetails
go run ".\practice_1\area.go"           # Shapes implementing Area()
go run ".\practice_2\FizzBuzz.go"
go run ".\practice_2\SumEvenOdd.go"     # prompts for n
go run ".\practice_2\PrimeNumber.go"    # prompts for number
```

Notes:

- Files that prompt for input: `switch_cases.go`, `SumEvenOdd.go`, `PrimeNumber.go`, and some others — follow the console prompts.
- Database example (`Error Handling/Examples/Example_2.go`) is illustrative. To run it:
  - Uncomment the Postgres driver import and install the driver
  - Provide a valid connection string
  - Have a database and `users` table ready

Optional driver setup:

```powershell
# Example: enable pgx and update code imports accordingly
go get github.com/jackc/pgx/v5/stdlib
```

## Folder-by-folder highlights

### Control Flow

- `if_else.go`: basic if / else-if / else chains (e.g., age categories)
- `loop.go`: all `for` loop forms — classic, while-style, infinite + `break`, and `range`
- `switch_cases.go`: several switch exercises, including a simple calculator with division-by-zero handling
- `HTTP.go`: switch over HTTP status codes and MIME types
- `practice.go`: grade calculation and sum via `range`

### Functions

- `function.go`: positional vs variadic params, slice spread (`...`), named returns, single/multiple returns, simple error returns (divide-by-zero)
- `practice.go`: `Max(...int)` implementation and example usage
- `function.txt`: primer explaining definitions, params, returns, and blank identifier

### Go Syntax

Short, didactic demos for variables, types, zero values, and scope.

Important: some files show multiple independent `main()` demos in one file. Treat them as reading/reference; if you want to run, comment out all but one `main()`.

### Error Handling

- Sentinel errors (`Sentinel_Errors.go`), custom error types (`Error_Types.go`), wrapping/unwrapping with `%w` (`Error_Wrapping.go`)
- Examples:
  - File I/O with wrapping (`Examples_1.go`)
  - DB error patterns (`Example_2.go`) — adjust driver/DSN to run
  - Retry utility wrapping final error (`Example_3.go`)

### Pointers & Memory

Pointers basics, dereferencing, passing by pointer, struct field mutation via pointer, `new` alloc, and examples that mutate arrays/slices. Includes a swap function via pointers.

### Practice Sets

- `practice_1`: Book struct with `PrintBookDetails`, plus Shapes implementing a `Shape` interface with `Area()`
- `practice_2`: FizzBuzz, sum of even/odd up to n, prime-checker

## Tips and troubleshooting

- One `main()` per build: ensure only one `main` per file is active. Comment out other demo `main()` blocks when running a file.
- Quoting paths: folders like `Control Flow` and `Pointers&Memory` contain spaces and `&`. Always quote paths in PowerShell.
- I/O utilities: some examples use `ioutil.ReadFile` (works, but modern Go favors `os.ReadFile`).
- Module name: `go.mod` sets module to `practice`. Running individual files as shown avoids cross-folder package conflicts.

## License

No license is declared yet. If you intend this to be public and reusable, consider adding an MIT license.

## Contributing

Small improvements and more micro-examples are welcome: add a new file with a single, focused `main()` and a short comment header explaining what it teaches.
