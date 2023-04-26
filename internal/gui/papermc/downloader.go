package papermc

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"io"
	"log"
	"net/http"
	"os"
	"papermc-downloader/internal/util/screen"
	"papermc-downloader/pkg/paper_api"
	"strconv"
	"time"
)

type DataField struct {
	Name  string
	Value string
}

func ShowBuildInfo(s tcell.Screen, line int, api paper_api.PapermcAPI, project string, version string, build string) error {
	buildInfo, err := api.GetBuildInfo(project, version, build)
	if err != nil {
		return err
	}

	dataTable := []DataField{
		{Name: "project:", Value: buildInfo.ProjectName},
		{Name: "version:", Value: buildInfo.Version},
		{Name: "build:", Value: strconv.Itoa(buildInfo.Build)},
		{Name: "time:", Value: buildInfo.Time},
	}

	s.Clear()
	NewInfoTable(s, line, dataTable)
	return nil
}

func NewInfoTable(s tcell.Screen, line int, data []DataField) {
	var keyMaxLen int
	var valMaxLen int
	for _, dataField := range data {
		if keyMaxLen < len(dataField.Name) {
			keyMaxLen = len(dataField.Name)
		}
		if valMaxLen < len(dataField.Value) {
			valMaxLen = len(dataField.Value)
		}
	}

	for _, dataField := range data {
		var writerPos int
		writerPos = drawCell(s, keyMaxLen, dataField.Name, writerPos, line, tcell.Style{}.Bold(true))
		writerPos = drawCell(s, 3, "   ", writerPos, line, tcell.Style{})
		writerPos = drawCell(s, valMaxLen, dataField.Value, writerPos, line, tcell.Style{})
		line++
	}
}

func drawCell(s tcell.Screen, maxLen int, str string, writerPos int, line int, style tcell.Style) int {
	var tagWriterPos int
	var char rune
	for tagWriterPos, char = range fmt.Sprintf("%v", str) {
		writerPos = screen.InsertChar(s, char, writerPos, line, style)
	}
	tagWriterPos++
	for n := tagWriterPos; n < maxLen; n++ {
		writerPos = screen.InsertChar(s, ' ', writerPos, line, style)
	}
	return writerPos
}

func PrintDownloadPercent(done chan int64, path string, total int64, s tcell.Screen, line int) {
	var stop bool = false

	for {
		select {
		case <-done:
			stop = true
		default:

			file, err := os.Open(path)
			if err != nil {
				log.Fatal(err)
			}

			fi, err := file.Stat()
			if err != nil {
				log.Fatal(err)
			}

			size := fi.Size()

			if size == 0 {
				size = 1
			}

			var percent float64 = float64(size) / float64(total) * 100

			screen.FullWidthField(s, fmt.Sprintf("%.0f%", percent), line)

			writerPos := 0
			w, _ := s.Size()
			//fmt.Println(float64(w), percent)
			//fmt.Println(int(float64(w) * (percent / 100)))
			s.SetContent(10, 10, 'X', nil, tcell.Style{}.Background(tcell.ColorGreen).Foreground(tcell.ColorWhite))
			writerPos = screen.InsertChars(s, int(float64(w)*(percent/100)), '|', writerPos, line+1, tcell.Style{}.Background(tcell.ColorGreen))
		}

		if stop {
			break
		}

		time.Sleep(time.Millisecond * 100)
	}
}

func DownloadWithProgressBar(url string, fileName string, s tcell.Screen, line int) error {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}
	defer file.Close()

	size, err := strconv.Atoi(resp.Header.Get("Content-Length"))
	if err != nil {
		fmt.Println("Error reading header:", err)
		return err
	}
	done := make(chan int64)

	go PrintDownloadPercent(done, fileName, int64(size), s, line)

	defer resp.Body.Close()

	n, err := io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	done <- n
	s.Beep()

	return nil
}
