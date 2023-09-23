# TarGET

List and download latest release asset for the given github repository

## Installation

Using Go native modules installation:

```bash
go install github.com/practical-coder/target
```

You should have `target` command available in your environment:

```bash
TarGET - download latest github project release

Usage:
  target [flags]
  target [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  get         Get / Download release asset file
  help        Help about any command
  list        List release assets
  version     TarGET version

Flags:
  -h, --help   help for target

Use "target [command] --help" for more information about a command.
```

## Description

CLI tool to list and download github release assets.
Let us take the repo of `haproxytech/dataplaneapi` as an example.

```bash
target list --repo='haproxytech/dataplaneapi'
```
