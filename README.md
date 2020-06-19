[![Tag](https://img.shields.io/github/v/tag/moore0n/gopr?sort=date)](https://github.com/moore0n/gopr/releases)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/moore0n/gopr)](https://golang.org/dl/)
[![GitHub stars](https://img.shields.io/github/stars/moore0n/gopr?style=social)](https://github.com/moore0n/gopr/stargazers)

# gopr
gopr is a simple CLI tool for opening the pull request url for github.com based on the current branch.

# Usage
```
NAME:
   gopr - Open the link for creating a pull request in Github

USAGE:
   gopr [options...] <target>

VERSION:
   1.0.0

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --title value   The title for your pull request
   --body value    The body for your pull request
   --remote value  The name of the remote to target (default: "origin")
   --help, -h      show help (default: false)
   --version, -v   print the version (default: false)
```

## Install 
```
go get -u github.com/moore0n/gopr
go install github.com/moore0n/gopr/...
```

## Try
```
gopr --title "New Feature" --body "This PR introduces a new feature" --remote origin master
```

## Build
If you so choose you can build a binary locally using the supplied build command.
```
$ make mac
```
