package ui

import (
	"fmt"
	"strings"
)

// View : bubbletea required method
func (m *Model) View() string {
	pad := strings.Repeat(" ", padding * 5)
	var loadingMessage = ""
	if m.loading {
		loadingMessage = baseStyle(fmt.Sprintf("%s Detecting silence segments...", m.spinner.View()))
	}
	return "\n" +
		loadingMessage + "\n" +
		m.table.View() + "\n" +
		pad + helpStyle("Press any key to quit") + "\n"
}
