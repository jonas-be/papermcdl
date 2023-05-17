package paper_api

import (
	"encoding/json"
	"fmt"
	"papermcdl/internal/util"
	"papermcdl/pkg/latest"
)

type Versions struct {
	ProjectID     string   `json:"project_id"`
	ProjectName   string   `json:"project_name"`
	VersionGroups []string `json:"version_groups"`
	Versions      []string `json:"versions"`
}

func (v Versions) PrintVersions() {
	for _, version := range v.Versions {
		fmt.Println(version)
	}
}

func (v Versions) GetLatestVersion() (string, error) {
	_, item, err := latest.GetLatestItem(util.ReverseStringArray(v.Versions))
	if err != nil {
		return "", err
	}
	return item, nil
}

func (v Versions) GetLatestSnapshotVersion() (string, error) {
	_, item, err := latest.GetLatestItemSnapshot(util.ReverseStringArray(v.Versions))
	if err != nil {
		return "", err
	}
	return item, nil
}

func (v Versions) GetLatestVersionGroup() (string, error) {
	_, item, err := latest.GetLatestItem(util.ReverseStringArray(v.VersionGroups))
	if err != nil {
		return "", err
	}
	return item, nil
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
