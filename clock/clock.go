package clock

import "time"

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

func System() Clock {
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

func Fixed(t time.Time) Clock {
	return &fixed{t}
}
