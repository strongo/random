package random

import (
	"regexp"
	"strconv"
	"sync"
	"testing"
)

var charsRegex = regexp.MustCompile("[a-zA-Z0-9]+")

func TestID(t *testing.T) {
	for _, l := range []int{1, 9, 15} {
		t.Run(strconv.Itoa(l), func(t *testing.T) {
			v := ID(l)
			if len(v) != l {
				t.Fatalf("len(id) != %v: %v", l, len(v))
			}
			if !charsRegex.MatchString(v) {
				t.Errorf("unexpected character(s): %v", v)
			}
		})
	}
}

func TestDigits(t *testing.T) {
	for _, l := range []int{1, 9, 15} {
		t.Run(strconv.Itoa(l), func(t *testing.T) {
			v := Digits(l)
			if len(v) != l {
				t.Fatalf("len(id) != %v: %v", l, len(v))
			}
			if _, err := strconv.Atoi(v); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestGeneratorsAreSafeForConcurrentUse(t *testing.T) {
	const (
		goroutines = 16
		iterations = 100
		length     = 32
	)

	var wg sync.WaitGroup
	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				if got := ID(length); len(got) != length || !charsRegex.MatchString(got) {
					t.Errorf("ID() = %q, want %d alphanumeric characters", got, length)
				}
				if got := Digits(length); len(got) != length {
					t.Errorf("len(Digits()) = %d, want %d", len(got), length)
				}
			}
		}()
	}
	wg.Wait()
}
