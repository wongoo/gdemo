// authors: wangoo
// created: 2018-05-25
// time demo

package basic

import (
	"fmt"
	"time"
	"testing"
)

func TestTime(t *testing.T) {
	timeDemo()

	durationDemo()

}
func timeDemo() {
	p := fmt.Println

	// We'll start by getting the current time.
	now := time.Now()
	p(now)

	// You can build a `time` struct by providing the
	// year, month, day, etc. Times are always associated
	// with a `Location`, i.e. time zone.
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// You can extract the various components of the time
	// value as expected.
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Unix())
	p(then.UnixNano())
	p(then.Location())

	p(then.Format("3:04PM"))
	p(then.Format("Mon Jan _2 15:04:05 2006"))
	p(then.Format("2006-01-02T15:04:05.999999-07:00"))
	p(then.Format("2006-01-02T15:04:05Z07:00"))
	p(then.Format("2006-01-02 15:04:05"))
	p(then.Format("2006-01-02"))
	p(then.Format("15:04:05"))

	// The Monday-Sunday `Weekday` is also available.
	p(then.Weekday())

	// These methods compare two times, testing if the
	// first occurs before, after, or at the same time
	// as the second, respectively.
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// The `Sub` methods returns a `Duration` representing
	// the interval between two times.
	diff := now.Sub(then)
	p(diff)

	// We can compute the length of the duration in
	// various units.
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// You can use `Add` to advance a time by a given
	// duration, or with a `-` to move backwards by a
	// duration.
	p(then.Add(diff))
	p(then.Add(-diff))

	layout := "2006-01-02 15:04:05"
	t, _ := time.Parse(layout, "2013-10-05 18:30:50")
	p(t.Year())
}

func durationDemo() {
	p := fmt.Println

	t, _ := time.ParseDuration("24h")
	p(t.Hours())
	t, _ = time.ParseDuration("1h")
	p(t.Minutes())
	t, _ = time.ParseDuration("1m")
	p(t.Seconds())

	now := time.Now()
	duration := time.Duration(-60 * time.Second)
	old := now.Add(duration)

	diff := time.Now().Sub(old)
	p(int(diff.Seconds()))

	du := time.Since(old)
	p(du)

}
