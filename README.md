# Ethereum Stealer

[![GoDoc](https://godoc.org/github.com/hexoul/ether-stealer?status.svg)](https://godoc.org/github.com/hexoul/ether-stealer)

## Contents
- [Build](#build)
- [Run](#run)
- [Test](#test)
- [License](#license)

## Build
```shell
dep ensure
go build
```

### Run
```shell
go run main.go -limiter=10 -apikey=[telegram_apikey] -chatid=[telegram_chat_id]
```

## Test
1. Move each module directory
2. Run testunit
```shell
go test -v
```

## License
MIT
