package papermc

import (
	"fmt"
	"papermc-downloader/internal/cli/list"
	"papermc-downloader/internal/util"
	"papermc-downloader/pkg/paper_api"
	"strconv"
	"strings"
)

type PapermcSelector struct {
	PapermcApi   paper_api.PapermcAPI
	view         string
	List         list.List
	project      string
	versionGroup string
	versions     paper_api.Versions
	version      string
	build        string
}

func (p *PapermcSelector) SelectorUp() {
	p.List.SelectorUp()
}

func (p *PapermcSelector) SelectorDown() {
	p.List.SelectorDown()
}

func (p PapermcSelector) Render() {
	p.List.Render()
}

func (p *PapermcSelector) EnterInput() {
	switch p.view {
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
		break
	}
}

func (p *PapermcSelector) ShowProjects() {
	projects, err := p.PapermcApi.GetProjects()
	if err != nil {
		fmt.Println("fetch failed: ", err)
		return
	}
	p.List.List = projects.Projects
	p.view = "projects"
}

func (p *PapermcSelector) ShowVersionGroups(project string) {
	versions, err := p.PapermcApi.GetVersions(project)
	if err != nil {
		fmt.Println("fetch failed: ", err)
		return
	}
	p.List.List = util.ReverseStringArray(versions.VersionGroups)
	p.versions = versions
	p.view = "version-groups"
}

func (p *PapermcSelector) ShowVersions(versions paper_api.Versions, version string) {
	p.List.List = util.ReverseStringArray(p.versionGroupFilter(versions, version))
	p.view = "versions"
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
	p.view = "builds"
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
