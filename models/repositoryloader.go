package models

import (
	"sf/utils"
	"strings"

	"golang.org/x/exp/slices"
)

type RepositoryLoader struct {
	Repository   *Repository
	Loads        []string
	GitRepo      bool
	Shapers      bool
	Framers      bool
	GitRepoLoads []string
	Preloads     []string
}

func NewRepositoryLoader(w *Repository, loads []string) (dl *RepositoryLoader) {
	dl = &RepositoryLoader{
		Repository: w,
		Loads:      loads,
	}
	return
}

func (dl *RepositoryLoader) load() (err error) {
	dl.dependencies()
	dl.settle()
	err = dl.assign()
	return
}

func (dl *RepositoryLoader) dependencies() {
	if slices.Contains(dl.Loads, "Shapers") || slices.Contains(dl.Loads, "Framers") {
		dl.Loads = append(dl.Loads,
			"GitRepo",
		)
	}
	utils.UniqStrings(&dl.Loads)
}

func (dl *RepositoryLoader) settle() {
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

func (dl *RepositoryLoader) assign() (err error) {
	if dl.GitRepo {
		dl.SetGitRepo()
		if err = dl.Repository.GitRepo.Load(dl.GitRepoLoads...); err != nil {
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

func (dl *RepositoryLoader) SetGitRepo() {
	dl.Repository.GitRepo = NewGitRepo(dl.Repository.Workspace, dl.Repository.Path)
}

func (dl *RepositoryLoader) SetFramers() {
	dl.Repository.Framers = dl.Repository.GitRepo.Framers
}

func (dl *RepositoryLoader) SetShapers() {
	dl.Repository.Shapers = dl.Repository.GitRepo.Shapers
}
