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

func NewRepositoryLoader(w *Repository, loads []string) (rl *RepositoryLoader) {
	rl = &RepositoryLoader{
		Repository: w,
		Loads:      loads,
	}
	return
}

func (rl *RepositoryLoader) load() (err error) {
	rl.dependencies()
	rl.settle()
	err = rl.assign()
	return
}

func (rl *RepositoryLoader) dependencies() {
	if slices.Contains(rl.Loads, "Shapers") || slices.Contains(rl.Loads, "Framers") {
		rl.Loads = append(rl.Loads,
			"GitRepo",
		)
	}
	utils.UniqStrings(&rl.Loads)
}

func (rl *RepositoryLoader) settle() {
	for _, load := range rl.Loads {
		elem := strings.SplitN(load, ".", 2)
		switch elem[0] {
		case "GitRepo":
			rl.GitRepo = true
		case "Shapers":
			rl.Shapers = true
			rl.GitRepoLoads = append(rl.GitRepoLoads, "Shapers")
		case "Framers":
			rl.Framers = true
			rl.GitRepoLoads = append(rl.GitRepoLoads, "Framers")
		default:
			rl.Preloads = append(rl.Preloads, load)
		}
	}
	utils.UniqStrings(&rl.Preloads)
}

func (rl *RepositoryLoader) assign() (err error) {
	if rl.GitRepo {
		rl.SetGitRepo()
		if err = rl.Repository.GitRepo.Load(rl.GitRepoLoads...); err != nil {
			return
		}
	}
	if rl.Shapers {
		rl.SetShapers()
	}
	if rl.Framers {
		rl.SetFramers()
	}
	return
}

func (rl *RepositoryLoader) SetGitRepo() {
	rl.Repository.GitRepo = NewGitRepo(rl.Repository.Workspace, rl.Repository.directory())
}

func (rl *RepositoryLoader) SetFramers() {
	rl.Repository.Framers = rl.Repository.GitRepo.Framers
}

func (rl *RepositoryLoader) SetShapers() {
	rl.Repository.Shapers = rl.Repository.GitRepo.Shapers
}
