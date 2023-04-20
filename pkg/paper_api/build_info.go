package paper_api

import (
	"encoding/json"
	"fmt"
)

type BuildInfo struct {
	ProjectID   string `json:"project_id"`
	ProjectName string `json:"project_name"`
	Version     string `json:"version"`
	Build       int    `json:"build"`
	Time        string `json:"time"`
	Channel     string `json:"channel"`
	Promoted    bool   `json:"promoted"`
	Changes     []struct {
		Commit  string `json:"commit"`
		Summary string `json:"summary"`
		Message string `json:"message"`
	} `json:"changes"`
	Downloads struct {
		Application struct {
			Name   string `json:"name"`
			Sha256 string `json:"sha256"`
		} `json:"application"`
	} `json:"downloads"`
}

func (p PapermcAPI) GetBuildInfo(project string, version string, build string) (BuildInfo, error) {
	res, err := p.sendRequest(fmt.Sprintf("/api/v2/projects/%v/versions/%v/builds/%v", project, version, build))
	if err != nil {
		return BuildInfo{}, err
	}

	var response BuildInfo
	err = json.Unmarshal(res, &response)
	if err != nil {
		return BuildInfo{}, err
	}

	return response, nil
}
