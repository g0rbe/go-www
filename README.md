# www

[![Go Reference](https://pkg.go.dev/badge/git.gorbe.io/go/www.svg)](https://pkg.go.dev/git.gorbe.io/go/www)
[![Go Report Card](https://goreportcard.com/badge/git.gorbe.io/go/www)](https://goreportcard.com/report/git.gorbe.io/go/www)

Package `www` provides HTTP utility functions.

## Get
```bash
go get git.gorbe.io/go/www@latest
```

Get the latest commit (if Go module proxy is not updated):

```bash
go get "git.gorbe.io/go/www@$(curl -s 'https://git.gorbe.io/api/v1/repos/go/www/commits' | jq -r '.[0].sha')"
```