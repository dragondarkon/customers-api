## go dependencies

```bash
go mod init
go mod tidy
```

## Run

```bash
go run main.go
```

## APIs

- `GET : /api/customers`

  response body

  ```json
    [{
        "id": "1234",
        "name": "Test",
        "age": 18
    }]
  ```

- `GET : /api/customers/1234`

  response body

  ```json
    [{
        "id": "1234",
        "name": "Test",
        "age": 18
    }]
  ```

- `POST : /api/customers`

  request body

  ```json
    [{
        "id": "1234",
        "name": "Test",
        "age": 18
    }]
  ```

  response body

  ```json
    [{
        "id": "1234",
        "name": "Test",
        "age": 18
    }]
  ```

- `PUT : /api/customers`

  request body

  ```json
    [{
        "id": "1234",
        "name": "Test",
        "age": 20
    }]
  ```

  response body

  ```json
    [{
        "id": "1234",
        "name": "Test",
        "age": 20
    }]
  ```

  - `DELETE : /api/customers/1234`

  response body

  ```json

  ```

## Unit Test

test

```bash
cd apis
```

```bash
go test -v
```
```bash
go test -cover
```

test with create cover profile

```bash
go test -coverprofile coverage.out ./...
```

export report html

```bash
go tool cover -html coverage.out -o report.html
```
