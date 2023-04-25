//package main
//
//import (
//	"fmt"
//	"github.com/gdamore/tcell/v2"
//	"io"
//	"log"
//	"net/http"
//	"os"
//	screen_util "papermc-downloader/internal/util/screen"
//	"papermc-downloader/pkg/paper_api"
//	"strconv"
//	"time"
//)
//
//func main() {
//	papermcAPI := paper_api.PapermcAPI{URL: "https://papermc.io"}
//	url, filename, _ := papermcAPI.GetDownloadString("paper", "1.19.4", "517")
//	//DownloadFile(url, ".")
//
//	screen, err := tcell.NewScreen()
//	if err != nil {
//		log.Fatalf("%+v", err)
//	}
//	if err := screen.Init(); err != nil {
//		log.Fatalf("%+v", err)
//	}
//	screen.SetStyle(tcell.Style{}.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite))
//	screen.EnablePaste()
//	screen.Clear()
//	screen.Show()
//
//	err = DownloadWithProgressBar(url, filename, screen, 4)
//	if err != nil {
//		return
//	}
//	//for {
//	//	print(percentage)
//	//}
//
//}
