package ui

import (
	"strings"
)

// View : bubbletea required method
func (m *Model) View() string {
	pad := strings.Repeat(" ", padding)
	return "\n" +
		pad + m.Spinner.View() +
		pad + helpStyle("Press any key to quit")
}
