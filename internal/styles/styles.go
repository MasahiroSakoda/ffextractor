// Package styles provides ui style for bubbletea
package styles

import "github.com/charmbracelet/lipgloss"

var (
	// ColorMain is primary color
	ColorMain = lipgloss.Color("#ff00ff")
	// ColorErr is to display error
	ColorErr  = lipgloss.Color("#ff0000")

	// StyleSpinner is style for spinner
	StyleSpinner            = lipgloss.NewStyle().Foreground(ColorMain)
	// StyleActive is style for active state
	StyleActive             = lipgloss.NewStyle().Bold(true)
	// StyleActionHeader is header style for action heaeder
	StyleActionHeader       = lipgloss.NewStyle().Bold(true).Padding(0, 1).Background(ColorMain)
	// StyleErrorHeader is header style for active state error
	StyleErrorHeader        = lipgloss.NewStyle().Bold(true).Padding(0, 1).Background(ColorErr)
	// StyleDone is style for ui state is done
	StyleDone               = lipgloss.NewStyle().Faint(true)
	// StyleTruncated is style for truncated state
	StyleTruncated          = lipgloss.NewStyle().Faint(true)
	// StyleNotificationBorder is style for notification
	StyleNotificationBorder = lipgloss.NewStyle().Foreground(ColorMain)
	// StyleNotificationText is style for notification
	StyleNotificationText   = lipgloss.NewStyle().Bold(true)
	// StyleLink is style for hyperlink
	StyleLink               = lipgloss.NewStyle().Underline(true)
)
