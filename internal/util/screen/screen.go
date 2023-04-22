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

func FullWidthField(s tcell.Screen, text string, line int) {
	width, _ := s.Size()

	writerPos := 0
	textWriterPos := 0
	for i := 0; i < width; i++ {

		if i >= (width)/2-(len(text)/2) {
			if !(textWriterPos >= len(text)) {
				writerPos = InsertChar(s, rune(text[textWriterPos]), writerPos, line, tcell.Style{}.Background(tcell.ColorGreen).Foreground(tcell.ColorBlack))
				textWriterPos++
			} else {
				writerPos = InsertChar(s, ' ', writerPos, line, tcell.Style{}.Background(tcell.ColorGreen).Foreground(tcell.ColorBlack))
			}
		} else {
			writerPos = InsertChar(s, ' ', writerPos, line, tcell.Style{}.Background(tcell.ColorGreen).Foreground(tcell.ColorBlack))
		}
	}

}
