package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Init : builder func for bubbletea
func (m *Model) Init() tea.Cmd {
	return tickCmd()
}
