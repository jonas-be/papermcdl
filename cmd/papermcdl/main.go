package main

import (
	"flag"
	"fmt"
	"papermcdl/internal/gui"
	"papermcdl/internal/util"
	"papermcdl/pkg/latest"
	"papermcdl/pkg/paper_api"
	"strings"
)

func main() {
	projectFlag := flag.String("p", "", "Project you want to download.")
	versionFlag := flag.String("v", "l", "(Optional) Version or Version Group of the Project. Defaults to latest version.")
	buildFlag := flag.String("b", "l", "(Optional) Build Number of the Project. You can use\"l\" to get the latest version, or \"ls\" to get the latest snapshot. Defaults to latest version.")
	infoFlag := flag.Bool("i", false, "(Optional) Show info or list available only.")
	flag.Parse()

	papermcAPI := paper_api.PapermcAPI{URL: "https://papermc.io"}

	// gui.StartGUI(papermcAPI) if no flags are provided
	if *projectFlag == "" && *versionFlag == "l" && *buildFlag == "l" && !*infoFlag {
		gui.StartGUI(papermcAPI)
		return
	}

	projects, err := papermcAPI.GetProjects()
	if err != nil {
		util.LogCon("error getting projects: ", err, util.Error)
		return
	}
	if !checkProjectFlag(projects, projectFlag) {
		return
	}

	versions, err := papermcAPI.GetVersions(*projectFlag)
	if err != nil {
		util.LogCon("error getting projects: ", err, util.Error)
		return
	}
	if !checkVersionFlag(versions, versionFlag) {
		return
	}

	builds, err := papermcAPI.GetBuilds(*projectFlag, *versionFlag)
	if err != nil {
		util.LogCon("error getting projects: ", err, util.Error)
		return
	}
	if !checkBuildFlag(builds, buildFlag) {
		return
	}

	downloadString, filename, err := papermcAPI.GetDownloadString(*projectFlag, *versionFlag, *buildFlag)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(downloadString)
	fmt.Println(filename)
	if *infoFlag {
		buildInfo, err := papermcAPI.GetBuildInfo(*projectFlag, *versionFlag, *buildFlag)
		if err != nil {
			util.LogCon("can not get buildInfo: ", err, util.Error)
			return
		}
		buildInfo.PrintBuildInfo()
		return
	}

	err = papermcAPI.Download(*projectFlag, *versionFlag, *buildFlag)
	if err != nil {
		util.LogCon("can not download: ", err, util.Error)
	}
}

func checkProjectFlag(projects paper_api.Projects, projectFlag *string) bool {
	if stringArrayContains(projects.Projects, *projectFlag) {
		return true
	}
	util.LogCon("Please provide a valid project:", "", util.Warning)
	projects.PrintProjects()
	return false
}

func checkVersionFlag(versions paper_api.Versions, versionFlag *string) bool {
	if stringArrayContains(versions.Versions, *versionFlag) {
		return true
	} else if strings.ToLower(*versionFlag) == "l" {
		var err error
		*versionFlag, err = versions.GetLatestVersion()
		if err != nil {
			util.LogCon("No latest version: ", err, util.Error)
			return false
		}
		return true
	} else if strings.ToLower(*versionFlag) == "s" {
		var err error
		*versionFlag, err = versions.GetLatestSnapshotVersion()
		if err != nil {
			util.LogCon("No snapshot version: ", err, util.Error)
			return false
		}
		return true
	}
	util.LogCon("Please provide a valid version: ", "", util.Warning)
	versions.PrintVersions()
	return false
}

func checkBuildFlag(builds paper_api.Builds, buildFlag *string) bool {
	if stringArrayContains(latest.ConvertIntArrayToStringArray(builds.Builds), *buildFlag) {
		return true
	} else if strings.ToLower(*buildFlag) == "l" {
		var err error
		*buildFlag, err = builds.GetLatestBuild()
		if err != nil {
			util.LogCon("No latest build: ", err, util.Error)
			return false
		}
		return true
	} else if strings.ToLower(*buildFlag) == "s" {
		var err error
		*buildFlag, err = builds.GetLatestSnapshotBuild()
		if err != nil {
			util.LogCon("No snapshot build: ", err, util.Error)
			return false
		}
		return true
	}
	util.LogCon("Please provide a valid build:", "", util.Warning)
	builds.PrintBuilds()
	return false
}

func stringArrayContains(arr []string, target string) bool {
	for _, s := range arr {
		if s == target {
			return true
		}
	}
	return false
}
