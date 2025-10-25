package paper_api

import (
	"github.com/jarcoal/httpmock"
	"reflect"
	"testing"
)

func TestPapermcAPI_GetProjects(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://fill.papermc.io/v2/projects/",
		httpmock.NewStringResponder(200, "{\"projects\":[\"paper\",\"travertine\",\"waterfall\",\"velocity\",\"folia\"]}"))

	want := Projects{Projects: []string{
		"paper",
		"travertine",
		"waterfall",
		"velocity",
		"folia",
	}}

	p := PapermcAPI{
		URL: "https://fill.papermc.io/v2",
	}
	projects, err := p.GetProjects()
	if err != nil {
		t.Errorf("GetProjects() unintentionally error = %v", err)
		return
	}
	if !reflect.DeepEqual(projects, want) {
		t.Errorf("GetProjects() got = %v, want %v", projects, want)
	}
}
