package gui

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"log"
	"papermcdl/internal/gui/header"
	"papermcdl/internal/gui/list"
	"papermcdl/internal/gui/papermc"
	"papermcdl/internal/util/screen"
	"papermcdl/pkg/paper_api"
	"time"
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
			screen.FullWidthField(s, "Downloading...", 4)
			s.Show()

			stopChan := make(chan bool)
			go downloadAnimation(s, stopChan)

			err := papermcSelector.Download()
			if err != nil {
				_ = s.Beep()
				fmt.Println("download failed: ", err)
				return
			}
			stopChan <- true
			quit()
			fmt.Println("Downloaded successfully!")
			return
		} else {
			papermcSelector.Render()
		}
		header.Title = papermcSelector.View
		header.Render(0)
	}
}

func downloadAnimation(s tcell.Screen, stop chan bool) {
	for {
		select {
		case <-stop:
			return
		default:
			milli := time.Now().UnixMilli() % 1000
			if milli < 333 {
				screen.FullWidthField(s, "Downloading.  ", 4)
			} else if milli < 666 {
				screen.FullWidthField(s, "Downloading.. ", 4)
			} else {
				screen.FullWidthField(s, "Downloading...", 4)
			}
			s.Show()
			time.Sleep(50 * time.Millisecond)
		}
	}
}
