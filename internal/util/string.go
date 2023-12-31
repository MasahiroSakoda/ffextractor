// Package util provides function for utility
package util

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/lipgloss"
	"github.com/jedib0t/go-pretty/v6/text"
)

// TruncateString returns truncated string
func TruncateString(s string, l int) (string, bool) {
	rows := strings.Split(s, "\n")
	trunc := false
	if len(rows) > 1 {
		s = rows[0]
		trunc = true
	}
	if utf8.RuneCountInString(s) > l {
		s = string([]rune(s)[:l])
		trunc = true
	}

	return s, trunc
}

// PaddingRight returns string layout for bubbletea
func PaddingRight(s string, l int) string {
	return text.Pad(s, l, ' ')
}

// String returns string
func String(v string) *string {
	return &v
}

// Border returns string to use bubbletea
func Border(str string, style lipgloss.Style) string {
	// NOTE: Don't use `lipgloss.Border“.
	//	See https://github.com/charmbracelet/lipgloss/issues/40 .
	lines := strings.Split(str, "\n")
	width := text.LongestLineLen(str)

	b := strings.Repeat("─", width+2)
	bt := style.Render(fmt.Sprintf("┌%s┐", b))
	bb := style.Render(fmt.Sprintf("└%s┘", b))

	rslt := []string{bt}
	for _, line := range lines {
		b := style.Render("│")
		rslt = append(rslt, fmt.Sprintf("%s %s %s", b, PaddingRight(line, width), b))
	}
	rslt = append(rslt, bb)

	return strings.Join(rslt, "\n")
}

// GetStringInBetween returns empty string if no start or end string found
func GetStringInBetween(str string, start string, end string) (result string) {
	s := strings.Index(str, start)
	if s == -1 {
		return
	}
	s += len(start)
	e := strings.Index(str[s:], end)
	if e == -1 {
		return
	}
	return str[s : s+e]
}
