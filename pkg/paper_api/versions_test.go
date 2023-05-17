package paper_api

import (
	"reflect"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestPapermcAPI_GetVersions(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://papermc.io/api/v2/projects/velocity",
		httpmock.NewStringResponder(200, "{\"project_id\":\"velocity\",\"project_name\":\"Velocity\",\"version_groups\":[\"1.0.0\",\"1.1.0\",\"3.0.0\"],\"versions\":[\"1.0.10\",\"1.1.9\",\"3.1.0\",\"3.1.1\",\"3.1.1-SNAPSHOT\",\"3.1.2-SNAPSHOT\",\"3.2.0-SNAPSHOT\"]}"))

	want := Versions{
		ProjectID:     "velocity",
		ProjectName:   "Velocity",
		VersionGroups: []string{"1.0.0", "1.1.0", "3.0.0"},
		Versions:      []string{"1.0.10", "1.1.9", "3.1.0", "3.1.1", "3.1.1-SNAPSHOT", "3.1.2-SNAPSHOT", "3.2.0-SNAPSHOT"},
	}

	p := PapermcAPI{
		URL: "https://papermc.io",
	}
	versions, err := p.GetVersions("velocity")
	if err != nil {
		t.Errorf("GetVersions() unintentionally error = %v", err)
		return
	}
	if !reflect.DeepEqual(versions, want) {
		t.Errorf("GetVersions() got = %v, want %v", versions, want)
	}
}
