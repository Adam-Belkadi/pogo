package logger

import (
	"errors"
	"reflect"
	"testing"
)

func TestLoggerDebugAndFlush(t *testing.T) {
	l := newLogger(false)

	l.Debug("first")
	l.Debug("second")

	got := l.Flush()
	want := []string{"first", "second"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Flush() = %v, want %v", got, want)
	}

	// Flush should clear internal state.
	if got := l.Flush(); len(got) != 0 {
		t.Fatalf("Flush() after clear = %v, want empty slice", got)
	}
}

func TestLoggerErrorAndFlush(t *testing.T) {
	l := newLogger(false)

	l.Error(errors.New("boom"))

	got := l.Flush()
	want := []string{"boom"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Flush() = %v, want %v", got, want)
	}
}

func TestLoggerFlushReturnsCopy(t *testing.T) {
	l := newLogger(false)

	l.Debug("alpha")
	got := l.Flush()
	got[0] = "mutated"

	// Internal state should not be affected by modifying the returned slice.
	l.Debug("beta")
	got2 := l.Flush()
	want := []string{"beta"}
	if !reflect.DeepEqual(got2, want) {
		t.Fatalf("Flush() after external mutation = %v, want %v", got2, want)
	}
}
