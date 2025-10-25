package paper_api

import (
	"encoding/json"
	"fmt"
)

type Projects struct {
	Projects []string `json:"projects"`
}

func (p Projects) PrintProjects() {
	for _, project := range p.Projects {
		fmt.Println(project)
	}
}

func (p PapermcAPI) GetProjects() (Projects, error) {
	res, err := p.sendRequest("/projects/")
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
