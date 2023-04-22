package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"log"
	"papermc-downloader/internal/cli/list"
	"papermc-downloader/internal/cli/papermc"
	"papermc-downloader/pkg/paper_api"
)

func main() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorPurple)

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.EnablePaste()
	s.Clear()

	papermcAPI := paper_api.PapermcAPI{URL: "https://papermc.io"}
	papermcSelector := papermc.PapermcSelector{
		PapermcApi: papermcAPI,
		List: list.List{
			Screen:        s,
			List:          nil,
			DefaultStyle:  defStyle,
			SelectedStyle: boxStyle,
		},
	}
	papermcSelector.ShowProjects()

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
		if papermcSelector.View != "no-render" {
			papermcSelector.Render()
		}
	}
}
