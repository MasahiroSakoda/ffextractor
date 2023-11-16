package ui

import (
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/progress"
	"github.com/charmbracelet/lipgloss"
	"github.com/MasahiroSakoda/ffextractor/internal/styles"

	tea "github.com/charmbracelet/bubbletea"
)

// Model : TUI model as Elm architecture
type Model struct {
	// ui components (view models)
	Spinner  spinner.Model
	Progress progress.Model
	Percent  float64
}

var _ tea.Model = (*Model)(nil)

var helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#626262")).Render

// New : initialize model
func New() *Model {
	return &Model{
		Spinner:  spinner.New(spinner.WithSpinner(spinner.Dot), spinner.WithStyle(styles.StyleSpinner)),
		Progress: progress.New(progress.WithDefaultGradient()),
	}
}
