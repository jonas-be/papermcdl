package papermc

import (
	"fmt"
	"github.com/jonas-be/papermcdl/internal/gui/list"
	"github.com/jonas-be/papermcdl/internal/util"
	"github.com/jonas-be/papermcdl/pkg/latest"
	"github.com/jonas-be/papermcdl/pkg/paper_api"
	"strconv"
	"strings"
)

type PapermcSelector struct {
	PapermcApi   paper_api.PapermcAPI
	View         string
	List         list.List
	Line         int
	project      string
	versionGroup string
	versions     paper_api.Versions
	version      string
	build        string
}

func (p PapermcSelector) GetSelections() (string, string, string) {
	return p.project, p.version, p.build
}

func (p *PapermcSelector) SelectorUp() {
	p.List.SelectorUp()
}

func (p *PapermcSelector) SelectorDown() {
	p.List.SelectorDown()
}

func (p PapermcSelector) Render() {
	p.List.Line = p.Line
	p.List.Render()
}

func (p PapermcSelector) Download() error {
	err := p.PapermcApi.Download(p.project, p.version, p.build)
	if err != nil {
		return err
	}
	return nil
}

func (p *PapermcSelector) EnterInput() error {
	switch p.View {
	case "projects":
		p.project = p.List.GetSelected()
		p.ShowVersionGroups(p.project)
		break
	case "version-groups":
		p.versionGroup = p.List.GetSelected()
		p.ShowVersions(p.versions, p.versionGroup)
		break
	case "versions":
		p.version = p.List.GetSelected()
		p.ShowBuilds(p.project, p.version)
		break
	case "builds":
		p.build = p.List.GetSelected()
		p.View = "build-info"
		err := ShowBuildInfo(p.List.Screen, p.Line, p.PapermcApi, p.project, p.version, p.build)
		if err != nil {
			return err
		}
		break
	case "build-info":
		p.View = "download"
		break
	}
	return nil
}

func (p *PapermcSelector) GoBack() bool {
	switch p.View {
	case "version-groups":
		p.ShowProjects()
		return true
	case "versions":
		p.ShowVersionGroups(p.project)
		return true
	case "builds":
		p.ShowVersions(p.versions, p.version)
		return true
	case "build-info":
		p.ShowBuilds(p.project, p.version)
		p.View = "builds"
		return true
	}
	return false
}

func (p *PapermcSelector) ShowProjects() {
	projects, err := p.PapermcApi.GetProjects()
	if err != nil {
		fmt.Println("fetch failed: ", err)
		return
	}
	p.List.List = projects.Projects
	p.View = "projects"
}

func (p *PapermcSelector) ShowVersionGroups(project string) {
	versions, err := p.PapermcApi.GetVersions(project)
	if err != nil {
		fmt.Println("fetch failed: ", err)
		return
	}
	p.List.List = util.ReverseStringArray(versions.VersionGroups)
	p.versions = versions
	p.selectAndTagLatest()
	p.View = "version-groups"
}

func (p *PapermcSelector) ShowVersions(versions paper_api.Versions, version string) {
	p.List.List = util.ReverseStringArray(p.versionGroupFilter(versions, version))
	p.selectAndTagLatest()
	p.View = "versions"
}

func (p *PapermcSelector) ShowBuilds(project string, version string) {
	builds, err := p.PapermcApi.GetBuilds(project, version)
	if err != nil {
		fmt.Println("fetch failed: ", err)
		return
	}

	var stringSlice []string

	for _, intValue := range builds.Builds {
		stringSlice = append(stringSlice, strconv.Itoa(intValue))
	}

	p.List.List = util.ReverseStringArray(stringSlice)
	p.selectAndTagLatest()
	p.View = "builds"
}

func (p PapermcSelector) versionGroupFilter(versions paper_api.Versions, version string) []string {
	var versionList []string
	var versionListTmp []string

	filter := ""
	for i, _ := range version {
		filter = filter + string([]rune(version)[i])
		versionListTmp = []string{}
		for _, str := range versions.Versions {
			if strings.HasPrefix(str, filter) {
				versionListTmp = append(versionListTmp, str)
			}
		}
		if len(versionListTmp) == 0 {
			return versionList
		}
		versionList = versionListTmp
	}
	return versionList
}

func (p *PapermcSelector) selectAndTagLatest() {
	id, item, err := latest.GetLatestItem(p.List.List)
	if err == nil {
		p.List.Tags = []list.Tag{{Label: "latest/stable", ItemName: item}}
		p.List.Selected = id
	}
}
