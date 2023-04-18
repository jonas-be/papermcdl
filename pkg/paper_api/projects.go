package paper_api

import (
	"encoding/json"
)

type Projects struct {
	Projects []string `json:"projects"`
}

func (p PapermcAPI) GetProjects() (Projects, error) {
	res, err := p.sendRequest("/api/v2/projects/")
	if err != nil {
		return Projects{}, err
	}

	var response Projects
	err = json.Unmarshal(res, &response)
	if err != nil {
		return Projects{}, err
	}

	return response, nil
}
