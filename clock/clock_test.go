package clock

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	past   = time.Now().Add(-time.Hour)
	future = time.Now().Add(time.Hour)
)

func Test_system_Now(t *testing.T) {
	t.Run("now returns time.now", func(t *testing.T) {
		s := system{}
		assert.WithinDurationf(t, time.Now(), s.Now(), time.Millisecond, "Now()")
	})
}

func Test_system_Since(t *testing.T) {
	t.Run("since returns time.since", func(t *testing.T) {
		s := system{}
		assert.InDeltaf(t, time.Since(past), s.Since(past), float64(time.Microsecond), "Since(%v)", past)
	})
}

func Test_system_Until(t *testing.T) {
	t.Run("until returns time.until", func(t *testing.T) {
		s := system{}
		assert.InDeltaf(t, time.Until(future), s.Until(future), float64(time.Microsecond), "Until(%v)", future)
	})
}

func TestFixed(t *testing.T) {
	type args struct {
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want Clock
	}{
		{
			name: "time passed, returns fixed clock with time",
			args: args{past},
			want: &fixed{past},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Fixed(tt.args.t), "Fixed(%v)", tt.args.t)
		})
	}
}

func Test_fixed_Now(t *testing.T) {
	type fields struct {
		fixed time.Time
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Time
	}{
		{
			name:   "now returns fixed time",
			fields: fields{past},
			want:   past,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fixed{
				fixed: tt.fields.fixed,
			}
			assert.Equalf(t, tt.want, f.Now(), "Now()")
		})
	}
}

func Test_fixed_Since(t *testing.T) {
	type fields struct {
		fixed time.Time
	}
	type args struct {
		t time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   time.Duration
	}{
		{
			name:   "since returns duration since fixed time",
			fields: fields{future},
			args:   args{past},
			want:   future.Sub(past),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fixed{
				fixed: tt.fields.fixed,
			}
			assert.Equalf(t, tt.want, f.Since(tt.args.t), "Since(%v)", tt.args.t)
		})
	}
}

func Test_fixed_Until(t *testing.T) {
	type fields struct {
		fixed time.Time
	}
	type args struct {
		t time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   time.Duration
	}{
		{
			name:   "since returns duration until fixed time",
			fields: fields{past},
			args:   args{future},
			want:   future.Sub(past),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := fixed{
				fixed: tt.fields.fixed,
			}
			assert.Equalf(t, tt.want, f.Until(tt.args.t), "Until(%v)", tt.args.t)
		})
	}
}
