// Package duration provides a partial implementation of ISO8601 durations. (no months)
package duration

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"text/template"
	"time"
)

var (
	// ErrBadFormat is returned when parsing fails
	ErrBadFormat = errors.New("bad format string")

	// ErrNoMonth is raised when a month is in the format string
	ErrNoMonth = errors.New("no months allowed")

	tmpl = template.Must(template.New("duration").Parse(`P{{if .Years}}{{.Years}}Y{{end}}{{if .Weeks}}{{.Weeks}}W{{end}}{{if .Days}}{{.Days}}D{{end}}{{if .HasTimePart}}T{{end }}{{if .Hours}}{{.Hours}}H{{end}}{{if .Minutes}}{{.Minutes}}M{{end}}{{if .Seconds}}{{.Seconds}}S{{end}}`))

	full = regexp.MustCompile(`P((?P<year>\d+)Y)?((?P<month>\d+)M)?((?P<day>\d+)D)?(T((?P<hour>\d+)H)?((?P<minute>\d+)M)?((?P<second>\d+)S)?)?`)
	week = regexp.MustCompile(`P((?P<week>\d+)W)`)
)

func ParseString(dur string) (*time.Duration, error) {
	var duration time.Duration
	var (
		match []string
		re    *regexp.Regexp
	)

	if week.MatchString(dur) {
		match = week.FindStringSubmatch(dur)
		re = week
	} else if full.MatchString(dur) {
		match = full.FindStringSubmatch(dur)
		re = full
	} else {
		return nil, ErrBadFormat
	}

	for i, name := range re.SubexpNames() {
		part := match[i]
		if i == 0 || name == "" || part == "" {
			continue
		}

		val, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		switch name {
		case "year":
			duration = duration + 24*365*time.Duration(val)*time.Hour
		case "month":
			return nil, ErrNoMonth
		case "week":
			duration = duration + 24*7*time.Duration(val)*time.Hour
		case "day":
			duration = duration + 24*time.Duration(val)*time.Hour
		case "hour":
			duration = duration + time.Duration(val)*time.Hour
		case "minute":
			duration = duration + time.Duration(val)*time.Minute
		case "second":
			duration = duration + time.Duration(val)*time.Second
		default:
			return nil, errors.New(fmt.Sprintf("unknown field %s", name))
		}
	}
	return &duration, nil
}
