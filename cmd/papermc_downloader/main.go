package main

import (
	"fmt"
	"papermc-downloader/pkg/paper_api"
)

func main() {
	papermcAPI := paper_api.PapermcAPI{URL: "https://papermc.io"}

	projects, err := papermcAPI.GetProjects()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(projects)

	versions, err := papermcAPI.GetVersions("paper")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(versions)

	builds, err := papermcAPI.GetBuilds("paper", "1.19.4")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(builds)
}
