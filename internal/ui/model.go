package ui

import (
	"github.com/MasahiroSakoda/ffextractor/internal/styles"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

// Model : TUI model as Elm architecture
type Model struct {
	// ui components (view models)
	Spinner  spinner.Model
}

// Init : builder func for bubbletea
func (m *Model) Init() tea.Cmd {
	// return m.Spinner.Tick
	return tea.Batch(
		m.Spinner.Tick,
	)
}

var _ tea.Model = (*Model)(nil)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

// New : initialize model
func New() *Model {
	return &Model{
		Spinner:  spinner.New(spinner.WithSpinner(spinner.Dot), spinner.WithStyle(styles.StyleSpinner)),
	}
}
