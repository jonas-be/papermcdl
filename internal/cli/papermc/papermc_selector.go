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
	View         string
	List         list.List
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
	p.List.Render()
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
		p.View = "no-render"
		err := ShowBuildInfo(p.List.Screen, p.PapermcApi, p.project, p.version, p.build)
		if err != nil {
			return err
		}
		break
	}
	p.List.Selected = 0
	return nil
}

func (p *PapermcSelector) GoBack() bool {
	p.List.Selected = 0
	switch p.View {
	case "version-groups":
		p.ShowProjects()
		return false
	case "versions":
		p.ShowVersionGroups(p.project)
		return false
	case "builds":
		p.ShowVersions(p.versions, p.version)
		return false
	case "no-render":
		p.ShowBuilds(p.project, p.version)
		p.View = "builds"
		return false
	}
	return true
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
	p.List.Tags = []list.Tag{{Name: "latest", ID: 0}}
	p.List.List = util.ReverseStringArray(versions.VersionGroups)
	p.versions = versions
	p.View = "version-groups"
}

func (p *PapermcSelector) ShowVersions(versions paper_api.Versions, version string) {
	p.List.Tags = []list.Tag{{Name: "latest", ID: 0}}
	p.List.List = util.ReverseStringArray(p.versionGroupFilter(versions, version))
	p.View = "versions"
}

func (p *PapermcSelector) ShowBuilds(project string, version string) {
	p.List.Tags = []list.Tag{{Name: "latest", ID: 0}}
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
