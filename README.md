# nudata-solver

A solver for NuData Security

## Description

Generates `nds-pmd` payload to bypass NuData Security.

## Supported

For now, only `iOS` is supported.

And the following sites for it:

- `Kohls` - `2.7.5`

## Installation
1. `git clone github.com/brianxor/nudata-solver.git`
2. `cd nudata-solver`
3. `go run .`

## Usage

```
http://127.0.0.1:8000/nudata/solve

{"websiteName":"kohls","proxy":"http://user:pass@ip:port"}
```

> [!NOTE]
> You can configure the server host & port through the .env file
