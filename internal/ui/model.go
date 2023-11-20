package ui

import (
	"strconv"
	"time"

	"github.com/MasahiroSakoda/ffextractor/internal/constants"
	"github.com/MasahiroSakoda/ffextractor/internal/util"
	"github.com/MasahiroSakoda/ffextractor/internal/ffmpeg"
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
type silenceDetectedMsg struct {
	rows []table.Row
}

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
		m.fetchSilenceSegments,
		m.spinner.Tick,
	)
}

// tickEvery : update every time
func tickEvery() tea.Cmd {
	return tea.Tick(time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
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
			table.WithFocused(false),
			table.WithHeight(8),
		),
	}
}

func (m *Model) fetchSilenceSegments() tea.Msg {
	segments, err := ffmpeg.DetectSilence(m.path)
	if err != nil {
		return constants.ErrSilenceDetect
	}

	var rows []table.Row
	for i, segment := range segments {
		var row = table.Row{
			strconv.Itoa(i + 1),
			util.GetFilenameFromPath(segment.Filename),
			strconv.FormatFloat(segment.Start,    'f', -1, 64),
			strconv.FormatFloat(segment.End,      'f', -1, 64),
			strconv.FormatFloat(segment.Duration, 'f', -1, 64),
		}
		rows = append(rows, row)
	}
	return silenceDetectedMsg{rows: rows}
}
