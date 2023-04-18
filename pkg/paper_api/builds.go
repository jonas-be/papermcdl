package paper_api

import (
	"encoding/json"
	"fmt"
)

type Builds struct {
	ProjectID   string `json:"project_id"`
	ProjectName string `json:"project_name"`
	Version     string `json:"version"`
	Builds      []int  `json:"builds"`
}

func (p PapermcAPI) GetBuilds(project string, version string) (Builds, error) {
	res, err := p.sendRequest(fmt.Sprintf("/api/v2/projects/%v/versions/%v", project, version))
	if err != nil {
		return Builds{}, err
	}

	var response Builds
	err = json.Unmarshal(res, &response)
	if err != nil {
		return Builds{}, err
	}

	return response, nil
}
