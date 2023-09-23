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

### Listing available assets from the latest release

```bash
target list --repo='haproxytech/dataplaneapi'
```

Output:

```bash
checksums.txt
dataplaneapi_2.8.1_darwin_arm64.tar.gz
dataplaneapi_2.8.1_darwin_x86_64.tar.gz
dataplaneapi_2.8.1_freebsd_arm.tar.gz
dataplaneapi_2.8.1_freebsd_arm64.tar.gz
dataplaneapi_2.8.1_freebsd_i386.tar.gz
dataplaneapi_2.8.1_freebsd_x86_64.tar.gz
dataplaneapi_2.8.1_linux_amd64.apk
dataplaneapi_2.8.1_linux_amd64.deb
dataplaneapi_2.8.1_linux_amd64.rpm
dataplaneapi_2.8.1_linux_arm.apk
dataplaneapi_2.8.1_linux_arm.deb
dataplaneapi_2.8.1_linux_arm.rpm
dataplaneapi_2.8.1_linux_arm.tar.gz
dataplaneapi_2.8.1_linux_arm64.apk
dataplaneapi_2.8.1_linux_arm64.deb
dataplaneapi_2.8.1_linux_arm64.rpm
dataplaneapi_2.8.1_linux_arm64.tar.gz
dataplaneapi_2.8.1_linux_i386.apk
dataplaneapi_2.8.1_linux_i386.deb
dataplaneapi_2.8.1_linux_i386.rpm
dataplaneapi_2.8.1_linux_i386.tar.gz
dataplaneapi_2.8.1_linux_ppc64le.apk
dataplaneapi_2.8.1_linux_ppc64le.deb
dataplaneapi_2.8.1_linux_ppc64le.rpm
dataplaneapi_2.8.1_linux_ppc64le.tar.gz
dataplaneapi_2.8.1_linux_s390x.apk
dataplaneapi_2.8.1_linux_s390x.deb
dataplaneapi_2.8.1_linux_s390x.rpm
dataplaneapi_2.8.1_linux_s390x.tar.gz
dataplaneapi_2.8.1_linux_x86_64.tar.gz
```

By default output formatting shows just a name of an asset.
It is possible to specify custom formatting for the output.

```bash
target list --repo 'haproxytech/dataplaneapi' --format='{{printf "%s\t\t%s" .Name .BrowserDownloadURL}}'
```

Output:

```bash
checksums.txt		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/checksums.txt
dataplaneapi_2.8.1_darwin_arm64.tar.gz		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/dataplaneapi_2.8.1_darwin_arm64.tar.gz
dataplaneapi_2.8.1_darwin_x86_64.tar.gz		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/dataplaneapi_2.8.1_darwin_x86_64.tar.gz
dataplaneapi_2.8.1_freebsd_arm.tar.gz		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/dataplaneapi_2.8.1_freebsd_arm.tar.gz
dataplaneapi_2.8.1_freebsd_arm64.tar.gz		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/dataplaneapi_2.8.1_freebsd_arm64.tar.gz
dataplaneapi_2.8.1_freebsd_i386.tar.gz		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/dataplaneapi_2.8.1_freebsd_i386.tar.gz
dataplaneapi_2.8.1_freebsd_x86_64.tar.gz		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/dataplaneapi_2.8.1_freebsd_x86_64.tar.gz
dataplaneapi_2.8.1_linux_amd64.apk		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/dataplaneapi_2.8.1_linux_amd64.apk
dataplaneapi_2.8.1_linux_amd64.deb		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/dataplaneapi_2.8.1_linux_amd64.deb
dataplaneapi_2.8.1_linux_amd64.rpm		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/dataplaneapi_2.8.1_linux_amd64.rpm
dataplaneapi_2.8.1_linux_arm.apk		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/dataplaneapi_2.8.1_linux_arm.apk
dataplaneapi_2.8.1_linux_arm.deb		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/dataplaneapi_2.8.1_linux_arm.deb
dataplaneapi_2.8.1_linux_arm.rpm		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/dataplaneapi_2.8.1_linux_arm.rpm
dataplaneapi_2.8.1_linux_arm.tar.gz		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/dataplaneapi_2.8.1_linux_arm.tar.gz
dataplaneapi_2.8.1_linux_arm64.apk		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/dataplaneapi_2.8.1_linux_arm64.apk
dataplaneapi_2.8.1_linux_arm64.deb		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/dataplaneapi_2.8.1_linux_arm64.deb
dataplaneapi_2.8.1_linux_arm64.rpm		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/dataplaneapi_2.8.1_linux_arm64.rpm
dataplaneapi_2.8.1_linux_arm64.tar.gz		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/dataplaneapi_2.8.1_linux_arm64.tar.gz
dataplaneapi_2.8.1_linux_i386.apk		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/dataplaneapi_2.8.1_linux_i386.apk
dataplaneapi_2.8.1_linux_i386.deb		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/dataplaneapi_2.8.1_linux_i386.deb
dataplaneapi_2.8.1_linux_i386.rpm		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/dataplaneapi_2.8.1_linux_i386.rpm
dataplaneapi_2.8.1_linux_i386.tar.gz		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/dataplaneapi_2.8.1_linux_i386.tar.gz
dataplaneapi_2.8.1_linux_ppc64le.apk		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/dataplaneapi_2.8.1_linux_ppc64le.apk
dataplaneapi_2.8.1_linux_ppc64le.deb		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/dataplaneapi_2.8.1_linux_ppc64le.deb
dataplaneapi_2.8.1_linux_ppc64le.rpm		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/dataplaneapi_2.8.1_linux_ppc64le.rpm
dataplaneapi_2.8.1_linux_ppc64le.tar.gz		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/dataplaneapi_2.8.1_linux_ppc64le.tar.gz
dataplaneapi_2.8.1_linux_s390x.apk		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/dataplaneapi_2.8.1_linux_s390x.apk
dataplaneapi_2.8.1_linux_s390x.deb		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/dataplaneapi_2.8.1_linux_s390x.deb
dataplaneapi_2.8.1_linux_s390x.rpm		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/dataplaneapi_2.8.1_linux_s390x.rpm
dataplaneapi_2.8.1_linux_s390x.tar.gz		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/dataplaneapi_2.8.1_linux_s390x.tar.gz
dataplaneapi_2.8.1_linux_x86_64.tar.gz		https://github.com/haproxytech/dataplaneapi/releases/download/v2.8.1/dataplaneapi_2.8.1_linux_x86_64.tar.gz
```

### Get asset file

Download asset file to current directory:

```bash
target get --repo 'haproxytech/dataplaneapi' --file='dataplaneapi_2.8.1_linux_x86_64.tar.gz'
```