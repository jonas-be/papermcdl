package paper_api

import (
	"github.com/jarcoal/httpmock"
	"reflect"
	"testing"
)

func TestPapermcAPI_GetDownloadString(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://papermc.io/api/v2/projects/paper/versions/1.8.8/builds/445",
		httpmock.NewStringResponder(200, "{\"project_id\":\"paper\",\"project_name\":\"Paper\",\"version\":\"1.8.8\",\"build\":445,\"time\":\"2021-12-20T00:10:48.936Z\",\"channel\":\"default\",\"promoted\":false,\"changes\":[{\"commit\":\"2ce7ea620a6d9590a9f93b47ac45e240dac7988a\",\"summary\":\"Update\",\"message\":\"Update\"}],\"downloads\":{\"application\":{\"name\":\"paper-1.8.8-445.jar\",\"sha256\":\"7ff6d2cec671ef0d95b3723b5c92890118fb882d73b7f8fa0a2cd31d97c55f86\"}}}"))

	p := PapermcAPI{
		URL: "https://papermc.io",
	}
	downloadString, fileName, err := p.GetDownloadString("paper", "1.8.8", "445")
	if err != nil {
		t.Errorf("GetDownloadString() unintentionally error = %v", err)
		return
	}
	want := "https://papermc.io/api/v2/projects/paper/versions/1.8.8/builds/445/downloads/paper-1.8.8-445.jar"
	if !reflect.DeepEqual(downloadString, want) {
		t.Errorf("GetDownloadString() got = %v, want %v", downloadString, want)
	}
	want = "paper-1.8.8-445.jar"
	if !reflect.DeepEqual(fileName, want) {
		t.Errorf("GetDownloadString() got = %v, want %v", downloadString, want)
	}
}
