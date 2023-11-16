package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/bubbles/spinner"

	tea "github.com/charmbracelet/bubbletea"
)

type tickMsg time.Time

const (
	padding  = 2
	maxWidth = 80
)

// Update : bubletea required method
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// spinner
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.Spinner, cmd = m.Spinner.Update(msg)
		return m, cmd
	// progress
	case tea.WindowSizeMsg:
		m.Progress.Width = msg.Width - padding * 2 - 4
		if m.Progress.Width > maxWidth {
			m.Progress.Width = maxWidth
		}
		return m, nil
	// key
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		default:
			return m, nil
		}
	case tickMsg:
		if m.Progress.Percent() == 1.0 {
			return m, tea.Quit
		}
		cmd := m.Progress.IncrPercent(0.25)
		return m, tea.Batch(tickCmd(), cmd)
	case progress.FrameMsg:
		progressModel, cmd := m.Progress.Update(msg)
		m.Progress = progressModel.(progress.Model)
		return m, cmd
	default:
		return m, nil
	}
}

func tickCmd() tea.Cmd {
	return tea.Tick(time.Second * 1, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
