package gui

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"log"
	"papermc-downloader/internal/gui/header"
	"papermc-downloader/internal/gui/list"
	"papermc-downloader/internal/gui/papermc"
	"papermc-downloader/internal/util/screen"
	"papermc-downloader/pkg/paper_api"
)

func StartGUI(papermcAPI paper_api.PapermcAPI) {
	defaultStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	selectedStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple)

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defaultStyle)
	s.EnablePaste()
	s.Clear()

	papermcSelector := papermc.PapermcSelector{
		PapermcApi: papermcAPI,
		Line:       2,
		List: list.List{
			Screen: s,
			//Line:          1,
			List:          nil,
			DefaultStyle:  defaultStyle,
			SelectedStyle: selectedStyle,
		},
	}
	papermcSelector.ShowProjects()
	header := header.Header{
		Screen: s,
		Title:  papermcSelector.View,
	}

	header.Render(0)

	quit := func() {
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	for {
		s.Show()

		ev := s.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyCtrlC {
				return
			} else if ev.Key() == tcell.KeyEscape {
				done := papermcSelector.GoBack()
				if !done {
					return
				}
			} else if ev.Key() == tcell.KeyDown {
				papermcSelector.SelectorDown()
			} else if ev.Key() == tcell.KeyUp {
				papermcSelector.SelectorUp()
			} else if ev.Key() == tcell.KeyEnter {
				err := papermcSelector.EnterInput()
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}
		if papermcSelector.View == "build-info" {
			screen.FullWidthField(s, "[ENTER] Download ", 8)
		} else if papermcSelector.View == "download" {
			s.Clear()
			//screen.FullWidthField(s, "Downloading ", 4)
			//go func(s tcell.Screen) {
			go func(s tcell.Screen) {
				err := papermcSelector.Download(s)
				if err != nil {
					fmt.Println("hmmmmm")
				}
			}(s)
			if err != nil {
				//_ = s.Beep()
				fmt.Println("download failed: ", err)
				return
			}

			//_ = s.Beep()
			//}(s)
		} else {
			papermcSelector.Render()
		}
		header.Title = papermcSelector.View
		header.Render(0)
	}
}
