<p align="center">
  <img alt="nem-toolchain logo" src="assets/logo.png" height="128" />
  <h3 align="center"><a href="https://git.io/nemtool">nem-toolchain</a></h3>
  <p align="center">Command line toolchain for <a href=https://nem.io>NEM blockchain</a>.</p>
  <p align="center">
    <a href="https://gitter.im/nem-toolchain/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge"><img alt="Join the chat at https://gitter.im/nem-toolchain/Lobby" src="https://img.shields.io/gitter/room/badges/shields.svg?style=flat-square"></a>
    <a href="https://circleci.com/gh/nem-toolchain/nem-toolchain"><img alt="CircleCI" src="https://img.shields.io/circleci/project/github/nem-toolchain/nem-toolchain/master.svg?style=flat-square"></a>
    <a href="https://travis-ci.org/nem-toolchain/nem-toolchain"><img alt="Travis" src="https://img.shields.io/travis/nem-toolchain/nem-toolchain/master.svg?style=flat-square"></a>
    <a href="https://ci.appveyor.com/project/dulanov/nem-toolchain"><img alt="AppVeyor" src="https://img.shields.io/appveyor/ci/dulanov/nem-toolchain/master.svg?style=flat-square"></a>
    <a href="https://codecov.io/gh/nem-toolchain/nem-toolchain"><img alt="Coverage Status" src="https://img.shields.io/codecov/c/github/nem-toolchain/nem-toolchain/master.svg?style=flat-square"></a>
    <a href="https://github.com/nem-toolchain/nem-toolchain/releases/latest"><img alt="Release" src="https://img.shields.io/github/release/nem-toolchain/nem-toolchain.svg?style=flat-square"></a>
    <a href="http://godoc.org/github.com/nem-toolchain/nem-toolchain"><img alt="Go Doc" src="https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square"></a>
    <a href="https://goreportcard.com/report/github.com/nem-toolchain/nem-toolchain"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/nem-toolchain/nem-toolchain?style=flat-square"></a>
    <a href="LICENSE"><img alt="Software License" src="https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square"></a>
  </p>
</p>

---

## How to try 

Run docker  container with interactive mode (will be deletaed after shutdown):

```console
$ docker run --rm -it nem-toolchain/latest /bin/sh
```

## How to install

Prerequisites:  [Go 1.8+](http://golang.org/doc/install)

To install, simply run:

```console
$ go get github.com/nem-toolchain/nem-toolchain/cmd/nem
```

Make sure your `PATH` includes the `$GOPATH/bin` directory so your commands can be easily used:

```console
$ export PATH=$PATH:$GOPATH/bin
```

Verify development snapshot installation with:

```console
$ nem -v
nem version git
```

If you look for releases instead, there are already prepared
[tarball releases](https://github.com/nem-toolchain/nem-toolchain/releases/latest).

## How to contribute

You are always welcome to fix
[not yet assigned bugs](https://github.com/nem-toolchain/nem-toolchain/issues?q=is%3Aopen+label%3A%22non-critical%20bug%22+no%3Aassignee)
or can help us with [extraordinary issues](https://github.com/nem-toolchain/nem-toolchain/labels/help%20wanted).

Check out the [technical documentation](https://github.com/nem-toolchain/nem-toolchain/wiki),
there you can find project roadmap, implementation details, useful links and another development related info.
In addition please don't forget to read and accept the [contributing guide](CONTRIBUTING.md).

## How to get support

Check out [our website](https://git.io/nemtool) for more usage specific details,
or chat with us in the [Gitter chat](https://gitter.im/nem-toolchain/Lobby).

## Donations

Our NEM address:
[`NCSY3D-EVBGVX-MEPZWK-TIJWCA-OXL5NE-OLJPHX-LFLD`](http://chain.nem.ninja/#/account/NCSY3D-EVBGVX-MEPZWK-TIJWCA-OXL5NE-OLJPHX-LFLD)
 
## Thanks to

* [JetBrains](https://www.jetbrains.com) for free open-source license for
[IntelliJ IDEA Ultimate](https://www.jetbrains.com/idea) and [GoLand](https://www.jetbrains.com/go/).

## Licence

[MIT](LICENSE)
