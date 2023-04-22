package papermc

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"papermc-downloader/internal/util/screen"
	"papermc-downloader/pkg/paper_api"
	"strconv"
)

type DataField struct {
	Name  string
	Value string
}

func ShowBuildInfo(s tcell.Screen, api paper_api.PapermcAPI, project string, version string, build string) error {
	buildInfo, err := api.GetBuildInfo(project, version, build)
	if err != nil {
		return err
	}

	dataTable := []DataField{
		{Name: "project:", Value: buildInfo.ProjectName},
		{Name: "version:", Value: buildInfo.Version},
		{Name: "build:", Value: strconv.Itoa(buildInfo.Build)},
		{Name: "time:", Value: buildInfo.Time},
	}

	s.Clear()
	NewInfoTable(s, dataTable)
	return nil
}

func NewInfoTable(s tcell.Screen, data []DataField) {
	var keyMaxLen int
	var valMaxLen int
	for _, dataField := range data {
		if keyMaxLen < len(dataField.Name) {
			keyMaxLen = len(dataField.Name)
		}
		if valMaxLen < len(dataField.Value) {
			valMaxLen = len(dataField.Value)
		}
	}

	var line int

	for _, dataField := range data {
		var writerPos int
		writerPos = drawCell(s, keyMaxLen, dataField.Name, writerPos, line, tcell.Style{}.Bold(true))
		writerPos = drawCell(s, 3, "   ", writerPos, line, tcell.Style{})
		writerPos = drawCell(s, valMaxLen, dataField.Value, writerPos, line, tcell.Style{})
		line++
	}
}

func drawCell(s tcell.Screen, maxLen int, str string, writerPos int, line int, style tcell.Style) int {
	var tagWriterPos int
	var char rune
	for tagWriterPos, char = range fmt.Sprintf("%v", str) {
		writerPos = screen.InsertChar(s, char, writerPos, line, style)
	}
	tagWriterPos++
	for n := tagWriterPos; n < maxLen; n++ {
		writerPos = screen.InsertChar(s, ' ', writerPos, line, style)
	}
	return writerPos
}
