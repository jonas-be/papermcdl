package list

import (
	"github.com/gdamore/tcell/v2"
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
	for i, listItem := range l.List {
		style := l.DefaultStyle
		if l.selected == i {
			style = l.SelectedStyle
			l.Screen.SetContent(0, i, '>', nil, style)
			l.Screen.SetContent(3, i, 'X', nil, style)

		} else {
			l.Screen.SetContent(0, i, ' ', nil, style)
			l.Screen.SetContent(3, i, rune(i), nil, style)

		}
		l.Screen.SetContent(2, i, '(', nil, style)
		l.Screen.SetContent(4, i, ')', nil, style)
		l.Screen.SetContent(5, i, ' ', nil, style)

		for n, char := range listItem {
			l.Screen.SetContent(n+6, i, char, nil, style)
		}
	}
}
