# Clock

Clock is an interface to provide current time in a way that can easily be mocked.

## Usage

```go
package main

import (
	"fmt"
	"github.com/ellogroup/ello-golang-clock/clock"
	"time"
)

func main() {
	// Returns a system clock
	system := clock.NewSystem()
	fmt.Printf("System time: %s\n\n", system.Now().String())

	// Returns a fixed clock
	fixed := clock.NewFixed(time.Date(2030, 1, 2, 12, 34, 56, 0, time.UTC))
	fmt.Printf("Fixed time: %s\n\n", fixed.Now().String())
}
```

### Interface

The `Clock` interface includes methods related to the current time.

```go
type Clock interface {
	Now() time.Time
	Since(t time.Time) time.Duration
	Until(t time.Time) time.Duration
}
```

### Implementations

#### System

Returns the system clock, which uses `time.Now()`.

```go
system := clock.NewSystem()
```

#### Fixed

Returns a fixed clock, which always returns the specified time.

```go
fixed := clock.NewFixed(time.Date(2030, 1, 2, 12, 34, 56, 0, time.UTC))
```