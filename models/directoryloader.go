package models

import (
	"sf/database/queries"
	"sf/utils"
	"strings"

	"golang.org/x/exp/slices"
)

type DirectoryLoader struct {
	Directory    *Directory
	Loads        []string
	GitRepo      bool
	Shapers      bool
	Framers      bool
	GitRepoLoads []string
	Preloads     []string
}

func NewDirectoryLoader(w *Directory, loads []string) (dl *DirectoryLoader) {
	dl = &DirectoryLoader{
		Directory: w,
		Loads:     loads,
	}
	return
}

func (dl *DirectoryLoader) load() (err error) {
	dl.dependencies()
	dl.settle()
	dl.query()
	err = dl.assign()
	return
}

func (dl *DirectoryLoader) dependencies() {
	if slices.Contains(dl.Loads, "Shapers") || slices.Contains(dl.Loads, "Framers") {
		dl.Loads = append(dl.Loads,
			"GitRepo",
		)
	}
	utils.UniqStrings(&dl.Loads)
}

func (dl *DirectoryLoader) settle() {
	for _, load := range dl.Loads {
		elem := strings.SplitN(load, ".", 2)
		switch elem[0] {
		case "GitRepo":
			dl.GitRepo = true
		case "Shapers":
			dl.Shapers = true
			dl.GitRepoLoads = append(dl.GitRepoLoads, "Shapers")
		case "Framers":
			dl.Framers = true
			dl.GitRepoLoads = append(dl.GitRepoLoads, "Framers")
		default:
			dl.Preloads = append(dl.Preloads, load)
		}
	}
}

func (dl *DirectoryLoader) query() {
	queries.Load(dl.Directory, dl.Directory.ID, dl.Preloads...)
}

func (dl *DirectoryLoader) assign() (err error) {
	if dl.GitRepo {
		dl.SetGitRepo()
		if err = dl.Directory.GitRepo.Load(dl.GitRepoLoads...); err != nil {
			return
		}
	}
	if dl.Shapers {
		dl.SetShapers()
	}
	if dl.Framers {
		dl.SetFramers()
	}
	return
}

func (dl *DirectoryLoader) SetGitRepo() {
	dl.Directory.GitRepo = NewGitRepo(dl.Directory.Workspace, dl.Directory.Path)
}

func (dl *DirectoryLoader) SetFramers() {
	dl.Directory.Framers = dl.Directory.GitRepo.Framers
}

func (dl *DirectoryLoader) SetShapers() {
	dl.Directory.Shapers = dl.Directory.GitRepo.Shapers
}
