# Reachable

[![Build Status](https://travis-ci.com/italolelis/reachable.svg)](https://travis-ci.com/italolelis/reachable) [![Coverage Status](https://coveralls.io/repos/github/italolelis/reachable/badge.svg?branch=master)](https://coveralls.io/github/italolelis/reachable?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/italolelis/reachable)](https://goreportcard.com/report/github.com/italolelis/reachable) [![GoDoc](https://godoc.org/github.com/italolelis/reachable?status.svg)](https://godoc.org/github.com/italolelis/reachable)

> A CLI tool to check if a domain is up

Welcome to reachable a CLI tool that helps you to check if a domain is up or down.

<p align="center">
<a href="https://asciinema.org/a/LJSooVahoiopp9Vx6THTEnwnP" target="_blank"><img src="https://asciinema.org/a/LJSooVahoiopp9Vx6THTEnwnP.png" width="70%"/></a>
</p>

## Installation

### MacOS

You can use our homebrew tap to install it

```sh
brew tap italolelis/homebrew-reachable
brew install reachable
```

### Manual

Go the [releases](https://github.com/italolelis/reachable/releases) and download the latest one for your platform.
Just place the binary in your $PATH and you are good to go.

## Usage

```
reachable [command] [--flags]
``` 

A few examples 

```
$ reachable check google.com

$ reachable check google.com facebook.com twitter.com

$ reachable check google.com -v
```

### Commands

| Command                  | Description                          |
|--------------------------|--------------------------------------|
| `reachable check [--flags]`   | Checks if a domain is reachable |
| `reachable version`           | Prints the version information  |

# Contributing

To start contributing, please check [CONTRIBUTING](CONTRIBUTING)
