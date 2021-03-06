# How to contribute

If you would like to contribute code you can do so through GitHub by forking
the repository and open a pull request against the master branch.

This project adheres to the Contributor Covenant [code of conduct](CODE_OF_CONDUCT.md).
By participating, you are expected to uphold this code.

## How to build

Prerequisites are:

* `make`
* [`Go 1.9+`](http://golang.org/doc/install)

Install the build and lint dependencies:

``` sh
$ make setup
```

A good way of making sure everything is all right is running the test suite:

``` sh
$ make test
```

## Test your change

You can create a branch for your changes and try to build from the source as you go:

``` sh
$ make build
```

Make sure your `PATH` includes the `bin` directory so your local builds can be easily used:

```bash
export PATH=$PATH:./bin
```

Simplest first test after build is:

```bash
nem -h
```

When you are satisfied with the changes, we suggest you run:

``` sh
$ make fmt && make ci
```

Which runs all the formatters, linters and tests.

## Special commit prefixes

Commits begin with `Fix`, `Merge` and `Update` will be excluded from the release notes
automatically, as they do not relate directly to development, but are more about supporting.

## License

By contributing your code, you agree to license your contribution under the terms of the
[MIT License](LICENSE).

If you are adding a new file it should have a header like this:

```
// Copyright 2017 The nem-toolchain project authors. All rights reserved.
// Use of this source code is governed by a MIT license that can be found in the LICENSE file.
 ```
