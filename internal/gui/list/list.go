package list

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"papermcdl/internal/util"
	"papermcdl/internal/util/screen"
	"strconv"
)

type List struct {
	Screen        tcell.Screen
	Line          int
	List          []string
	Tags          []Tag
	Selected      int
	DefaultStyle  tcell.Style
	SelectedStyle tcell.Style
}

func (l *List) GetSelected() string {
	return l.List[l.Selected]
}

func (l *List) SelectorDown() {
	l.Selected++
	if l.Selected >= len(l.List) {
		l.Selected = 0
	}
}

func (l *List) SelectorUp() {
	l.Selected--
	if l.Selected < 0 {
		l.Selected = len(l.List) - 1
	}
}

func (l *List) Render() {
	l.Screen.Clear()

	var maxTagLen int
	maxTagLen = l.getMaxTagLen(maxTagLen)

	idMaxLen := len(strconv.Itoa(len(l.List)))
	for i, listItem := range l.List {
		line := i + l.Line
		writerPos := 0
		style := l.DefaultStyle

		style, writerPos = l.insertSelectorColumn(line, style, writerPos)

		writerPos = screen.InsertChars(l.Screen, 1, ' ', writerPos, line, style)

		writerPos = l.insertRowId(writerPos, idMaxLen, i, style)

		writerPos = screen.InsertChars(l.Screen, 1, ' ', writerPos, line, style)
		writerPos = screen.InsertChars(l.Screen, 1, '|', writerPos, line, style)

		writerPos = screen.InsertChars(l.Screen, 1, ' ', writerPos, line, style)
		writerPos = l.drawTagColumn(maxTagLen, writerPos, i, listItem, style)
		writerPos = screen.InsertChars(l.Screen, 1, ' ', writerPos, line, style)

		for n, char := range listItem {
			l.Screen.SetContent(n+writerPos, line, char, nil, style)
		}
	}
}

func (l *List) insertSelectorColumn(i int, style tcell.Style, writerPos int) (tcell.Style, int) {
	if l.Selected == i-l.Line {
		style = l.SelectedStyle
		writerPos = screen.InsertChars(l.Screen, 1, '>', writerPos, i, style)
	} else {
		writerPos = screen.InsertChars(l.Screen, 1, ' ', writerPos, i, style)
	}
	return style, writerPos
}

func (l *List) insertRowId(writerPos int, idMaxLen int, i int, style tcell.Style) int {
	writerPos += idMaxLen - 1
	var ni int
	var s rune
	for ni, s = range util.ReverseString(strconv.Itoa(i)) {
		l.Screen.SetContent(writerPos, i+l.Line, s, nil, style)
		writerPos--
	}
	for n := ni; ni < idMaxLen-1; ni++ {
		l.Screen.SetContent(writerPos-n, i+l.Line, ' ', nil, style)
		writerPos--
	}
	writerPos += idMaxLen + 1
	return writerPos
}

func (l *List) getMaxTagLen(maxTagLen int) int {
	for _, tag := range l.Tags {
		tagLen := len(tag.Label) + 2
		if tagLen > maxTagLen {
			maxTagLen = tagLen
		}
	}
	return maxTagLen
}

func (l *List) drawTagColumn(maxTagLen int, writerPos int, i int, listItem string, style tcell.Style) int {
	var tagToDraw Tag
	for _, tag := range l.Tags {
		if listItem == tag.ItemName {
			tagToDraw = tag
		}
	}
	var tagWriterPos int
	var char rune
	if tagToDraw.Label != "" {
		for tagWriterPos, char = range fmt.Sprintf("(%v)", tagToDraw.Label) {
			writerPos = screen.InsertChars(l.Screen, 1, char, writerPos, i+l.Line, style)
		}
		tagWriterPos++
	}
	for n := tagWriterPos; n < maxTagLen; n++ {
		writerPos = screen.InsertChars(l.Screen, 1, ' ', writerPos, i+l.Line, style)
	}
	return writerPos
}
