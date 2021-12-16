# Mu Torere

[![forthebadge](http://forthebadge.com/images/badges/built-with-love.svg)](http://forthebadge.com)

Mu Torere is a game whose origins date back to the Maori tribes of New Zealand. It is a strategy game based on a low number of pawns (4 pawns per player) with 9 slots.
This program implements SARSA reinforcement algorithm to teach an AI the rules of mutorere. Once the rules have been learned, a human player can play against the AI.

## Getting Started

### Prerequisites

- You must install a recent version of [Go](https://go.dev/)

### Installation

- ```shell
  $ go get github.com/go-echarts/go-echarts/
  ```
- ```shell
  $ go get github.com/manifoldco/promptui
  ```

## Starting

To start the project, just move to the project root and run the command

```shell
$ go run cmd/mutorere/main.go
```

You can also export the application as an executable binary with the command

```shell
$ go build -o bin/mutorere cmd/mutorere/main.go
```

If you have the make command installed, you can run the commands

```shell
$ make run
```

```shell
$ make build
```

## Built with

* [go](https://github.com/golang/go) - The Go programming language
* [go-echarts](https://github.com/go-echarts/go-echarts/) - 🎨 The adorable charts library for Golang
* [promptui](https://github.com/manifoldco/promptui) - Interactive prompt for command-line applications

## Authors

* **Maël Donnart** _alias_ [@maeldonn](https://github.com/maeldonn)