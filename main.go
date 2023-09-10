package main

import (
	"sort"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Filter and Sort")

	// Create mock data
	repositories := []string{
		"Repository 1",
		"Repository 2",
		"Repository 3",
		"Another Repository",
		"Example Repo",
	}

	repoList := widget.NewList(
		func() int {
			return len(repositories)
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(index widget.ListItemID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(repositories[index])
		},
	)

	filterEntry := widget.NewEntry()
	filterEntry.SetPlaceHolder("Filter by RepoName")

	sortButton := widget.NewButton("Sort by RepoName", func() {
		sort.Strings(repositories)
		repoList.Refresh()
	})

	container := container.NewVBox(
		repoList,
		filterEntry,
		sortButton,
	)

	filterEntry.OnChanged = func(text string) {
		filteredRepositories := []string{}
		for _, repo := range repositories {
			if strings.Contains(strings.ToLower(repo), strings.ToLower(text)) {
				filteredRepositories = append(filteredRepositories, repo)
			}
		}
		repoList.Refresh()
	}

	myWindow.SetContent(container)
	myWindow.ShowAndRun()
}
