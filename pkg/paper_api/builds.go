package paper_api

import (
	"encoding/json"
	"fmt"
	"github.com/jonas-be/papermcdl/internal/util"
	"github.com/jonas-be/papermcdl/pkg/latest"
)

type Builds struct {
	ProjectID   string `json:"project_id"`
	ProjectName string `json:"project_name"`
	Version     string `json:"version"`
	Builds      []int  `json:"builds"`
}

func (b Builds) PrintBuilds() {
	for _, build := range b.Builds {
		fmt.Println(build)
	}
}

func (b Builds) GetLatestBuild() (string, error) {
	_, item, err := latest.GetLatestItem(util.ReverseStringArray(latest.ConvertIntArrayToStringArray(b.Builds)))
	if err != nil {
		return "", err
	}
	return item, nil
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
