package screen

import "github.com/gdamore/tcell/v2"

func InsertChar(s tcell.Screen, char rune, writerPos int, line int, style tcell.Style) int {
	s.SetContent(writerPos, line, char, nil, style)
	writerPos++
	return writerPos
}

func InsertChars(s tcell.Screen, count int, char rune, writerPos int, line int, style tcell.Style) int {
	for i := 0; i < count; i++ {
		s.SetContent(writerPos, line, char, nil, style)
		writerPos++
	}
	return writerPos
}
