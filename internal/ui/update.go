package ui

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

const (
	padding  = 1
	maxWidth = 80
)

// Update : bubletea required method
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		cmd    tea.Cmd
		cmds []tea.Cmd
	)

	m.table, cmd = m.table.Update(msg)
	cmds = append(cmds, cmd)

	switch msg := msg.(type) {
	case spinner.TickMsg:
		switch {
		case m.loading:
			m.spinner, cmd = m.spinner.Update(msg)
			cmds = append(cmds, cmd)
		}
	case silenceDetectedMsg:
		m.loading = false
		m.table.SetRows(msg.rows)

	case blackoutDetectedMsg:
		m.loading = false

	case tea.WindowSizeMsg:
		m.termWidth, m.termHeight = msg.Width, msg.Height
		return m, nil
	case tickMsg:
		tickEvery()
	// key
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		}
	default:
		return m, tea.Batch(cmds...)
	}
	return m, tea.Batch(cmds...)
}
