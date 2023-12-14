// Package ui provides ui interaction built with bubbletea
package ui

import (
	"fmt"
	"os"
	"time"

	"github.com/MasahiroSakoda/ffextractor/internal/constants"
	"github.com/MasahiroSakoda/ffextractor/internal/extractor"

	tea "github.com/charmbracelet/bubbletea"
)

// Init : builder func for bubbletea
func (m *Model) Init() tea.Cmd {
	return tea.Batch(
		tickEvery(),
		m.spinner.Tick,
		m.fetchSilenceSegments(),
	)
}

// tickEvery : update every time
func tickEvery() tea.Cmd {
	return tea.Tick(time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func errorDetected(e error) tea.Msg {
	if e != nil { return errMsg{err: e} }
	return nil
}

// fetchSilenceSegments return command to detect silence segments
func (m *Model) fetchSilenceSegments() tea.Cmd {
	return func() tea.Msg {
		segments, err := extractor.DetectSilentSegments(m.path)
		if err != nil {
			return errMsg{err: constants.ErrSilenceDetect}
		}
		m.segments = segments
		return silenceDetectedMsg{}
	}
}

// fetchSplittedSegments returns command to split silence segments
func (m *Model) fetchSplittedSegments() tea.Cmd {
	return func() tea.Msg {

		tmpFile, err := os.CreateTemp("", "concat.txt")
		if err != nil { return errorDetected(err) }
		m.concatFile = tmpFile

		for i, segment := range m.segments {
			err = extractor.SplitDetectedSegment(segment, m.tempDir)
			if err != nil { errorDetected(err) }
			m.table.SetCursor(i)
			m.splitProcessing(i)
			if segment.Output != "" {
				_, err = m.concatFile.WriteString(fmt.Sprintf("file '%s'\n", segment.Output))
				if err != nil {
					return errMsg{err: constants.ErrSplitSegment}
				}
			}
		}
		return splitCompletedMsg{}
	}
}

func (m *Model) splitProcessing(index int) tea.Msg {
	return splitProcessingMsg{index: index}
}

// func (m *Model) executeConcatSegments(filePath string) tea.Cmd {
// 	return func() tea.Msg {
// 		output, err := extractor.ConcatDetectedSegments(m.concatFile.Name(), filePath)
// 		if err != nil {
// 			return errorDetected(err)
// 		}
// 		return concatCompletedMsg{output: output}
// 	}
// }
