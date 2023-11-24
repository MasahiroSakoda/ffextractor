package ui

import (
	"os"
	"fmt"
	"strconv"
	"time"

	"github.com/MasahiroSakoda/ffextractor/internal/constants"
	"github.com/MasahiroSakoda/ffextractor/internal/extractor"
	"github.com/MasahiroSakoda/ffextractor/internal/segment"
	"github.com/MasahiroSakoda/ffextractor/internal/util"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
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

	// update every time
	case tickMsg:
		tickEvery()

	case errMsg:
		if msg.err != nil {
			fmt.Println(msg.err)
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		m.termWidth, m.termHeight = msg.Width, msg.Height
		return m, nil

	// key
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit
		}

	// spinner
	case spinner.TickMsg:
		switch {
		case m.loading:
			m.spinner, cmd = m.spinner.Update(msg)
			cmds = append(cmds, cmd)
		}

	// after silence detected
	case silenceDetectedMsg:
		m.loading = false
		m.makeTableRows(msg.segments)

		// create temp directory
		dirPrefix := constants.CommandName + "_"
		tempDir, err := os.MkdirTemp("", dirPrefix)
		errorDetected(err)
		defer func() {
			os.RemoveAll(tempDir)
		}()

		for i, segment := range msg.segments {
			m.table.SetCursor(i)
			err := extractor.SplitDetectedSegment(segment, tempDir)
			errorDetected(err)
		}
		m.table.Blur()
		return m, tea.Quit

	// msg for split detected segments
	case splitProcessingMsg:
	}

	return m, tea.Batch(cmds...)
}

// tickEvery : update every time
func tickEvery() tea.Cmd {
	return tea.Tick(time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func errorDetected(e error) tea.Msg {
	if e != nil {
		return errMsg{err: e}
	}
	return nil
}

// makeTableRows makes table using detected segments
func (m *Model) makeTableRows(segments []segment.Model) {
	var rows []table.Row
	for i, segment := range segments {
		var row = table.Row{
			strconv.Itoa(i + 1),
			util.GetFilenameFromPath(segment.Input),
			strconv.FormatFloat(segment.Start,    'f', -1, 64),
			strconv.FormatFloat(segment.End,      'f', -1, 64),
			strconv.FormatFloat(segment.Duration, 'f', -1, 64),
		}
		rows = append(rows, row)
	}
	m.table.SetRows(rows)
}

// fetchSilenceSegments return message to detect silence segments
func (m *Model) fetchSilenceSegments() tea.Msg {
	segments, err := extractor.DetectSilentSegments(m.path)
	if err != nil {
		return errMsg{err: constants.ErrSilenceDetect}
	}
	return silenceDetectedMsg{segments: segments}
}
