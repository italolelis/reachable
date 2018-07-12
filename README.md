# Reachable

> A CLI tool to check if a domain is up

[![Build Status](https://travis-ci.org/italolelis/reachable.svg)](https://travis-ci.org/italolelis/reachable) [![Go Report Card](https://goreportcard.com/badge/github.com/italolelis/reachable)](https://goreportcard.com/report/github.com/italolelis/reachable) [![GoDoc](https://godoc.org/github.com/italolelis/reachable?status.svg)](https://godoc.org/github.com/italolelis/reachable)

Welcome to reachable a CLI tool that helps you to check if a domain is up or down.

## Installation

### MacOS

You can use our homebrew tap to install it

```sh
brew tep italolelis/homebrew-reachable
brew install reachable
```

### Manual

Go the [releases](https://github.com/italolelis/reachable/releases) and download the latest one for your platform.
Just place the binary in your $PATH and you are good to go.

## Usage

```
reachable [command] [--flags]
``` 

### Commands

| Command                  | Description                          |
|--------------------------|--------------------------------------|
| `reachable check [--flags]`   | Checks if a domain is reachable |
| `reachable version`           | Prints the version information  |

# Contributing

To start contributing, please check [CONTRIBUTING](CONTRIBUTING)
