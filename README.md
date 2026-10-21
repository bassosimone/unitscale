# Unit Scaling and Formatting

[![GoDoc](https://pkg.go.dev/badge/github.com/bassosimone/unitscale)](https://pkg.go.dev/github.com/bassosimone/unitscale) [![Build Status](https://github.com/bassosimone/unitscale/actions/workflows/go.yml/badge.svg)](https://github.com/bassosimone/unitscale/actions) [![codecov](https://codecov.io/gh/bassosimone/unitscale/branch/main/graph/badge.svg)](https://codecov.io/gh/bassosimone/unitscale)

The `unitscale` Go package contains helpers for formatting and
scaling numeric values using IEC and SI prefixes.

```Go
import (
	"fmt"

	"github.com/bassosimone/unitscale"
)

fmt.Println(unitscale.IEC(3 * 1024 * 1024, "B"))
fmt.Println(unitscale.SI(10e6, "bit/s"))
```

## Installation

To add this package as a dependency to your module:

```sh
go get github.com/bassosimone/unitscale
```

## Development

To run the tests:
```sh
go test -v .
```

To measure test coverage:
```sh
go test -v -cover .
```

## License

```
SPDX-License-Identifier: GPL-3.0-or-later
```

## History

Adapted from [ooni/probe-cli/internal/humanize](https://github.com/ooni/probe-cli/blob/647b03f4270eb758106523fe6273e3ebdbcd599c/internal/humanize/humanize.go).
