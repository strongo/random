package random

import (
	"regexp"
	"strconv"
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
