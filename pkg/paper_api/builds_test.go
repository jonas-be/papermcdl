package paper_api

import (
	"github.com/jarcoal/httpmock"
	"reflect"
	"testing"
)

func TestPapermcAPI_GetBuilds(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://fill.papermc.io/v2/projects/velocity/versions/3.1.1",
		httpmock.NewStringResponder(200, "{\"project_id\":\"velocity\",\"project_name\":\"Velocity\",\"version\":\"3.1.1\",\"builds\":[98,99,102]}"))

	want := Builds{
		ProjectID:   "velocity",
		ProjectName: "Velocity",
		Version:     "3.1.1",
		Builds:      []int{98, 99, 102},
	}

	p := PapermcAPI{
		URL: "https://fill.papermc.io/v2",
	}
	builds, err := p.GetBuilds("velocity", "3.1.1")
	if err != nil {
		t.Errorf("GetBuilds() unintentionally error = %v", err)
		return
	}
	if !reflect.DeepEqual(builds, want) {
		t.Errorf("GetBuilds() got = %v, want %v", builds, want)
	}
}
