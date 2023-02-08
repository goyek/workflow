# goyek reusable workflow

[![Go Reference](https://pkg.go.dev/badge/github.com/goyek/workflow.svg)](https://pkg.go.dev/github.com/goyek/workflow)
[![go.mod](https://img.shields.io/github/go-mod/go-version/goyek/workflow)](go.mod)
[![LICENSE](https://img.shields.io/github/license/goyek/workflow)](LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/goyek/workflow/build.yml?branch=main)](https://github.com/goyek/workflow/actions?query=workflow%3Abuild+branch%3Amain)
[![Go Report Card](https://goreportcard.com/badge/github.com/goyek/workflow)](https://goreportcard.com/report/github.com/goyek/workflow)

⭐ `Star` this repository if you find it valuable and worth maintaining.

## Description

`workflow` is an example of a reusable, yet customizable, `goyek` workflow
that could be used to share common build pipeline code.

It defines reusable tasks, stages and configures the `goyek.DefaultFlow`
to use common flags, middlewares, printing.

You can find example usage (and customization) in <https://github.com/goyek/demo>.

## Build Pipeline

The default `all` build pipeline consists of following stages:

- `init`
- `build`
- `test`

### Stages

The `init` stage consists of following tasks:

- `go-tidy`
- `go-generate`

The `build` stage consists of following tasks:

- `go-build`

The `test` stage consists of following tasks:

- `go-vet`
- `go-test`

## Additional Tasks

`clean` tasks is used to clean the workspace from files created during build pipeline.

`git-diff` is used to detect if there were changes (e.g. done by `go generate`)
that were not commited.
