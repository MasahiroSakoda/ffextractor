package ui

import (
	"time"

	"github.com/MasahiroSakoda/ffextractor/internal/segment"
	"github.com/MasahiroSakoda/ffextractor/internal/styles"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Model : TUI model as Elm architecture
type Model struct {
	// Parameter
	path    string

	// Status
	loading bool

	// ui components (view models)
	spinner spinner.Model
	table   table.Model

	termWidth  int
	termHeight int
}

type tickMsg time.Time
type errMsg struct { err error}
type silenceDetectedMsg struct {
	segments []segment.Model
}
type splitProcessingMsg struct {}

var (
	_ tea.Model = (*Model)(nil)
	baseStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).Render
	helpStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#626262")).Render
)

func columns() []table.Column {
	columns := []table.Column{
		{ Title: "ID",       Width: 5},
		{ Title: "File",     Width: 35},
		{ Title: "Start",    Width: 10},
		{ Title: "End",      Width: 10},
		{ Title: "Duration", Width: 10},
	}
	return columns
}

// Init : builder func for bubbletea
func (m *Model) Init() tea.Cmd {
	return tea.Batch(
		tickEvery(),
		m.spinner.Tick,
		m.fetchSilenceSegments,
	)
}

// New : initialize model
func New(path string) *Model {
	return &Model{
		loading: true,
		path:    path,
		spinner: spinner.New(
			spinner.WithSpinner(spinner.Dot),
			spinner.WithStyle(styles.StyleSpinner)),
		table: table.New(
			table.WithColumns(columns()),
			table.WithRows(make([]table.Row, 0)),
			table.WithFocused(true),
			table.WithHeight(10),
		),
	}
}
