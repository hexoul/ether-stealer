# Ethereum Stealer
[![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/hexoul/ether-stealer/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/hexoul/ether-stealer)](https://goreportcard.com/report/github.com/hexoul/ether-stealer)
[![GoDoc](https://godoc.org/github.com/hexoul/ether-stealer?status.svg)](https://godoc.org/github.com/hexoul/ether-stealer)

> This project conducts account mining, not block mining. In general, this project can do nothing stochastically. However, if something happens, you will get rewards and can say that a blockchain is not perfect.

## Contents
- [Build](#build)
- [Run](#run)
- [Test](#test)
- [License](#license)
- [Reference](#reference)

## Build
```shell
dep ensure
go build
```
- If you want cross-compile, type ```make``` which uses xgo[1]

## Run
```shell
go run main.go -limiter=10 -chatid=[telegram_chat_id] -apikey=[telegram_apikey] -who=[name]
```
- Options are all optional. No need to put all

## Test
1. Move each module directory
1. Run testunit
```shell
go test -v
```

## Add ERC20 tokens
1. Get code at etherscan.io and put into contract/sol
2. Run abigen
 
  ```shell
  abigen -sol contract/sol/[target].sol -pkg [target] -out contract/abigen/[target]/[target].go
  ```
 
  - If you need specific solidity version,
    - Check history of [solidity.rb](https://github.com/ethereum/homebrew-ethereum/commits/master/solidity.rb)
    - Re-install solidity following a)'s commit hash
 
    ```shell
    brew unlink solidity
    brew install https://raw.githubusercontent.com/ethereum/homebrew-ethereum/[commit_hash]/solidity.rb
    ```
 
3. Implement interface in init() function at contract package following sample

## License
MIT

## Reference
[1] https://github.com/karalabe/xgo
