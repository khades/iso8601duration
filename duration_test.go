package duration

import (
	"log"
	"testing"
	"time"
)

func TestFromString(t *testing.T) {

	// test with bad format
	_, err := ParseString("asdf")
	if err != ErrBadFormat {
		t.Error("Should give exception on random string")
	}
	// test with month
	_, err = ParseString("P1M")

	if err != ErrNoMonth  {
		t.Error("Should give exception when there's month")
	}
	// test with good full string
	dur, err := ParseString("P1Y2DT3H4M5S")
	if err != nil {
		t.Error("Full string should parse properly")
	}
	fullDuration := 24*365*time.Hour + 2*24*time.Hour + 3*time.Hour + 4*time.Minute + 5*time.Second
	log.Println(dur)
	if fullDuration.Nanoseconds() != dur.Nanoseconds() {
		t.Error("Full string parsing is inaccurate")

	}

	// test with good week string
	dur, err = ParseString("P1W")
	if err != nil {
		t.Error("Week string should parse properly")
	}
	smallDuration := 7 * 24 * time.Hour
	if smallDuration.Nanoseconds() != dur.Nanoseconds() {
		t.Error("Week string parsing is inaccurate")

	}
}
