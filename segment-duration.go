package main

import (
	"fmt"
	"strings"
	"time"
)

const (
	micro  rune = '\u00B5'
	milli  rune = 'm'
	second rune = 's'
	minute rune = 'm'
	hour   rune = 'h'
)

const (
	nanoseconds  int64 = 1
	microseconds int64 = nanoseconds * 1000
	milliseconds int64 = microseconds * 1000
	seconds      int64 = milliseconds * 1000
	minutes      int64 = seconds * 60
	hours        int64 = minutes * 60
)

func segmentDuration(p *powerline) {
	if p.args.Duration == nil || *p.args.Duration == "" {
		p.appendSegment("duration", segment{
			content:    "No duration",
			foreground: p.theme.DurationFg,
			background: p.theme.DurationBg,
		})
		return
	}

	durationValue := strings.Trim(*p.args.Duration, "'\"")

	// Arguments may come in two forms, "int" and "float".  An int represents
	// just whole seconds, while a float is the number of seconds and fractions
	// of a second.

	hasPrecision := strings.Index(durationValue, ".") != -1

	d, err := time.ParseDuration(durationValue + "s")
	if err != nil {
		p.appendSegment("duration", segment{
			content:    fmt.Sprintf("Failed to convert '%s' to a number", *p.args.Duration),
			foreground: p.theme.DurationFg,
			background: p.theme.DurationBg,
		})
		return
	}

	var content string
	if hasPrecision {
		ns := d.Nanoseconds()
		if ns > hours {
			hrs := ns / hours
			ns -= hrs * hours
			mins := ns / minutes
			content = fmt.Sprintf("%dh, %dm", hrs, mins)
		} else if ns > minutes {
			mins := ns / minutes
			ns -= mins * minutes
			secs := ns / seconds
			content = fmt.Sprintf("%dm, %ds", mins, secs)
		} else if ns > seconds {
			secs := ns / seconds
			ns -= secs * seconds
			millis := ns / milliseconds
			content = fmt.Sprintf("%ds, %dms", secs, millis)
		} else if ns > milliseconds {
			millis := ns / milliseconds
			ns -= millis * milliseconds
			micros := ns / microseconds
			content = fmt.Sprintf("%dms, %d\u00B5s", millis, micros)
		} else {
			content = fmt.Sprintf("%d\u00B5s", ns/microseconds)
		}
	} else {
		s := int64(d.Seconds())
		if s > hours {
			hrs := s / hours
			s -= hrs * hours
			mins := s / minutes
			content = fmt.Sprintf("%dh, %dm", hrs, mins)
		} else if s > minutes {
			mins := s / minutes
			s -= mins * minutes
			secs := s / seconds
			content = fmt.Sprintf("%dm, %ds", mins, secs)
		} else {
			content = fmt.Sprintf("%ds", s)
		}
	}

	p.appendSegment("duration", segment{
		content:    content,
		foreground: p.theme.DurationFg,
		background: p.theme.DurationBg,
	})
}
