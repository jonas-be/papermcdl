package paper_api

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func (p PapermcAPI) GetFileName(buildInfo BuildInfo) string {
	return buildInfo.Downloads.Application.Name
}

func (p PapermcAPI) GetDownloadString(project string, version string, build string) (string, string, error) {
	buildInfo, err := p.GetBuildInfo(project, version, build)
	if err != nil {
		return "", "", err
	}
	return fmt.Sprintf("%v/projects/%v/versions/%v/builds/%v/downloads/%v",
		p.URL, project, version, build, p.GetFileName(buildInfo)), p.GetFileName(buildInfo), nil
}

func (p PapermcAPI) Download(project string, version string, build string) error {
	url, fileName, err := p.GetDownloadString(project, version, build)
	if err != nil {
		return err
	}

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

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return err
	}

	fmt.Println("File downloaded successfully!")
	return nil
}
