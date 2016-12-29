# walker-wormkit

[![standard-readme compliant](https://img.shields.io/badge/standard--readme-OK-green.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)

> An educational toolkit for worm creation.

## Table of Contents

- [Security](#security)
- [Install](#install)
- [Usage](#usage)
    - [Overview](#overview)
    - [Patchers](#patchers)
    - [File Associations](#file-associations)
    - [Ignored Paths](#ignored-paths)
- [Compatibility](#compatibility)
- [Contribute](#contribute)
- [License](#license)

## Security

Always remember that **whatever is wrapped in this toolkit may or may not**, depending on your modifications, **damage your OS and must be treated as hazardous.**

It is advisable to run all the operations inside a VM, a container, or another kind of secluded system. A Vagrantfile is included for ease of use, and to ensure the total safety of your host system you should unmount the `/vagrant` share in the guest, making the VM effectively untied from your host enviroment.

If you choose to use alternative restriction means, be sure to act accordingly.

### Disclaimer

The author(s) of this toolkit hold no responsibility or liability for any consequences that the use of it might cause to your computer, cat, microwave, house or pocket spaceship.

Use is expected from people who know what they are doing; documentation is purposefully scarce in general details.

> _And on the Eighth day, the Lord made a backup copy._

## Install

To download all the needed dependencies `cd` into the repo's folder and run:

```bash
go get -d ./...
```

## Usage

### Overview

_Walker_ is a basic toolkit that traverses the filesystem and executes a patcher function against any writable file with a given extension that is found. The patchers are modular, pluggable and must not follow any specific behavior apart from having a specific signature.

### Patchers

Patcher functions must **belong to the `patchers` package** and **have the following signature**, where `path` is the full absulte path to the writable file: 

```go
func LanguageName(path string) bool
```

The patcher must return `true` if the given file was successfully processed or `false` if, for any intentional or unintentional reason, it was not correctly or fully processed.

The files containing such functions must be placed in the [patchers](https://github.com/nmaggioni/walker-wormkit/tree/master/patchers) directory.

### File Associations

Patchers are associated to file extensions in the `filePatchers` map in the [main.go](https://github.com/nmaggioni/walker-wormkit/blob/master/main.go) file. The map has a `map[string]func(string) bool` signature.

_Check the code for examples._

### Ignored Paths

If you want Walker to avoid stepping inside a certain folder at all, place its absolute path in the `ignoredPaths` slice in the [main.go](https://github.com/nmaggioni/walker-wormkit/blob/master/main.go) file.

_Check the code for examples._

## Compatibility

The current implementation of the toolkit is based on the UNIX filesystem model (Linux + Mac OS), but some core functions were **untestedly** ported to Windows.

Executing it as-is will probably not work coherently on Windows - you should at least check the ignored paths.

## Contribute

PRs are gladly accepted.

Small note: If editing the README, please conform to the [standard-readme](https://github.com/RichardLitt/standard-readme) specification.

## License

_GPLv3_ - the [LICENSE](https://github.com/nmaggioni/walker-wormkit/blob/master/LICENSE) file is the source of truth.
