package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type CoundownOperationsSpy struct {
	Calls []string
}

const write = "write"
const sleep = "sleep"

func (s *CoundownOperationsSpy) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *CoundownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

func TestCountdown(t *testing.T) {
	t.Run("initial test", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleeper := &SpySleeper{}
		Countdown(buffer, spySleeper)
		got := buffer.String()
		want := `3
2
1
Go!`
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
		if spySleeper.Calls != 3 {
			t.Errorf("not enough calls to sleeper, want %d, got %d", spySleeper.Calls, 3)
		}
	})
	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &CoundownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)
		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}
func TestConfigurableSleeper(t *testing.T) {
	t.Run("sleeps the expected amount of time", func(t *testing.T) {
		sleepTime := 5 * time.Second
		spyTime := &SpyTime{}
		sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
		sleeper.Sleep()
		if spyTime.durationSlept != sleepTime {
			t.Errorf("should have slept for %v, but slept for %v", sleepTime, spyTime.durationSlept)
		}
	})
}
