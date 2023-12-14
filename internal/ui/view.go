// Package ui provides ui interaction built with bubbletea
package ui

import (
	"fmt"
	"strings"

)

// View : bubbletea required method
func (m *Model) View() string {
	pad := strings.Repeat(" ", padding * 5)
	var loadingMessage = ""
	// var helpMessage = ""

	switch {
	case m.detecting:
		loadingMessage = baseStyle(fmt.Sprintf("%s Detecting silence segments...", m.spinner.View()))
	case m.splitting:
		loadingMessage = baseStyle(fmt.Sprintf("%s Splitting silence segments...", m.spinner.View()))
	case m.concatenating:
		loadingMessage = baseStyle(fmt.Sprintf("%s Concatenating splitted segments...", m.spinner.View()))
	}
	return "\n" +
		loadingMessage + "\n" +
		m.table.View() + "\n" +
		"\n" +
		pad + helpStyle("Press \"esc\" or \"q\" key to quit") + "\n"
}
