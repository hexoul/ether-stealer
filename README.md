# Ethereum Stealer
[![License](http://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/hexoul/ether-stealer/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/hexoul/ether-stealer)](https://goreportcard.com/report/github.com/hexoul/ether-stealer)
[![GoDoc](https://godoc.org/github.com/hexoul/ether-stealer?status.svg)](https://godoc.org/github.com/hexoul/ether-stealer)

> This project conducts account mining, not block mining. In general, this project can do nothing stochastically. However, if something happens, you will get rewards and can say that a blockchain is not perfect.

## Contents
- [Installation](#installation)
- [Getting started](#getting-started)
- [Test](#test)
- [License](#license)
- [Reference](#reference)


## Installation
```shell
go get -u github.com/hexoul/ether-stealer
```
- If you want cross-compile, type ```make``` which uses xgo[1]


## Getting started
```shell
go run main.go \
  -concurrency 10
  -infura-apikey [infura_apikey] \
  -telegram-chatid [telegram_chat_id] \
  -telegram-apikey [telegram_apikey] \
  -identifier [name]
```
- Options are all optional. No need to put all.
- `concurrency`: The number of threads can be executed concurrently.
- `infura-apikey`: API key of your Infura project.
- `identifier`: An identifier of a client.


## Test
```shell
go test -v
go test -v ./infura -args -infura-apikey xxx
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
