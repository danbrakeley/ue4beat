package main

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Output Log Levels
const (
	LevelInfo    = "info"    // Unreal's Display and Verbose
	LevelWarning = "warning" // Unreal's Warning
	LevelError   = "error"   // Unreal's Error
)

// UE4Line holds a fully parsed server log line
type UE4Line struct {
	Timestamp *time.Time `json:"@timestamp,omitempty"`
	Message   string     `json:"message"`
	Frame     int        `json:"fields.frame"`
	Category  string     `json:"fields.category,omitempty"`
	Level     string     `json:"fields.level"`
}

// regular expressions
var (
	// ANSI escape sequences are used to change the text color for Warnings and Errors
	// From: https://stackoverflow.com/questions/14693701/how-can-i-remove-the-ansi-escape-sequences-from-a-string-in-python#14693789
	rANSIEsc = regexp.MustCompile(`\x1B\[[0-?]*[ -/]*[@-~]`)

	// The Unreal timestamp is in the form "[2018.09.21-21.44.44:949]"
	rTimestamp = regexp.MustCompile(`^\[([0-9.\-:]+)\]`)

	// The Unreal frame number is potentially padded with whitespace, ie "[  0]"
	rFrame = regexp.MustCompile(`^\[\s*([0-9]+)\]`)

	// The Unreal categories are alpha-numeric (no whitespace) and end with ": ".
	// To rule out "sh: ", which is a real server log entry, I also mandate a category must be at least 3 chars long.
	rCategory = regexp.MustCompile(`^([a-zA-Z0-9_]{3,}): `)

	// Unreal doesn't always print the level, but if one was included, we want to find and remove it.
	rVerbose = regexp.MustCompile(`^Verbose: `)
	rDisplay = regexp.MustCompile(`^Display: `)
	rWarning = regexp.MustCompile(`^Warning: `)
	rError   = regexp.MustCompile(`^Error: `)
)

// ParseLine parses a raw UE4 log line and makes a UE4Line object
func ParseLine(line string) UE4Line {
	var u UE4Line
	var matches []string

	// First, we may encounter ANSI escape sequences (text coloring).
	// We don't care about them, so remove them.
	line = rANSIEsc.ReplaceAllString(line, "")

	// Second, we may encounter a timestamp, which we should parse, then remove.
	matches = rTimestamp.FindStringSubmatch(line)
	if len(matches) == 2 {
		// To preserve fractional seconds (ms), Go needs the seconds to be ##.### (Unreal provides ##:###)
		raw := strings.Replace(matches[1], ":", ".", -1)
		// Now parse our YYYY.MM.DD-hh.mm.ss.sss
		if t, err := time.Parse("2006.01.02-15.04.05", raw); err == nil {
			u.Timestamp = &t
			line = rTimestamp.ReplaceAllString(line, "")
		}
	}

	// Third, we may encounter a frame number, which we should parse, then remove.
	u.Frame = -1
	matches = rFrame.FindStringSubmatch(line)
	if len(matches) == 2 {
		if frame, err := strconv.Atoi(matches[1]); err == nil {
			u.Frame = frame
			line = rFrame.ReplaceAllString(line, "")
		}
	}

	// Fourth, we may encounter a Log Category.
	matches = rCategory.FindStringSubmatch(line)
	if len(matches) == 2 {
		u.Category = matches[1]
		line = rCategory.ReplaceAllString(line, "")
	}

	u.Level = LevelInfo
	switch {
	case rDisplay.MatchString(line):
		line = rDisplay.ReplaceAllString(line, "")
	case rVerbose.MatchString(line):
		line = rVerbose.ReplaceAllString(line, "")
	case rWarning.MatchString(line):
		u.Level = LevelWarning
		line = rWarning.ReplaceAllString(line, "")
	case rError.MatchString(line):
		u.Level = LevelError
		line = rError.ReplaceAllString(line, "")
	}

	u.Message = strings.TrimSpace(line)
	return u
}
