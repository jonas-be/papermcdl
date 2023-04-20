package list

import (
	"github.com/gdamore/tcell/v2"
	"papermc-downloader/internal/util"
	"strconv"
)

type List struct {
	Screen        tcell.Screen
	List          []string
	selected      int
	DefaultStyle  tcell.Style
	SelectedStyle tcell.Style
}

func (l *List) GetSelected() string {
	return l.List[l.selected]
}

func (l *List) SelectorDown() {
	l.selected++
	if l.selected >= len(l.List) {
		l.selected = 0
	}
}

func (l *List) SelectorUp() {
	l.selected--
	if l.selected < 0 {
		l.selected = len(l.List) - 1
	}
}

func (l *List) Render() {
	l.Screen.Clear()
	idMaxLen := len(strconv.Itoa(len(l.List)))
	for i, listItem := range l.List {
		writerPos := 0
		style := l.DefaultStyle

		style, writerPos = l.insertSelectorColumn(i, style, writerPos)

		writerPos = l.insertSpacer(1, writerPos, i, style)

		writerPos = l.insertRowId(writerPos, idMaxLen, i, style)

		writerPos = l.insertSpacer(1, writerPos, i, style)
		writerPos = l.insertChar(1, '|', writerPos, i, style)
		writerPos = l.insertSpacer(4, writerPos, i, style)

		for n, char := range listItem {
			l.Screen.SetContent(n+writerPos, i, char, nil, style)
		}
	}
}

func (l *List) insertSelectorColumn(i int, style tcell.Style, writerPos int) (tcell.Style, int) {
	if l.selected == i {
		style = l.SelectedStyle
		writerPos = l.insertChar(1, '>', writerPos, i, style)
	} else {
		writerPos = l.insertSpacer(1, writerPos, i, style)
	}
	return style, writerPos
}

func (l *List) insertRowId(writerPos int, idMaxLen int, i int, style tcell.Style) int {
	writerPos += idMaxLen - 1
	var ni int
	var s rune
	for ni, s = range util.ReverseString(strconv.Itoa(i)) {
		l.Screen.SetContent(writerPos, i, s, nil, style)
		writerPos--
	}
	for n := ni; ni < idMaxLen-1; ni++ {
		l.Screen.SetContent(writerPos-n, i, ' ', nil, style)
		writerPos--
	}
	writerPos += idMaxLen + 1
	return writerPos
}

func (l *List) insertSpacer(count int, writerPos int, i int, style tcell.Style) int {
	for n := 0; n < count; n++ {
		l.Screen.SetContent(writerPos, i, ' ', nil, style)
		writerPos++
	}
	return writerPos
}

func (l *List) insertChar(count int, char rune, writerPos int, i int, style tcell.Style) int {
	for n := 0; n < count; n++ {
		l.Screen.SetContent(writerPos, i, char, nil, style)
		writerPos++
	}
	return writerPos
}
