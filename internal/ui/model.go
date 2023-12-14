// Package ui provides ui interaction built with bubbletea
package ui

import (
	"os"
	"time"

	"github.com/MasahiroSakoda/ffextractor/internal/constants"
	"github.com/MasahiroSakoda/ffextractor/internal/segment"
	"github.com/MasahiroSakoda/ffextractor/internal/styles"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/lipgloss"
)

// Model : TUI model as Elm architecture
type Model struct {
	// Parameter
	path    string

	// Status
	detecting     bool
	splitting     bool
	concatenating bool
	err error

	// ui components (view models)
	spinner spinner.Model
	table   table.Model

	tempDir    string
	concatFile *os.File
	segments []segment.Model

	termWidth  int
	termHeight int
}

type tickMsg time.Time
type errMsg struct { err error }
type silenceDetectedMsg struct {}
type splitProcessingMsg struct { index int }
type splitCompletedMsg  struct {}
type concatCompletedMsg struct { output string }

var (
	_ tea.Model = (*Model)(nil)
	baseStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.DoubleBorder()).
		BorderForeground(lipgloss.Color("240")).Render
	helpStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#626262")).Render
)

func columns() []table.Column {
	columns := []table.Column{
		{ Title: "ID",       Width: 5},
		{ Title: "File",     Width: 35},
		{ Title: "Start",    Width: 15},
		{ Title: "End",      Width: 15},
		{ Title: "Duration", Width: 15},
	}
	return columns
}

// New : initialize model
func New(path string) *Model {
	// create temp directory
	dirPrefix := constants.CommandName + "_"
	tempDir, err := os.MkdirTemp("", dirPrefix)
	errorDetected(err)

	return &Model{
		detecting: true,
		splitting: false,
		concatenating: false,
		path:    path,
		tempDir: tempDir,
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
