package paper_api

import "encoding/json"

type Versions struct {
	ProjectID     string   `json:"project_id"`
	ProjectName   string   `json:"project_name"`
	VersionGroups []string `json:"version_groups"`
	Versions      []string `json:"versions"`
}

func (p PapermcAPI) GetVersions(project string) (Versions, error) {
	res, err := p.sendRequest("/api/v2/projects/" + project)
	if err != nil {
		return Versions{}, err
	}

	var response Versions
	err = json.Unmarshal(res, &response)
	if err != nil {
		return Versions{}, err
	}

	return response, nil
}
