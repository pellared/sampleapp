# sampleapp

[![GitHub Release](https://img.shields.io/github/v/release/pellared/sampleapp)](https://github.com/pellared/sampleapp/releases)
[![go.dev](https://img.shields.io/badge/go.dev-reference-blue.svg)](https://pkg.go.dev/github.com/pellared/sampleapp)
[![go.mod](https://img.shields.io/github/go-mod/go-version/pellared/sampleapp)](go.mod)
[![Build Status](https://img.shields.io/github/workflow/status/pellared/sampleapp/build)](https://github.com/pellared/sampleapp/actions?query=workflow%3Abuild+branch%3Amaster)
[![Go Report Card](https://goreportcard.com/badge/github.com/pellared/sampleapp)](https://goreportcard.com/report/github.com/pellared/sampleapp)

## Build

- Terminal: `make` to get help for make targets.
- Terminal: `make all` to execute a full build.
- Visual Studio Code: `Terminal` → `Run Build Task... (CTRL+ALT+B)` to execute a fast build.

## Maintainance

Remember to update Go version in [.github/workflows](.github/workflows), [Makefile](Makefile) and [devcontainer.json](.devcontainer/devcontainer.json).

Notable files:

- [devcontainer.json](.devcontainer/devcontainer.json) - Visual Studio Code Remote Container configuration
- [.github/workflows](.github/workflows) - GitHub Actions workflows
- [.vscode](.vscode) - Visual Studio Code configuration files
- [.golangci.yml](.golangci.yml) - golangci-lint configuration
- [.goreleaser.yml](.goreleaser.yml) - GoReleaser configuration
- [install.sh](install.sh) - build tools installation script
- [Makefile](Makefile) - Make targets used for development, [CI build](.github/workflows) and [.vscode/tasks.json](.vscode/tasks.json)
