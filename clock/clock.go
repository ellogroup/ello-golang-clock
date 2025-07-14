package clock

import "time"

// Clock is an interface to provide current time in a way that can easily be mocked.
type Clock interface {
	Now() time.Time
	Since(t time.Time) time.Duration
	Until(t time.Time) time.Duration
}

type system struct{}

func (s system) Now() time.Time {
	return time.Now()
}

func (s system) Since(t time.Time) time.Duration {
	return time.Since(t)
}

func (s system) Until(t time.Time) time.Duration {
	return time.Until(t)
}

// NewSystem returns a system time implementation of Clock
func NewSystem() Clock {
	return &system{}
}

type fixed struct {
	fixed time.Time
}

func (f fixed) Now() time.Time {
	return f.fixed
}

func (f fixed) Since(t time.Time) time.Duration {
	return f.fixed.Sub(t)
}

func (f fixed) Until(t time.Time) time.Duration {
	return t.Sub(f.fixed)
}

// NewFixed returns a fixed time implementation of Clock
func NewFixed(t time.Time) Clock {
	return &fixed{t}
}
