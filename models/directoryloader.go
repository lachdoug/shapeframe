package models

import (
	"sf/database/queries"
	"sf/utils"

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
	primaries := primaryLoads(dl.Loads)
	if slices.Contains(primaries, "Shapers") {
		dl.Loads = append(dl.Loads,
			"Workspace",
			"GitRepo.Shapers",
		)
	}
	if slices.Contains(primaries, "Framers") {
		dl.Loads = append(dl.Loads,
			"Workspace",
			"GitRepo.Framers",
		)
	}
}

func (dl *DirectoryLoader) settle() {
	utils.UniqStrings(&dl.Loads)
	for _, load := range dl.Loads {
		switch primaryLoad(load) {
		case "GitRepo":
			abstractAssociation(load, "GitRepo", &dl.GitRepo, &dl.GitRepoLoads)
		case "Shapers":
			abstractAssociation(load, "Shapers", &dl.Shapers)
		case "Framers":
			abstractAssociation(load, "Framers", &dl.Framers)
		default:
			dl.Preloads = append(dl.Preloads, load)
		}
	}
}

func (dl *DirectoryLoader) query() {
	utils.UniqStrings(&dl.Preloads)
	queries.Load(dl.Directory, dl.Directory.ID, dl.Preloads...)
}

func (dl *DirectoryLoader) assign() (err error) {
	if err = dl.LoadGitRepo(); err != nil {
		return
	}
	dl.LoadShapers()
	dl.LoadFramers()
	return
}

func (dl *DirectoryLoader) LoadGitRepo() (err error) {
	if dl.GitRepo {
		dl.SetGitRepo()
		if err = dl.Directory.GitRepo.Load(dl.GitRepoLoads...); err != nil {
			return
		}
	}
	return
}

func (dl *DirectoryLoader) LoadShapers() {
	if dl.Shapers {
		dl.SetShapers()
	}
}

func (dl *DirectoryLoader) LoadFramers() {
	if dl.Framers {
		dl.SetFramers()
	}
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
