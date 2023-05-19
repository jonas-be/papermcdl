package header

import (
	"github.com/gdamore/tcell/v2"
	"github.com/jonas-be/papermcdl/internal/util/screen"
)

type Header struct {
	Screen tcell.Screen
	Title  string
}

const backButtonText = " [ESC] Back "

var backButtonStyle = tcell.Style{}.Background(tcell.ColorOrange).Foreground(tcell.ColorBlack)

const quitButtonText = " [CRTL-C] Quit "

var quitButtonStyle = tcell.Style{}.Background(tcell.ColorRed).Foreground(tcell.ColorBlack)

func (h Header) Render(line int) {
	width, _ := h.Screen.Size()

	writerPos := 0
	backButtonWriterPos := 0
	quitButtonWriterPos := 0
	titleWriterPos := 0
	for i := 0; i < width; i++ {
		if i < len(backButtonText) {
			writerPos = screen.InsertChar(h.Screen, rune(backButtonText[backButtonWriterPos]), writerPos, line, backButtonStyle)
			backButtonWriterPos++
			continue
		}

		if width-i <= len(quitButtonText) {
			writerPos = screen.InsertChar(h.Screen, rune(quitButtonText[quitButtonWriterPos]), writerPos, line, quitButtonStyle)
			quitButtonWriterPos++
			continue
		}
		if i >= len(backButtonText)+(width-(len(backButtonText)+len(quitButtonText)))/2-(len(h.Title)/2) {
			if !(titleWriterPos >= len(h.Title)) {
				writerPos = screen.InsertChar(h.Screen, rune(h.Title[titleWriterPos]), writerPos, line, tcell.Style{}.Background(tcell.ColorGray).Foreground(tcell.ColorBlack))
				titleWriterPos++
			} else {
				writerPos = screen.InsertChar(h.Screen, ' ', writerPos, line, tcell.Style{}.Background(tcell.ColorGray).Foreground(tcell.ColorBlack))
			}
		} else {
			writerPos = screen.InsertChar(h.Screen, ' ', writerPos, line, tcell.Style{}.Background(tcell.ColorGray).Foreground(tcell.ColorBlack))
		}
	}
}
