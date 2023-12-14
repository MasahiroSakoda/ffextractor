// Package ui provides ui interaction built with bubbletea
package ui

import (
	"os"
	"strconv"

	"github.com/MasahiroSakoda/ffextractor/internal/constants"
	"github.com/MasahiroSakoda/ffextractor/internal/segment"
	"github.com/MasahiroSakoda/ffextractor/internal/util"

	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/cockroachdb/errors"
	"github.com/sirupsen/logrus"
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
		errors.WithDetail(msg.err, "Error detected")
		if !m.detecting && !m.splitting && !m.concatenating {
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

	case tea.QuitMsg:
		m.removeTempFiles(m.tempDir)

	// spinner
	case spinner.TickMsg:
		switch {
		case m.detecting || m.splitting || m.concatenating:
			m.spinner, cmd = m.spinner.Update(msg)
			cmds = append(cmds, cmd)
		}

	// after silence detected
	case silenceDetectedMsg:
		m.detecting = false
		m.splitting = true
		m.makeTableRows()
		cmd := m.fetchSplittedSegments()
		cmds = append(cmds, cmd)

	// msg for split detected segments
	case splitProcessingMsg:
		// m.segments = m.removeSegment(msg.index)
		// m.makeTableRows()

	// msg for split process completed
	case splitCompletedMsg:
		m.splitting = false
		m.concatenating = true
		m.concatFile.Close()
		// cmd := m.executeConcatSegments(m.path)
		// cmds = append(cmds, cmd)

	// msg for concatenate process completed
	case concatCompletedMsg:
		m.concatenating = false
		// m.removeTempFiles(m.tempDir)
		logrus.Debugf("Concatenate completed: %s", msg.output)
		// return m, tea.Quit
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) removeTempFiles(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		logrus.Errorf("%s: %v", constants.ErrFileRemove, err)
	}
}

// makeTableRows makes table using detected segments
func (m *Model) makeTableRows() {
	var rows []table.Row
	for i, segment := range m.segments {
		var row = table.Row{
			strconv.Itoa(i + 1),
			util.GetFilenameFromPath(segment.Input),
			util.SecondsToString(segment.Start),
			util.SecondsToString(segment.End),
			util.SecondsToString(segment.Duration),
		}
		rows = append(rows, row)
	}
	m.table.SetRows(rows)
}

// removeSegment returns removed segments with index
func (m *Model) removeSegment(index int) []segment.Model {
	return m.segments[:index+copy(m.segments[index:], m.segments[index+1:])]
}
