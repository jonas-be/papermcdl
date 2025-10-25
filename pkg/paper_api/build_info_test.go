package paper_api

import (
	"github.com/jarcoal/httpmock"
	"reflect"
	"testing"
)

func TestPapermcAPI_GetBuildInfo(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://fill.papermc.io/v2/projects/velocity/versions/3.1.1/builds/445",
		httpmock.NewStringResponder(200, "{\"project_id\":\"paper\",\"project_name\":\"Paper\",\"version\":\"1.8.8\",\"build\":445,\"time\":\"2021-12-20T00:10:48.936Z\",\"channel\":\"default\",\"promoted\":false,\"changes\":[{\"commit\":\"2ce7ea620a6d9590a9f93b47ac45e240dac7988a\",\"summary\":\"Update\",\"message\":\"Update\"}],\"downloads\":{\"application\":{\"name\":\"paper-1.8.8-445.jar\",\"sha256\":\"7ff6d2cec671ef0d95b3723b5c92890118fb882d73b7f8fa0a2cd31d97c55f86\"}}}"))

	want := BuildInfo{
		ProjectID:   "paper",
		ProjectName: "Paper",
		Version:     "1.8.8",
		Build:       445,
		Time:        "2021-12-20T00:10:48.936Z",
		Channel:     "default",
		Promoted:    false,
		Changes: []struct {
			Commit  string `json:"commit"`
			Summary string `json:"summary"`
			Message string `json:"message"`
		}{
			{
				Commit:  "2ce7ea620a6d9590a9f93b47ac45e240dac7988a",
				Summary: "Update",
				Message: "Update",
			},
		},
		Downloads: struct {
			Application struct {
				Name   string `json:"name"`
				Sha256 string `json:"sha256"`
			} `json:"application"`
		}{
			Application: struct {
				Name   string `json:"name"`
				Sha256 string `json:"sha256"`
			}{
				Name:   "paper-1.8.8-445.jar",
				Sha256: "7ff6d2cec671ef0d95b3723b5c92890118fb882d73b7f8fa0a2cd31d97c55f86",
			},
		},
	}

	p := PapermcAPI{
		URL: "https://fill.papermc.io/v2",
	}
	builds, err := p.GetBuildInfo("velocity", "3.1.1", "445")
	if err != nil {
		t.Errorf("GetBuildInfo() unintentionally error = %v", err)
		return
	}
	if !reflect.DeepEqual(builds, want) {
		t.Errorf("GetBuildInfo() got = %v, want %v", builds, want)
	}
}
