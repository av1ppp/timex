# Timex

Timex is a Go package that provides utilities for working with time durations and time zones. It is designed to provide convenience functions and types for handling time-related operations.

## Installation

To install the package, use the following command:

```bash
go get github.com/av1ppp/timex
```

## Usage

Here are some examples of how to use the Timex package:

### Working with Duration

```go
package main

import (
	"fmt"
	"github.com/av1ppp/timex"
)

func main() {
	// Create a duration of 1 hour, 2 minutes, and 3 seconds
	d := timex.Hour + timex.Minute*2 + timex.Second*3

	// Convert duration to string
	fmt.Println(d.String()) // Output: 01:02:03
}
```

## License

This project is licensed under the MIT License.
