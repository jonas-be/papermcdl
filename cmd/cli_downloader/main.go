package main

import (
	"flag"
	"fmt"
	"papermc-downloader/internal/gui"
	"papermc-downloader/pkg/latest"
	"papermc-downloader/pkg/paper_api"
	"strings"
)

func main() {
	projectFlag := flag.String("p", "", "Project you want to download.")
	versionFlag := flag.String("v", "", "Version or Version Group of the Project. You can use\"l\" to get the latest version.")
	buildFlag := flag.String("b", "l", "(Optional) Build Number of the Project. You can use\"l\" to get the latest version.")
	infoFlag := flag.Bool("i", false, "(Optional) Show info or list available only.")
	flag.Parse()

	papermcAPI := paper_api.PapermcAPI{URL: "https://papermc.io"}

	if *projectFlag == "" && *versionFlag == "" && *buildFlag == "l" && *infoFlag == false {
		gui.StartGUI(papermcAPI)
		return
	}

	projects, err := papermcAPI.GetProjects()
	if err != nil {
		fmt.Println("error getting projects: ", err)
		return
	}
	if !checkProjectFlag(projects, projectFlag) {
		return
	}

	versions, err := papermcAPI.GetVersions(*projectFlag)
	if err != nil {
		fmt.Println("error getting projects: ", err)
		return
	}
	if !checkVersionFlag(versions, versionFlag) {
		return
	}

	builds, err := papermcAPI.GetBuilds(*projectFlag, *versionFlag)
	if err != nil {
		fmt.Println("error getting projects: ", err)
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
			fmt.Println("can not get buildInfo: ", err)
			return
		}
		buildInfo.PrintBuildInfo()
		return
	}

	err = papermcAPI.Download(*projectFlag, *versionFlag, *buildFlag)
	if err != nil {
		fmt.Println("can not download: ", err)
	}
}

func checkProjectFlag(projects paper_api.Projects, projectFlag *string) bool {
	if stringArrayContains(projects.Projects, *projectFlag) {
		return true
	}
	fmt.Println("Please provide a valid project:")
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
			fmt.Println("No latest version:", err)
			return false
		}
		return true
	}
	fmt.Println("Please provide a valid version:")
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
			fmt.Println("No latest build:", err)
			return false
		}
		return true
	}
	fmt.Println("Please provide a valid build:")
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
